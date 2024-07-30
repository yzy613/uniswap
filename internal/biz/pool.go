package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"time"
	"uniswap/internal/lib/tickmath"
	"uniswap/internal/model/entity"
	"uniswap/internal/util"
)

const (
	// TODO: 精度问题
	//MinTick = -887272
	MinTick = 886463
	MaxTick = -MinTick
)

type Slot0 struct {
	entity.Slot0
}

type Pool struct {
	entity.Pool
	slot0 *Slot0
}

type Position struct {
	entity.Position
}

type PoolRepo interface {
	GetPool(token0, token1 string, fee uint32) (*Pool, error)
	CreatePool(token0, token1 string, fee uint32, tickSpacing int32,
		tickSpacingToMaxLiquidityPerTick decimal.Decimal) (*Pool, error)
	FeeAmountTickSpacing(fee uint32) (tickSpacing int32)
	GetPosition(poolId int64, owner string, tickLower, tickUpper int32) (*Position, error)
	UpdatePosition(position Position,
		liquidityDelta, feeGrowthInside0, feeGrowthInside1 decimal.Decimal) error
	GetSlot0(poolId int64) (*Slot0, error)
	SaveSlot0(slot0 Slot0) error
	GetFeeGrowthGlobal(poolId int64) (feeGrowthGlobal0, feeGrowthGlobal1 decimal.Decimal, err error)
}

type PoolUsecase struct {
	repo PoolRepo
	log  *log.Helper

	tickUsecase *TickUsecase
	observation *ObservationUsecase
	tickBitmap  *TickBitmapUsecase
}

func NewPoolUsecase(repo PoolRepo, logger log.Logger, tickUsecase *TickUsecase) *PoolUsecase {
	return &PoolUsecase{repo: repo, log: log.NewHelper(logger), tickUsecase: tickUsecase}
}

func (uc *PoolUsecase) initialize(pool Pool, price decimal.Decimal) error {
	slot0, err := uc.repo.GetSlot0(pool.Id)
	if err != nil {
		return err
	}
	if slot0 != nil && !slot0.Price.IsZero() {
		return errors.BadRequest("SLOT0_EXISTS", "slot0 already exists")
	}
	tick, err := tickmath.GetTickAtSqrtRatio(price)
	if err != nil {
		return err
	}
	// TODO: _blockTimestamp()
	cardinality, cardinalityNext, err := uc.observation.Initialize(pool, uint32(time.Now().Unix()))
	if err != nil {
		return err
	}

	return uc.repo.SaveSlot0(Slot0{
		Slot0: entity.Slot0{
			PoolId:                     pool.Id,
			Price:                      price,
			CurrentTick:                int(tick),
			ObservationIndex:           0,
			ObservationCardinality:     int(cardinality),
			ObservationCardinalityNext: int(cardinalityNext),
			FeeProtocol:                0,
			Unlocked:                   1,
		},
	})
}

func (uc *PoolUsecase) CreatePool(token0, token1 string, fee uint32) (*Pool, error) {
	if token0 == token1 {
		return nil, errors.BadRequest("INVALID_TOKEN", "token0 and token1 must be different")
	}
	if token0 > token1 {
		token0, token1 = token1, token0
	}
	if token0 == "" || token0 == "0" {
		return nil, errors.BadRequest("INVALID_TOKEN", "token0 is invalid")
	}
	tickSpacing := uc.repo.FeeAmountTickSpacing(fee)
	if tickSpacing == 0 {
		return nil, errors.BadRequest("INVALID_FEE", "fee is invalid")
	}
	if p, e := uc.repo.GetPool(token0, token1, fee); e == nil && p != nil {
		return nil, errors.BadRequest("POOL_EXISTS", "pool already exists")
	}

	return uc.repo.CreatePool(token0, token1, fee, tickSpacing, uc.tickSpacingToMaxLiquidityPerTick(tickSpacing))
}

// TODO: 用这个来测试精度
func (uc *PoolUsecase) tickSpacingToMaxLiquidityPerTick(tickSpacing int32) decimal.Decimal {
	minTick := MinTick / tickSpacing * tickSpacing
	maxTick := MaxTick / tickSpacing * tickSpacing
	numTicks := uint32((maxTick-minTick)/tickSpacing) + 1

	return util.BitMaxNumDecimal(128).DivRound(decimal.NewFromInt(int64(numTicks)), 0)
}

func (uc *PoolUsecase) checkTick(tickLower, tickUpper int32) error {
	if tickLower >= tickUpper {
		return errors.BadRequest("INVALID_TICK", "tickLower must be less than tickUpper")
	}
	if tickLower < MinTick || tickUpper > MaxTick {
		return errors.BadRequest("INVALID_TICK", "tickLower and tickUpper must be within the range")
	}
	return nil
}

func (uc *PoolUsecase) Mint(pool Pool, tickLower, tickUpper int32, amount decimal.Decimal,
) (amount0, amount1 decimal.Decimal, err error) {

}

func (uc *PoolUsecase) modifyPosition(owner string, tickLower, tickUpper int32, liquidityDelta decimal.Decimal,
) (position *Position, amount0, amount1 decimal.Decimal, err error) {
	err = uc.checkTick(tickLower, tickUpper)
	if err != nil {
		return
	}
}

func (uc *PoolUsecase) updatePosition(pool Pool, owner string, tickLower, tickUpper int32, liquidityDelta decimal.Decimal,
	tick int32) (*Position, error) {
	position, err := uc.repo.GetPosition(pool.Id, owner, tickLower, tickUpper)
	if err != nil {
		return nil, err
	}
	feeGrowthGlobal0, feeGrowthGlobal1, err := uc.repo.GetFeeGrowthGlobal(pool.Id)
	if err != nil {
		return nil, err
	}
	var flippedLower, flippedUpper bool
	if !liquidityDelta.IsZero() {
		// TODO: _blockTimestamp()
		time_ := uint32(time.Now().Unix())
		tickCumulative, secondsPerLiquidityCumulative, err := uc.observation.ObserveSingle(
			pool.Id,
			time_,
			0,
			tick,
			uint16(pool.slot0.ObservationIndex),
			liquidityDelta,
			uint16(pool.slot0.ObservationCardinality),
		)
		if err != nil {
			return nil, err
		}
		flippedLower, err = uc.tickUsecase.Update(pool.Id, tick, tickLower, liquidityDelta, feeGrowthGlobal0,
			feeGrowthGlobal1, secondsPerLiquidityCumulative, tickCumulative, time_,
			false, pool.MaxLiquidityPerTick)
		if err != nil {
			return nil, err
		}
		flippedUpper, err = uc.tickUsecase.Update(pool.Id, tick, tickUpper, liquidityDelta, feeGrowthGlobal0,
			feeGrowthGlobal1, secondsPerLiquidityCumulative, tickCumulative, time_,
			true, pool.MaxLiquidityPerTick)
		if err != nil {
			return nil, err
		}

		if flippedLower {
			err = uc.tickBitmap.FlipTick(pool.Id, tickLower, int32(pool.TickSpacing))
			if err != nil {
				return nil, err
			}
		}
		if flippedUpper {
			err = uc.tickBitmap.FlipTick(pool.Id, tickUpper, int32(pool.TickSpacing))
			if err != nil {
				return nil, err
			}
		}

		feeGrowthInside0, feeGrowthInside1 :=
			uc.tickUsecase.
	}
}
