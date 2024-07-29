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

func (r *tickBitmapRepo) GetBitmap(poolId int64, wordPos int16) (string, error) {
	ctx := context.TODO()
	var bitmap entity.TickBitmap

	err := dao.TickBitmap.Ctx(ctx).
		Where(g.Map{
			dao.TickBitmap.Columns().PoolId:       poolId,
			dao.TickBitmap.Columns().WordPosition: wordPos,
		}).
		Scan(&bitmap)
	if err != nil {
		return "", err
	}

	return bitmap.Bitmap, nil
}

func (r *tickBitmapRepo) SetBitmap(poolId int64, wordPos int16, bitmap string) error {
	ctx := context.TODO()

	_, err := dao.TickBitmap.Ctx(ctx).
		Data(g.Map{
			dao.TickBitmap.Columns().PoolId:       poolId,
			dao.TickBitmap.Columns().WordPosition: wordPos,
			dao.TickBitmap.Columns().Bitmap:       bitmap,
		}).
		OnConflict(dao.TickBitmap.Columns().PoolId, dao.TickBitmap.Columns().WordPosition).
		Save()

	return err
}
