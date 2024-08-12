package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
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
		OnConflict(
			dao.Pool.Columns().Token0Address,
			dao.Pool.Columns().Token1Address,
			dao.Pool.Columns().Fee,
		).
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

func (r *poolRepo) GetPosition(poolId int64, owner string, tickLower, tickUpper int32) (*biz.Position, error) {
	ctx := context.TODO()
	var position *entity.Position

	err := dao.Position.Ctx(ctx).
		Where(g.Map{
			dao.Position.Columns().PoolId:       poolId,
			dao.Position.Columns().OwnerAddress: owner,
			dao.Position.Columns().TickLower:    tickLower,
			dao.Position.Columns().TickUpper:    tickUpper,
		}).
		Scan(&position)
	if err != nil {
		return nil, err
	}
	if position == nil {
		return nil, nil
	}

	return &biz.Position{Position: *position}, nil
}

func (r *poolRepo) UpdatePosition(position biz.Position,
	liquidityDelta, feeGrowthInside0, feeGrowthInside1 decimal.Decimal,
) error {
	ctx := context.TODO()
	data := entity.Position{
		Liquidity:            position.Liquidity.Add(liquidityDelta),
		FeeGrowthInside0Last: feeGrowthInside0,
		FeeGrowthInside1Last: feeGrowthInside1,
	}

	_, err := dao.Position.Ctx(ctx).
		Data(data).OmitEmptyData().
		Where(g.Map{
			dao.Position.Columns().PoolId:       position.PoolId,
			dao.Position.Columns().OwnerAddress: position.OwnerAddress,
			dao.Position.Columns().TickLower:    position.TickLower,
			dao.Position.Columns().TickUpper:    position.TickUpper,
		}).
		Update()

	return err
}

func (r *poolRepo) SavePosition(position biz.Position) error {
	ctx := context.TODO()

	_, err := dao.Position.Ctx(ctx).
		Data(position.Position).OmitEmptyData().
		OnConflict(
			dao.Position.Columns().PoolId,
			dao.Position.Columns().OwnerAddress,
			dao.Position.Columns().TickLower,
			dao.Position.Columns().TickUpper,
		).
		Save()

	return err
}

func (r *poolRepo) GetSlot0(poolId int64) (*biz.Slot0, error) {
	ctx := context.TODO()
	var slot0 *entity.Slot0

	err := dao.Slot0.Ctx(ctx).
		Where(dao.Slot0.Columns().PoolId, poolId).
		Scan(&slot0)
	if err != nil {
		return nil, err
	}
	if slot0 == nil {
		return nil, nil
	}

	return &biz.Slot0{Slot0: *slot0}, nil
}

func (r *poolRepo) SaveSlot0(slot0 biz.Slot0) error {
	if slot0.PoolId == 0 {
		return errors.BadRequest("INVALID_POOL_ID", "invalid pool id")
	}

	ctx := context.TODO()

	_, err := dao.Slot0.Ctx(ctx).
		Data(slot0.Slot0).OmitEmptyData().
		OnConflict(dao.Slot0.Columns().PoolId).
		Save()

	return err
}

func (r *poolRepo) TryLockSlot0(poolId int64) error {
	ctx := context.TODO()
	var slot0 *entity.Slot0

	err := dao.Slot0.Ctx(ctx).
		Fields(dao.Slot0.Columns().Unlocked).
		Where(dao.Slot0.Columns().PoolId, poolId).
		Scan(&slot0)
	if err != nil {
		return err
	}

	if slot0.Unlocked == 0 {
		return errors.BadRequest("SLOT0_LOCKED", "slot0 is locked")
	}

	_, err = dao.Slot0.Ctx(ctx).
		Data(g.Map{
			dao.Slot0.Columns().Unlocked: 1,
		}).
		Where(dao.Slot0.Columns().PoolId, poolId).
		Update()
	return err
}
func (r *poolRepo) UnlockSlot0(poolId int64) error {
	ctx := context.TODO()
	var slot0 *entity.Slot0

	err := dao.Slot0.Ctx(ctx).
		Fields(dao.Slot0.Columns().Unlocked).
		Where(dao.Slot0.Columns().PoolId, poolId).
		Scan(&slot0)
	if err != nil {
		return err
	}

	if slot0.Unlocked != 0 {
		return errors.BadRequest("SLOT0_UNLOCKED", "slot0 is unlocked")
	}

	_, err = dao.Slot0.Ctx(ctx).
		Data(g.Map{
			dao.Slot0.Columns().Unlocked: 0,
		}).
		Where(dao.Slot0.Columns().PoolId, poolId).
		Update()
	return err
}

func (r *poolRepo) GetFeeGrowthGlobal(poolId int64,
) (feeGrowthGlobal0, feeGrowthGlobal1 decimal.Decimal, err error) {
	ctx := context.TODO()
	var feeGrowthGlobal *entity.FeeGrowthGlobal

	err = dao.FeeGrowthGlobal.Ctx(ctx).
		Where(dao.FeeGrowthGlobal.Columns().PoolId, poolId).
		Scan(&feeGrowthGlobal)

	return feeGrowthGlobal.FeeGrowthGlobal0, feeGrowthGlobal.FeeGrowthGlobal1, err
}

func (r *poolRepo) SaveFeeGrowthGlobal0(poolId int64, feeGrowthGlobal0 decimal.Decimal) error {
	if feeGrowthGlobal0.IsNegative() {
		return errors.BadRequest("INVALID_FEE_GROWTH_GLOBAL0", "invalid fee growth global0")
	}

	ctx := context.TODO()

	_, err := dao.FeeGrowthGlobal.Ctx(ctx).
		Data(g.Map{
			dao.FeeGrowthGlobal.Columns().FeeGrowthGlobal0: feeGrowthGlobal0,
		}).
		Where(dao.FeeGrowthGlobal.Columns().PoolId, poolId).
		Update()

	return err
}

func (r *poolRepo) SaveFeeGrowthGlobal1(poolId int64, feeGrowthGlobal1 decimal.Decimal) error {
	if feeGrowthGlobal1.IsNegative() {
		return errors.BadRequest("INVALID_FEE_GROWTH_GLOBAL1", "invalid fee growth global1")
	}

	ctx := context.TODO()

	_, err := dao.FeeGrowthGlobal.Ctx(ctx).
		Data(g.Map{
			dao.FeeGrowthGlobal.Columns().FeeGrowthGlobal1: feeGrowthGlobal1,
		}).
		Where(dao.FeeGrowthGlobal.Columns().PoolId, poolId).
		Update()

	return err
}

func (r *poolRepo) GetProtocolFee(poolId int64) (token0, token1 decimal.Decimal, err error) {
	ctx := context.TODO()
	var protocolFee *entity.ProtocolFee

	err = dao.ProtocolFee.Ctx(ctx).
		Where(dao.ProtocolFee.Columns().PoolId, poolId).
		Scan(&protocolFee)

	return protocolFee.Token0Fees, protocolFee.Token1Fees, err
}

func (r *poolRepo) SaveProtocolFeeToken0(poolId int64, token0 decimal.Decimal) error {
	if token0.IsNegative() {
		return errors.BadRequest("INVALID_PROTOCOL_FEE_TOKEN0", "invalid protocol fee token0")
	}

	ctx := context.TODO()

	_, err := dao.ProtocolFee.Ctx(ctx).
		Data(g.Map{
			dao.ProtocolFee.Columns().Token0Fees: token0,
		}).
		Where(dao.ProtocolFee.Columns().PoolId, poolId).
		Update()

	return err
}

func (r *poolRepo) SaveProtocolFeeToken1(poolId int64, token1 decimal.Decimal) error {
	if token1.IsNegative() {
		return errors.BadRequest("INVALID_PROTOCOL_FEE_TOKEN1", "invalid protocol fee token1")
	}

	ctx := context.TODO()

	_, err := dao.ProtocolFee.Ctx(ctx).
		Data(g.Map{
			dao.ProtocolFee.Columns().Token1Fees: token1,
		}).
		Where(dao.ProtocolFee.Columns().PoolId, poolId).
		Update()

	return err
}

func (r *poolRepo) GetBalance(poolId int64) (balance0, balance1 decimal.Decimal, err error) {
	ctx := context.TODO()
	var pool *entity.Pool

	err = dao.Pool.Ctx(ctx).
		Fields(dao.Pool.Columns().Balance0, dao.Pool.Columns().Balance1).
		Where(dao.Pool.Columns().Id, poolId).
		Scan(&pool)

	return pool.Balance0, pool.Balance1, err
}

func (r *poolRepo) SetBalance(poolId int64, balance0, balance1 decimal.Decimal) error {
	ctx := context.TODO()

	_, err := dao.Pool.Ctx(ctx).
		Data(g.Map{
			dao.Pool.Columns().Balance0: balance0,
			dao.Pool.Columns().Balance1: balance1,
		}).
		Where(dao.Pool.Columns().Id, poolId).
		Update()

	return err
}
