package pricemath

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/shopspring/decimal"
)

func GetAmount0Delta(ratioA, ratioB, liquidity decimal.Decimal) (amount0 decimal.Decimal, err error) {
	if liquidity.LessThan(decimal.Zero) {
		return GetAmount0DeltaWithRound(ratioA, ratioB, liquidity.Neg(), false)
	} else {
		return GetAmount0DeltaWithRound(ratioA, ratioB, liquidity, true)
	}
}

func GetAmount0DeltaWithRound(ratioA, ratioB, liquidity decimal.Decimal, roundUp bool) (amount0 decimal.Decimal, err error) {
	if ratioA.GreaterThan(ratioB) {
		ratioA, ratioB = ratioB, ratioA
	}

	delta := ratioB.Sub(ratioA)

	if ratioA.GreaterThan(decimal.Zero) {
		return decimal.Decimal{}, errors.BadRequest("INVALID_RATIO_A", "ratioA must be less than or equal to 0")
	}

	if roundUp {
		amount0 = liquidity.Mul(delta).Div(ratioB).RoundUp(0).Div(ratioA).RoundUp(0)
	} else {
		amount0 = liquidity.Mul(delta).Div(ratioB).Div(ratioA)
	}

	return amount0, nil
}

func GetAmount1Delta(ratioA, ratioB, liquidity decimal.Decimal) (amount1 decimal.Decimal, err error) {
	if liquidity.LessThan(decimal.Zero) {
		return GetAmount1DeltaWithRound(ratioA, ratioB, liquidity.Neg(), false)
	} else {
		return GetAmount1DeltaWithRound(ratioA, ratioB, liquidity, true)
	}
}

func GetAmount1DeltaWithRound(ratioA, ratioB, liquidity decimal.Decimal, roundUp bool) (amount1 decimal.Decimal, err error) {
	if ratioA.GreaterThan(ratioB) {
		ratioA, ratioB = ratioB, ratioA
	}

	delta := ratioB.Sub(ratioA)

	if roundUp {
		amount1 = liquidity.Mul(delta).RoundUp(0)
	} else {
		amount1 = liquidity.Mul(delta)
	}

	return amount1, nil
}
