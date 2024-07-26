package util

import (
	"github.com/shopspring/decimal"
	"math/big"
)

func BitMaxNumDecimal(bit uint) decimal.Decimal {
	oneInt := big.NewInt(1)

	return decimal.NewFromBigInt(new(big.Int).Sub(new(big.Int).Lsh(oneInt, bit), oneInt), 0)
}
