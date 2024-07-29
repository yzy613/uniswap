package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"math/big"
	"uniswap/internal/lib/bitmath"
	"uniswap/internal/model/entity"
)

type TickBitmap struct {
	entity.TickBitmap
}

type TickBitmapRepo interface {
	GetBitmap(poolId int64, wordPos int16) (string, error)
	SetBitmap(poolId int64, wordPos int16, bitmap string) error
}

type TickBitmapUsecase struct {
	repo TickBitmapRepo
	log  *log.Helper
}

func NewTickBitMapUsecase(repo TickBitmapRepo, logger log.Logger) *TickBitmapUsecase {
	return &TickBitmapUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *TickBitmapUsecase) position(tick int32) (wordPos int16, bitPos uint8) {
	wordPos = int16(tick >> 8)
	bitPos = uint8(tick % 256)
	return
}

func (uc *TickBitmapUsecase) FlipTick(poolId int64, tick, tickSpacing int32) error {
	if tick%tickSpacing != 0 {
		return errors.BadRequest("INVALID_TICK", "ensure that the tick is spaced")
	}
	wordPos, bitPos := uc.position(tick / tickSpacing)
	mask := decimal.NewFromBigInt(new(big.Int).Lsh(big.NewInt(1), uint(bitPos)), 0)
	return uc.repo.SetBitmap(poolId, wordPos, mask.String())
}

func (uc *TickBitmapUsecase) NextInitializedTickWithinOneWord(poolId int64, tick, tickSpacing int32, lte bool,
) (next int32, initialized bool, err error) {
	compressed := tick / tickSpacing
	if tick < 0 && tick%tickSpacing != 0 {
		compressed--
	}

	if lte {
		wordPos, bitPos := uc.position(compressed)
		_one := big.NewInt(1)
		_lsh := new(big.Int).Lsh(_one, uint(bitPos))
		mask := new(big.Int).Sub(_lsh, _one)
		mask = new(big.Int).Add(mask, _lsh)

		var _masked string
		_masked, err = uc.repo.GetBitmap(poolId, wordPos)
		if err != nil {
			return 0, false, err
		}
		masked, ok := new(big.Int).SetString(_masked, 10)
		if !ok {
			return 0, false, errors.BadRequest("INVALID_BITMAP", "invalid bitmap")
		}

		masked = masked.And(masked, mask)

		initialized = masked.Cmp(big.NewInt(0)) != 0
		if initialized {
			next = (compressed - int32(bitPos-bitmath.MostSignificantBit(masked))) * tickSpacing
		} else {
			next = (compressed - int32(bitPos)) * tickSpacing
		}
	} else {
		wordPos, bitPos := uc.position(compressed + 1)
		_one := big.NewInt(1)
		_lsh := new(big.Int).Lsh(_one, uint(bitPos))
		mask := new(big.Int).Sub(_lsh, _one)
		mask = mask.Not(mask)

		var _masked string
		_masked, err = uc.repo.GetBitmap(poolId, wordPos)
		if err != nil {
			return 0, false, err
		}
		masked, ok := new(big.Int).SetString(_masked, 10)
		if !ok {
			return 0, false, errors.BadRequest("INVALID_BITMAP", "invalid bitmap")
		}

		masked = masked.And(masked, mask)

		initialized = masked.Cmp(big.NewInt(0)) != 0
		if initialized {
			next = (compressed + 1 + int32(bitmath.LeastSignificantBit(masked)-bitPos)) * tickSpacing
		} else {
			next = (compressed + 1 + int32((1<<8)-1-bitPos)) * tickSpacing
		}
	}
	return
}
