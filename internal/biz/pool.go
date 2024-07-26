package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"uniswap/internal/model/entity"
	"uniswap/internal/util"
)

type Pool struct {
	entity.Pool
	slot0 *entity.Slot0
}

type PoolRepo interface {
	GetPool(token0, token1 string, fee uint32) (*Pool, error)
	CreatePool(token0, token1 string,
		fee uint32, tickSpacing int32, tickSpacingToMaxLiquidityPerTick decimal.Decimal,
	) (*Pool, error)
	FeeAmountTickSpacing(fee uint32) (tickSpacing int32)
}

type PoolUsecase struct {
	repo PoolRepo
	log  *log.Helper
}

func NewPoolUsecase(repo PoolRepo, logger log.Logger) *PoolUsecase {
	return &PoolUsecase{repo: repo, log: log.NewHelper(logger)}
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

func (uc *PoolUsecase) tickSpacingToMaxLiquidityPerTick(tickSpacing int32) decimal.Decimal {
	const (
		MinTick = -887272
		MaxTick = -MinTick
	)

	minTick := MinTick / tickSpacing * tickSpacing
	maxTick := MaxTick / tickSpacing * tickSpacing
	numTicks := uint32((maxTick-minTick)/tickSpacing) + 1

	return util.BitMaxNumDecimal(128).DivRound(decimal.NewFromInt(int64(numTicks)), 0)
}
