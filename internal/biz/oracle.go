package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"math/big"
	"uniswap/internal/dao"
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
			PoolId:                        pool.Id,
			ObservationIndex:              0,
			BlockTimestamp:                int(time),
			TickCumulative:                0,
			SecondsPerLiquidityCumulative: decimal.Zero,
			Initialized:                   1,
		},
	}

	err = uc.repo.SaveObservation(o)
	if err != nil {
		return
	}

	return 1, 1, nil
}

func (uc *ObservationUsecase) Write(poolId int64, index uint16, blockTimestamp uint32, tick int32,
	liquidity decimal.Decimal, cardinality, cardinalityNext uint16,
) (indexUpdated, cardinalityUpdated uint16, err error) {
	last, err := uc.repo.GetObservation(poolId, index)
	if err != nil {
		return 0, 0, err
	}

	if last.BlockTimestamp == int(blockTimestamp) {
		return index, cardinality, nil
	}

	if cardinalityNext > cardinality && index == (cardinality-1) {
		cardinalityUpdated = cardinalityNext
	} else {
		cardinalityUpdated = cardinality
	}

	indexUpdated = (index + 1) % cardinalityUpdated

	err = uc.repo.SaveObservation(*uc.transform(*last, blockTimestamp, tick, liquidity))
	if err != nil {
		return 0, 0, err
	}

	return
}

func (uc *ObservationUsecase) transform(
	last Observation, blockTimestamp uint32, tick int32, liquidity decimal.Decimal) *Observation {
	delta := blockTimestamp - uint32(last.BlockTimestamp)

	liquidityInt := liquidity.BigInt()
	if liquidity.LessThanOrEqual(decimal.Zero) {
		liquidityInt = big.NewInt(1)
	}
	deltaInt := big.NewInt(int64(delta))
	_splc := new(big.Int).Lsh(deltaInt, 128)
	_splc = new(big.Int).Div(_splc, liquidityInt)
	splc := decimal.NewFromBigInt(_splc, 0)

	return &Observation{
		Observation: entity.Observation{
			BlockTimestamp:                int(blockTimestamp),
			TickCumulative:                last.TickCumulative + int(tick)*int(delta),
			SecondsPerLiquidityCumulative: splc,
			Initialized:                   1,
		},
	}
}

func (uc *ObservationUsecase) lte(time, a, b uint32) bool {
	if a <= time && b <= time {
		return a <= b
	}

	var aAdjusted, bAdjusted uint64
	if a > time {
		aAdjusted = uint64(a)
	} else {
		aAdjusted = uint64(a) + 1<<32
	}
	if b > time {
		bAdjusted = uint64(b)
	} else {
		bAdjusted = uint64(b) + 1<<32
	}

	return aAdjusted <= bAdjusted
}

func (uc *ObservationUsecase) binarySearch(poolId int64, time, target uint32, index uint16, cardinality uint16,
) (beforeOrAt, atOrAfter *Observation, err error) {
	ctx := context.TODO()

	err = dao.Observation.Ctx(ctx).
		Where(dao.Observation.Columns().PoolId, poolId).
		WhereLTE(dao.Observation.Columns().BlockTimestamp, target).
		OrderDesc(dao.Observation.Columns().BlockTimestamp).
		Scan(&beforeOrAt)
	if err != nil {
		return nil, nil, err
	}

	err = dao.Observation.Ctx(ctx).
		Where(dao.Observation.Columns().PoolId, poolId).
		WhereGTE(dao.Observation.Columns().BlockTimestamp, target).
		OrderAsc(dao.Observation.Columns().BlockTimestamp).
		Scan(&atOrAfter)
	if err != nil {
		return nil, nil, err
	}

	return
}

func (uc *ObservationUsecase) getSurroundingObservations(poolId int64, time, target uint32,
	tick int32, index uint16, liquidity decimal.Decimal, cardinality uint16,
) (beforeOrAt, atOrAfter *Observation, err error) {
	beforeOrAt, err = uc.repo.GetObservation(poolId, index)
	if err != nil {
		return nil, nil, err
	}

	if uc.lte(time, uint32(beforeOrAt.BlockTimestamp), target) {
		if beforeOrAt.BlockTimestamp == int(target) {
			return
		} else {
			return beforeOrAt, uc.transform(*beforeOrAt, target, tick, liquidity), nil
		}
	}

	beforeOrAt, err = uc.repo.GetObservation(poolId, (index+1)%cardinality)
	if err != nil {
		return nil, nil, err
	}
	if beforeOrAt.Initialized == 0 {
		beforeOrAt, err = uc.repo.GetObservation(poolId, 0)
		if err != nil {
			return nil, nil, err
		}
	}

	if !uc.lte(time, uint32(beforeOrAt.BlockTimestamp), target) {
		return nil, nil, errors.BadRequest("NO_OBSERVATIONS", "no observations")
	}

	return uc.binarySearch(poolId, time, target, index, cardinality)
}

func (uc *ObservationUsecase) ObserveSingle(poolId int64, time, secondsAgo uint32, tick int32, index uint16,
	liquidity decimal.Decimal, cardinality uint16,
) (tickCumulative int64, secondsPerLiquidityCumulative decimal.Decimal, err error) {
	if secondsAgo == 0 {
		var last *Observation
		last, err = uc.repo.GetObservation(poolId, index)
		if err != nil {
			return
		}

		if last.BlockTimestamp != int(time) {
			last = uc.transform(*last, time, tick, liquidity)
			return int64(last.TickCumulative), last.SecondsPerLiquidityCumulative, nil
		}
	}

	target := time - secondsAgo

	beforeOrAt, atOrAfter, err :=
		uc.getSurroundingObservations(poolId, time, target, tick, index, liquidity, cardinality)
	if err != nil {
		return 0, decimal.Decimal{}, err
	}

	if target == uint32(beforeOrAt.BlockTimestamp) {
		return int64(beforeOrAt.TickCumulative), beforeOrAt.SecondsPerLiquidityCumulative, nil
	} else if target == uint32(atOrAfter.BlockTimestamp) {
		return int64(atOrAfter.TickCumulative), atOrAfter.SecondsPerLiquidityCumulative, nil
	} else {
		observationTimeDelta := uint32(atOrAfter.BlockTimestamp - beforeOrAt.BlockTimestamp)
		targetDelta := target - uint32(beforeOrAt.BlockTimestamp)

		tickCumulative = int64(beforeOrAt.TickCumulative) +
			((int64(atOrAfter.TickCumulative)-int64(beforeOrAt.TickCumulative))/int64(observationTimeDelta))*
				int64(targetDelta)
		secondsPerLiquidityCumulative =
			beforeOrAt.SecondsPerLiquidityCumulative.Add(
				atOrAfter.SecondsPerLiquidityCumulative.Sub(
					beforeOrAt.SecondsPerLiquidityCumulative).Mul(
					decimal.NewFromInt(int64(targetDelta))).Div(
					decimal.NewFromInt(int64(observationTimeDelta))))
	}

	return
}
