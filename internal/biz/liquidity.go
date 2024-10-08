package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"uniswap/internal/model/entity"
)

type Liquidity struct {
	entity.Liquidity
}

type LiquidityRepo interface {
	GetLiquidity(poolId int64) (*Liquidity, error)
	SaveLiquidity(liquidity Liquidity) error
}

type LiquidityUsecase struct {
	repo LiquidityRepo
	log  *log.Helper
}

func NewLiquidityUsecase(repo LiquidityRepo, logger log.Logger) *LiquidityUsecase {
	return &LiquidityUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *LiquidityUsecase) Get(poolId int64) (decimal.Decimal, error) {
	l, err := uc.repo.GetLiquidity(poolId)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return l.Liquidity.Liquidity, nil
}

func (uc *LiquidityUsecase) Save(poolId int64, liquidity decimal.Decimal) error {
	l := Liquidity{entity.Liquidity{
		PoolId:    poolId,
		Liquidity: liquidity,
	}}

	return uc.repo.SaveLiquidity(l)
}
