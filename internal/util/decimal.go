package util

import (
	"github.com/shopspring/decimal"
	"math/big"
)

func NewDecimalFromHex(hex string) decimal.Decimal {
	i, _ := new(big.Int).SetString(hex, 0)
	return decimal.NewFromBigInt(i, 0)
}

func BitMaxNumDecimal(bit uint) decimal.Decimal {
	oneInt := big.NewInt(1)

	return decimal.NewFromBigInt(new(big.Int).Sub(new(big.Int).Lsh(oneInt, bit), oneInt), 0)
}

// decimal Bitwise operations

func DecimalAnd(d1, d2 decimal.Decimal) decimal.Decimal {
	i1, i2 := d1.BigInt(), d2.BigInt()
	return decimal.NewFromBigInt(i1.And(i1, i2), 0)
}

func DecimalOr(d1, d2 decimal.Decimal) decimal.Decimal {
	i1, i2 := d1.BigInt(), d2.BigInt()
	return decimal.NewFromBigInt(i1.Or(i1, i2), 0)
}

func DecimalXor(d1, d2 decimal.Decimal) decimal.Decimal {
	i1, i2 := d1.BigInt(), d2.BigInt()
	return decimal.NewFromBigInt(i1.Xor(i1, i2), 0)
}

func DecimalNot(d decimal.Decimal) decimal.Decimal {
	i := d.BigInt()
	return decimal.NewFromBigInt(i.Not(i), 0)
}

func DecimalLsh(d decimal.Decimal, bit uint) decimal.Decimal {
	i := d.BigInt()
	i.Lsh(i, bit)
	return decimal.NewFromBigInt(i, 0)
}

func DecimalRsh(d decimal.Decimal, bit uint) decimal.Decimal {
	i := d.BigInt()
	i.Rsh(i, bit)
	return decimal.NewFromBigInt(i, 0)
}
