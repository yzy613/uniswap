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
}

type TickUsecase struct {
	repo TickRepo
	log  *log.Helper
}

func NewTickUsecase(repo TickRepo, logger log.Logger) *TickUsecase {
	return &TickUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TickUsecase) Update(poolId int64, tick, tickCurrent int32,
	liquidityDelta, feeGrowthGlobal0X128, feeGrowthGlobal1X128, secondsPerLiquidityCumulativeX128 decimal.Decimal,
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
			info.FeeGrowthOutside0X128 = feeGrowthGlobal0X128
			info.FeeGrowthOutside1X128 = feeGrowthGlobal1X128
			info.SecondsPerLiquidityOutsideX128 = secondsPerLiquidityCumulativeX128
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