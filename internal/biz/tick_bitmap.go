package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"uniswap/internal/model/entity"
)

type TickBitmap struct {
	entity.TickBitmap
}

type TickBitmapRepo interface {
	GetBit(poolId int64, position int32) (bool, error)
	SetBit(poolId int64, position int32, bit bool) error
	GetNextTruePosition(poolId int64, position int32, lte bool, fallbackPos int32) (int32, bool, error)
}

type TickBitmapUsecase struct {
	repo TickBitmapRepo
	log  *log.Helper
}

func NewTickBitMapUsecase(repo TickBitmapRepo, logger log.Logger) *TickBitmapUsecase {
	return &TickBitmapUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TickBitmapUsecase) FlipTick(poolId int64, tick, tickSpacing int32) error {
	if tick%tickSpacing != 0 {
		return errors.BadRequest("INVALID_TICK", "ensure that the tick is spaced")
	}

	position := tick / tickSpacing
	bit, err := uc.repo.GetBit(poolId, position)
	if err != nil {
		return err
	}

	return uc.repo.SetBit(poolId, position, !bit)
}

func (uc *TickBitmapUsecase) NextInitializedTickWithinOneWord(poolId int64, tick, tickSpacing int32, lte bool,
) (next int32, initialized bool, err error) {
	compressed := tick / tickSpacing
	if tick < 0 && tick%tickSpacing != 0 {
		compressed--
	}

	if lte {
		fallbackPos := compressed & ^(1<<8 - 1)

		next, initialized, err = uc.repo.GetNextTruePosition(poolId, compressed, lte, fallbackPos)
		if err != nil {
			return
		}
	} else {
		fallbackPos := compressed | (1<<8 - 1)

		next, initialized, err = uc.repo.GetNextTruePosition(poolId, compressed, lte, fallbackPos)
		if err != nil {
			return
		}
	}

	next *= tickSpacing

	return
}
