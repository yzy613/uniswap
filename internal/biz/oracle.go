package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"uniswap/internal/model/entity"
)

type Observation struct {
	entity.Observation
}

type ObservationRepo interface {
	GetObservation(poolId int64, index uint16) (*Observation, error)
	SaveObservation(observation Observation) error
}

type ObservationUsecase struct {
	repo ObservationRepo
	log  *log.Helper
}

func NewObservationUsecase(repo ObservationRepo, logger log.Logger) *ObservationUsecase {
	return &ObservationUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *ObservationUsecase) Initialize(pool Pool, time uint32) (cardinality, cardinalityNext uint16, err error) {
	o := Observation{
		Observation: entity.Observation{
			PoolId:                            pool.Id,
			ObservationIndex:                  0,
			BlockTimestamp:                    int(time),
			TickCumulative:                    decimal.Zero,
			SecondsPerLiquidityCumulativeX128: decimal.Zero,
			Initialized:                       1,
		},
	}

	err = uc.repo.SaveObservation(o)
	if err != nil {
		return
	}

	return 1, 1, nil
}

func (uc *ObservationUsecase) ObserveSingle(poolId int64, time, secondsAgo uint32, tick int32, index uint16,
	liquidity decimal.Decimal, cardinality uint16,
) (tickCumulative int64, secondsPerLiquidityCumulativeX128 decimal.Decimal, err error) {
	if secondsAgo == 0 {
		var last *Observation
		last, err = uc.repo.GetObservation(poolId, index)
		if err != nil {
			return
		}
		if last.BlockTimestamp != int(time) {

		}
	}
}
