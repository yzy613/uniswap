package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"uniswap/internal/biz"
	"uniswap/internal/dao"
	"uniswap/internal/model/entity"
)

var _ biz.LiquidityRepo = (*liquidityRepo)(nil)

type liquidityRepo struct {
	data *Data
	log  *log.Helper
}

func NewLiquidityRepo(data *Data, logger log.Logger) biz.LiquidityRepo {
	return &liquidityRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *liquidityRepo) GetLiquidity(poolId int64) (*biz.Liquidity, error) {
	ctx := context.TODO()
	var l *entity.Liquidity

	err := dao.Liquidity.Ctx(ctx).
		Where(dao.Liquidity.Columns().PoolId, poolId).
		Scan(&l)
	if err != nil {
		return nil, err
	}
	if l == nil {
		return nil, nil
	}

	return &biz.Liquidity{Liquidity: *l}, nil
}

func (r *liquidityRepo) SaveLiquidity(liquidity biz.Liquidity) error {
	if liquidity.PoolId == 0 {
		return errors.BadRequest("INVALID_POOL_ID", "invalid pool id")
	}

	ctx := context.TODO()

	_, err := dao.Liquidity.Ctx(ctx).
		Data(liquidity).
		OnConflict(dao.Liquidity.Columns().PoolId).
		Save()

	return err
}
