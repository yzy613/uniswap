package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
	"uniswap/internal/lib/tickmath"
)

type PositionManagerUsecase struct {
	log *log.Helper

	pool *PoolUsecase
}

func NewPositionManagerUsecase(logger log.Logger) *PositionManagerUsecase {
	return &PositionManagerUsecase{log: log.NewHelper(logger)}
}

func (uc *PositionManagerUsecase) Mint(
	poolId int64,
	token0,
	token1 string,
	fee uint32,
	tickLower,
	tickUpper int32,
	amount0Desired,
	amount1Desired,
	amount0Min,
	amount1Min decimal.Decimal,
	recipient string,
	deadline int64,
) (tokenId, liquidity, amount0, amount1 decimal.Decimal, err error) {
	if gtime.New().Unix() > deadline {
		err = errors.BadRequest("invalid", "deadline")
		return
	}

	liquidity, amount0, amount1, pool, err := uc.addLiquidity(
		token0,
		token1,
		fee,
		"this",
		tickLower,
		tickUpper,
		amount0Desired,
		amount1Desired,
		amount0Min,
		amount1Min,
	)
	if err != nil {
		return
	}

	return
}

func (uc *PositionManagerUsecase) addLiquidity(
	token0,
	token1 string,
	fee uint32,
	recipient string,
	tickLower,
	tickUpper int32,
	amount0Desired,
	amount1Desired,
	amount0Min,
	amount1Min decimal.Decimal,
) (liquidity, amount0, amount1 decimal.Decimal, pool *Pool, err error) {
	pool, err = uc.pool.GetPool(token0, token1, fee)
	if err != nil {
		return
	}

	{
		price := pool.slot0.Price
		var ratioA, ratioB decimal.Decimal
		ratioA, err = tickmath.GetRatioAtTick(tickLower)
		if err != nil {
			return
		}
		ratioB, err = tickmath.GetRatioAtTick(tickUpper)
		if err != nil {
			return
		}

		liquidity = uc.getLiquidityForAmounts(price, ratioA, ratioB, amount0Desired, amount1Desired)
	}

	amount0, amount1, err = uc.pool.Mint(pool.Id, recipient, tickLower, tickUpper, liquidity)
	if err != nil {
		return
	}

	if amount0.LessThan(amount0Min) || amount1.LessThan(amount1Min) {
		err = errors.BadRequest("insufficient", "liquidity")
		return
	}

	return
}

func (uc *PositionManagerUsecase) getLiquidityForAmounts(
	ratio decimal.Decimal,
	ratioA decimal.Decimal,
	ratioB decimal.Decimal,
	amount0 decimal.Decimal,
	amount1 decimal.Decimal,
) decimal.Decimal {
	// Ensure ratioA <= ratioB
	if ratioA.GreaterThan(ratioB) {
		ratioA, ratioB = ratioB, ratioA
	}

	var liquidity decimal.Decimal

	if ratio.LessThanOrEqual(ratioA) {
		liquidity = uc.getLiquidityForAmount0(ratioA, ratioB, amount0)
	} else if ratio.LessThan(ratioB) {
		liquidity0 := uc.getLiquidityForAmount0(ratio, ratioB, amount0)
		liquidity1 := uc.getLiquidityForAmount1(ratioA, ratio, amount1)
		if liquidity0.LessThan(liquidity1) {
			liquidity = liquidity0
		} else {
			liquidity = liquidity1
		}
	} else {
		liquidity = uc.getLiquidityForAmount1(ratioA, ratioB, amount1)
	}

	return liquidity
}

func (uc *PositionManagerUsecase) getLiquidityForAmount0(ratioA, ratioB, amount0 decimal.Decimal) decimal.Decimal {
	// Ensure ratioA is less than or equal to ratioB
	if ratioA.GreaterThan(ratioB) {
		ratioA, ratioB = ratioB, ratioA
	}

	// intermediate = ratioA * ratioB
	intermediate := ratioA.Mul(ratioB)

	// liquidity = amount0 * intermediate / (ratioB - ratioA)
	liquidity := amount0.Mul(intermediate).Div(ratioB.Sub(ratioA))

	return liquidity
}

func (uc *PositionManagerUsecase) getLiquidityForAmount1(ratioA, ratioB, amount1 decimal.Decimal) decimal.Decimal {
	// Ensure ratioA is less than or equal to ratioB
	if ratioA.GreaterThan(ratioB) {
		ratioA, ratioB = ratioB, ratioA
	}

	// liquidity = amount1 / (ratioB - ratioA)
	liquidity := amount1.Div(ratioB.Sub(ratioA))

	return liquidity
}
