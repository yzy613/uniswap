package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/frame/g"
	"uniswap/internal/biz"
	"uniswap/internal/dao"
	"uniswap/internal/model/entity"
)

type tickBitmapRepo struct {
	data *Data
	log  *log.Helper
}

func NewTickBitmapRepo(data *Data, logger log.Logger) biz.TickBitmapRepo {
	return &tickBitmapRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *tickBitmapRepo) GetBit(poolId int64, position int32) (bool, error) {
	ctx := context.TODO()
	var bitmap entity.TickBitmap

	err := dao.TickBitmap.Ctx(ctx).
		Where(g.Map{
			dao.TickBitmap.Columns().PoolId:   poolId,
			dao.TickBitmap.Columns().Position: position,
		}).
		Scan(&bitmap)
	if err != nil {
		return false, err
	}

	return bitmap.Bit != 0, nil
}

func (r *tickBitmapRepo) SetBit(poolId int64, position int32, bit bool) error {
	ctx := context.TODO()

	_, err := dao.TickBitmap.Ctx(ctx).
		Data(g.Map{
			dao.TickBitmap.Columns().PoolId:   poolId,
			dao.TickBitmap.Columns().Position: position,
			dao.TickBitmap.Columns().Bit:      bit,
		}).
		OnConflict(dao.TickBitmap.Columns().PoolId, dao.TickBitmap.Columns().Position).
		Save()

	return err
}

func (r *tickBitmapRepo) GetNextTruePosition(poolId int64, position int32, lte bool, fallbackPos int32) (int32, bool, error) {
	ctx := context.TODO()
	var bitmap *entity.TickBitmap

	query := dao.TickBitmap.Ctx(ctx).
		Where(g.Map{
			dao.TickBitmap.Columns().PoolId: poolId,
			dao.TickBitmap.Columns().Bit:    true,
		})
	if lte {
		query = query.WhereLTE(dao.TickBitmap.Columns().Position, position).
			WhereGTE(dao.TickBitmap.Columns().Position, fallbackPos).
			OrderDesc(dao.TickBitmap.Columns().Position)
	} else {
		query = query.WhereGT(dao.TickBitmap.Columns().Position, position).
			WhereLTE(dao.TickBitmap.Columns().Position, fallbackPos).
			OrderAsc(dao.TickBitmap.Columns().Position)
	}
	err := query.Scan(&bitmap)
	if err != nil {
		return 0, false, err
	}

	if bitmap == nil {
		return fallbackPos, false, nil
	}

	return int32(bitmap.Position), bitmap.Bit != 0, nil
}
