package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"uniswap/internal/model/entity"
)

type Tick struct {
	entity.Tick
}

type TickRepo interface {
	Get(poolId int64, tick int32) (*Tick, error)
	Save(tick *Tick) error
	Clear(poolId int64, tick int32) error
}

type TickUsecase struct {
	repo TickRepo
	log  *log.Helper
}

func NewTickUsecase(repo TickRepo, logger log.Logger) *TickUsecase {
	return &TickUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TickUsecase) Update(poolId int64, tick, tickCurrent int32,
	liquidityDelta, feeGrowthGlobal0, feeGrowthGlobal1, secondsPerLiquidityCumulative decimal.Decimal,
	tickCumulative int64, time uint32, upper bool, maxLiquidity decimal.Decimal,
) (flipped bool, err error) {
	info, err := uc.repo.Get(poolId, tick)
	if err != nil {
		return
	}

	liquidityGrossBefore := info.LiquidityGross
	liquidityGrossAfter := liquidityGrossBefore.Add(liquidityDelta)

	if liquidityGrossAfter.GreaterThan(maxLiquidity) {
		return false, errors.BadRequest("LIQUIDITY_EXCEEDS_MAXIMUM", "liquidity exceeds maximum")
	}

	flipped = liquidityGrossAfter.IsZero() != liquidityGrossBefore.IsZero()

	if liquidityGrossBefore.IsZero() {
		if tick <= tickCurrent {
			info.FeeGrowthOutside0 = feeGrowthGlobal0
			info.FeeGrowthOutside1 = feeGrowthGlobal1
			info.SecondsPerLiquidityOutside = secondsPerLiquidityCumulative
			info.TickCumulativeOutside = decimal.NewFromInt(tickCumulative)
			info.SecondsOutside = decimal.NewFromInt(int64(time))
		}
		info.Initialized = 1
	}

	info.LiquidityGross = liquidityGrossAfter

	if upper {
		info.LiquidityNet = info.LiquidityNet.Sub(liquidityDelta)
	} else {
		info.LiquidityNet = info.LiquidityNet.Add(liquidityDelta)
	}

	err = uc.repo.Save(info)

	return
}

func (uc *TickUsecase) GetFeeGrowthInside(poolId int64,
	tickLower, tickUpper, tickCurrent int32, feeGrowthGlobal0, feeGrowthGlobal1 decimal.Decimal,
) (feeGrowthInside0, feeGrowthInside1 decimal.Decimal, err error) {
	lower, err := uc.repo.Get(poolId, tickLower)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}
	upper, err := uc.repo.Get(poolId, tickUpper)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}

	var feeGrowthBelow0, feeGrowthBelow1 decimal.Decimal

	if tickCurrent >= tickLower {
		feeGrowthBelow0 = lower.FeeGrowthOutside0
		feeGrowthBelow1 = lower.FeeGrowthOutside1
	} else {
		feeGrowthBelow0 = feeGrowthGlobal0.Sub(lower.FeeGrowthOutside0)
		feeGrowthBelow1 = feeGrowthGlobal1.Sub(lower.FeeGrowthOutside1)
	}

	var feeGrowthAbove0, feeGrowthAbove1 decimal.Decimal

	if tickCurrent < tickUpper {
		feeGrowthAbove0 = upper.FeeGrowthOutside0
		feeGrowthAbove1 = upper.FeeGrowthOutside1
	} else {
		feeGrowthAbove0 = feeGrowthGlobal0.Sub(upper.FeeGrowthOutside0)
		feeGrowthAbove1 = feeGrowthGlobal1.Sub(upper.FeeGrowthOutside1)
	}

	feeGrowthInside0 = feeGrowthGlobal0.Sub(feeGrowthBelow0).Sub(feeGrowthAbove0)
	feeGrowthInside1 = feeGrowthGlobal1.Sub(feeGrowthBelow1).Sub(feeGrowthAbove1)

	return
}

func (uc *TickUsecase) Clear(poolId int64, tick int32) error {
	return uc.repo.Clear(poolId, tick)
}

func (uc *TickUsecase) Cross(
	poolId int64,
	tick int32,
	feeGrowthGlobal0,
	feeGrowthGlobal1,
	secondsPerLiquidityCumulative decimal.Decimal,
	tickCumulative int64,
	time uint32,
) (liquidityNet decimal.Decimal, err error) {
	info, err := uc.repo.Get(poolId, tick)
	if err != nil {
		return
	}

	info.FeeGrowthOutside0 = feeGrowthGlobal0.Sub(info.FeeGrowthOutside0)
	info.FeeGrowthOutside1 = feeGrowthGlobal1.Sub(info.FeeGrowthOutside1)
	info.SecondsPerLiquidityOutside = secondsPerLiquidityCumulative.Sub(info.SecondsPerLiquidityOutside)
	info.TickCumulativeOutside = decimal.NewFromInt(tickCumulative).Sub(info.TickCumulativeOutside)
	info.SecondsOutside = decimal.NewFromInt(int64(time)).Sub(info.SecondsOutside)
	liquidityNet = info.LiquidityNet

	err = uc.repo.Save(info)

	return
}
