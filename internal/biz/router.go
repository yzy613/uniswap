package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
)

type Path struct {
	TokenA string
	TokenB string
	Fee    uint32
}

type SwapCallbackData struct {
	Path  Path
	Payer string
}

type RouterUsecase struct {
	log *log.Helper

	pool *PoolUsecase
}

func NewRouterUsecase(logger log.Logger, pool *PoolUsecase) *RouterUsecase {
	return &RouterUsecase{log: log.NewHelper(logger), pool: pool}
}

func (uc *RouterUsecase) ExactInputInternal(
	amountIn decimal.Decimal,
	recipient string,
	priceLimit decimal.Decimal,
	data SwapCallbackData,
) (amountOut decimal.Decimal, err error) {
	if recipient == "" || recipient == "0" {
		recipient = data.Payer
	}

	tokenIn, tokenOut, fee := data.Path.TokenA, data.Path.TokenB, data.Path.Fee

	zeroForOne := tokenIn < tokenOut

	pool, err := uc.pool.GetPool(tokenIn, tokenOut, fee)
	if err != nil {
		return decimal.Decimal{}, err
	}
	amount0, amount1, err := uc.pool.Swap(
		*pool,
		recipient,
		zeroForOne,
		amountIn,
		priceLimit,
		data,
	)
	if err != nil {
		return decimal.Decimal{}, err
	}

	if zeroForOne {
		return amount1.Neg(), nil
	} else {
		return amount0.Neg(), nil
	}
}

func (uc *RouterUsecase) ExactOutputInternal(
	amountOut decimal.Decimal,
	recipient string,
	priceLimit decimal.Decimal,
	data SwapCallbackData,
) (amountIn decimal.Decimal, err error) {
	if recipient == "" || recipient == "0" {
		recipient = data.Payer
	}

	tokenOut, tokenIn, fee := data.Path.TokenA, data.Path.TokenB, data.Path.Fee

	zeroForOne := tokenIn < tokenOut

	pool, err := uc.pool.GetPool(tokenIn, tokenOut, fee)
	if err != nil {
		return decimal.Decimal{}, err
	}
	amount0, amount1, err := uc.pool.Swap(
		*pool,
		recipient,
		zeroForOne,
		amountOut.Neg(),
		priceLimit,
		data,
	)
	if err != nil {
		return decimal.Decimal{}, err
	}

	var amountOutReceived decimal.Decimal
	if zeroForOne {
		amountIn, amountOutReceived = amount0, amount1.Neg()
	} else {
		amountIn, amountOutReceived = amount1, amount0.Neg()
	}

	if priceLimit.IsZero() && !amountOutReceived.Equal(amountOut) {
		return decimal.Decimal{}, errors.InternalServer("INVALID_AMOUNT_OUT", "invalid amount out")
	}

	return amountIn, nil
}
