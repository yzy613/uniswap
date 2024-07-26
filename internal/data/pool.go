package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/shopspring/decimal"
	"uniswap/internal/biz"
	"uniswap/internal/dao"
	"uniswap/internal/model/entity"
)

var _ biz.PoolRepo = (*poolRepo)(nil)

type poolRepo struct {
	data *Data
	log  *log.Helper
}

func NewPoolRepo(data *Data, logger log.Logger) biz.PoolRepo {
	return &poolRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *poolRepo) GetPool(token0, token1 string, fee uint32) (*biz.Pool, error) {
	ctx := context.TODO()
	var pool *entity.Pool

	err := dao.Pool.Ctx(ctx).
		Where(g.Map{
			dao.Pool.Columns().Token0Address: token0,
			dao.Pool.Columns().Token1Address: token1,
			dao.Pool.Columns().Fee:           fee,
		}).
		Scan(&pool)
	if err != nil {
		return nil, err
	}
	if pool == nil {
		return nil, nil
	}

	return &biz.Pool{Pool: *pool}, nil
}

func (r *poolRepo) CreatePool(token0, token1 string,
	fee uint32, tickSpacing int32, tickSpacingToMaxLiquidityPerTick decimal.Decimal,
) (*biz.Pool, error) {
	ctx := context.TODO()

	pool := entity.Pool{
		Token0Address:       token0,
		Token1Address:       token1,
		Fee:                 int(fee),
		TickSpacing:         int(tickSpacing),
		MaxLiquidityPerTick: tickSpacingToMaxLiquidityPerTick,
	}

	result, err := dao.Pool.Ctx(ctx).
		Data(pool).OmitEmptyData().
		Save()
	if err != nil {
		return nil, err
	}

	pool.Id, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &biz.Pool{Pool: pool}, nil
}

func (r *poolRepo) FeeAmountTickSpacing(fee uint32) (tickSpacing int32) {
	switch fee {
	case 500:
		tickSpacing = 10
	case 3000:
		tickSpacing = 60
	case 10000:
		tickSpacing = 200
	}

	ctx := context.TODO()
	var feeAmount *entity.FeeAmount

	err := dao.FeeAmount.Ctx(ctx).
		Where(dao.FeeAmount.Columns().Fee, fee).
		Scan(&feeAmount)
	if err != nil {
		return
	}

	return int32(feeAmount.TickSpacing)
}
