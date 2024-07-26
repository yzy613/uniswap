package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/frame/g"
	"uniswap/internal/biz"
	"uniswap/internal/dao"
	"uniswap/internal/model/entity"
)

var _ biz.TickRepo = (*tickRepo)(nil)

type tickRepo struct {
	data *Data
	log  *log.Helper
}

func NewTickRepo(data *Data, logger log.Logger) biz.TickRepo {
	return &tickRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *tickRepo) Get(poolId int64, tickIndex int32) (*biz.Tick, error) {
	ctx := context.TODO()
	var tick *entity.Tick

	err := dao.Tick.Ctx(ctx).
		Where(g.Map{
			dao.Tick.Columns().PoolId:    poolId,
			dao.Tick.Columns().TickIndex: tickIndex,
		}).
		Scan(&tick)
	if err != nil {
		return nil, err
	}
	if tick == nil {
		return nil, nil
	}

	return &biz.Tick{Tick: *tick}, nil
}

func (r *tickRepo) Save(tick *biz.Tick) error {
	ctx := context.TODO()

	_, err := dao.Tick.Ctx(ctx).
		Data(tick).
		OnConflict(
			dao.Tick.Columns().PoolId,
			dao.Tick.Columns().TickIndex,
		).
		Save()

	return err
}
