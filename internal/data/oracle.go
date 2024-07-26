package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/frame/g"
	"uniswap/internal/biz"
	"uniswap/internal/dao"
	"uniswap/internal/model/entity"
)

var _ biz.ObservationRepo = (*observationRepo)(nil)

type observationRepo struct {
	data *Data
	log  *log.Helper
}

func NewObservationRepo(data *Data, logger log.Logger) biz.ObservationRepo {
	return &observationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *observationRepo) GetObservation(poolId int64, index uint16) (*biz.Observation, error) {
	ctx := context.TODO()
	var observation *entity.Observation

	err := dao.Observation.Ctx(ctx).
		Where(g.Map{
			dao.Observation.Columns().PoolId:           poolId,
			dao.Observation.Columns().ObservationIndex: index,
		}).
		Scan(&observation)
	if err != nil {
		return nil, err
	}
	if observation == nil {
		return nil, nil
	}

	return &biz.Observation{Observation: *observation}, nil
}

func (r *observationRepo) SaveObservation(observation biz.Observation) error {
	ctx := context.TODO()

	_, err := dao.Observation.Ctx(ctx).
		Data(observation.Observation).OmitEmptyData().
		OnConflict(
			dao.Observation.Columns().PoolId,
			dao.Observation.Columns().ObservationIndex,
		).
		Save()

	return err
}
