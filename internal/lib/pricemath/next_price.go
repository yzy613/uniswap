package pricemath

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/shopspring/decimal"
)

func GetNextPriceFromInput(price, liquidity, amountIn decimal.Decimal, zeroForOne bool) (decimal.Decimal, error) {
	if price.LessThanOrEqual(decimal.Zero) {
		return decimal.Decimal{}, errors.BadRequest("INVALID_PRICE", "price must be greater than 0")
	}
	if liquidity.LessThanOrEqual(decimal.Zero) {
		return decimal.Decimal{}, errors.BadRequest("INVALID_LIQUIDITY", "liquidity must be greater than 0")
	}

	if zeroForOne {
		return GetNextPriceFromAmount0(price, liquidity, amountIn, true)
	} else {
		return GetNextPriceFromAmount1(price, liquidity, amountIn, true)
	}
}

func GetNextPriceFromOutput(price, liquidity, amountOut decimal.Decimal, zeroForOne bool) (decimal.Decimal, error) {
	if price.LessThanOrEqual(decimal.Zero) {
		return decimal.Decimal{}, errors.BadRequest("INVALID_PRICE", "price must be greater than 0")
	}
	if liquidity.LessThanOrEqual(decimal.Zero) {
		return decimal.Decimal{}, errors.BadRequest("INVALID_LIQUIDITY", "liquidity must be greater than 0")
	}

	if zeroForOne {
		return GetNextPriceFromAmount1(price, liquidity, amountOut, false)
	} else {
		return GetNextPriceFromAmount0(price, liquidity, amountOut, false)
	}
}

func GetNextPriceFromAmount0(price, liquidity, amount decimal.Decimal, add bool) (decimal.Decimal, error) {
	if amount.IsZero() {
		return price, nil
	}

	var newPrice decimal.Decimal

	if add {
		newPrice = liquidity.Mul(price).Div(liquidity.Add(amount.Mul(price)))
	} else {
		if amount.Mul(price).GreaterThan(liquidity) {
			return decimal.Decimal{}, errors.BadRequest("INVALID_AMOUNT", "amount is too large")
		}
		newPrice = liquidity.Mul(price).Div(liquidity.Sub(amount.Mul(price)))
	}

	return newPrice, nil
}

func GetNextPriceFromAmount1(price, liquidity, amount decimal.Decimal, add bool) (decimal.Decimal, error) {
	if amount.IsZero() {
		return price, nil
	}

	var newPrice decimal.Decimal

	if add {
		newPrice = price.Add(amount.Div(liquidity))
	} else {
		quotient := amount.Add(liquidity).Sub(decimal.NewFromInt(1)).Sub(liquidity)
		if price.LessThanOrEqual(quotient) {
			return decimal.Decimal{}, errors.BadRequest("INVALID_AMOUNT", "amount is too large")
		}
		newPrice = price.Sub(quotient)
	}

	return newPrice, nil
}
