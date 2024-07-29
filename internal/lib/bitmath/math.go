package bitmath

import (
	"math/big"
	"uniswap/internal/util"
)

// MostSignificantBit returns the index of the most significant bit of the number,
// where the least significant bit is at index 0 and the most significant bit is at index 255
func MostSignificantBit(x *big.Int) uint8 {
	if x.Cmp(big.NewInt(0)) <= 0 {
		panic("value must be greater than 0")
	}

	var r uint8
	t := new(big.Int).Set(x)

	if t.Cmp(new(big.Int).Lsh(big.NewInt(1), 128)) >= 0 {
		t.Rsh(t, 128)
		r += 128
	}
	if t.Cmp(new(big.Int).Lsh(big.NewInt(1), 64)) >= 0 {
		t.Rsh(t, 64)
		r += 64
	}
	if t.Cmp(new(big.Int).Lsh(big.NewInt(1), 32)) >= 0 {
		t.Rsh(t, 32)
		r += 32
	}
	if t.Cmp(new(big.Int).Lsh(big.NewInt(1), 16)) >= 0 {
		t.Rsh(t, 16)
		r += 16
	}
	if t.Cmp(new(big.Int).Lsh(big.NewInt(1), 8)) >= 0 {
		t.Rsh(t, 8)
		r += 8
	}
	if t.Cmp(new(big.Int).Lsh(big.NewInt(1), 4)) >= 0 {
		t.Rsh(t, 4)
		r += 4
	}
	if t.Cmp(new(big.Int).Lsh(big.NewInt(1), 2)) >= 0 {
		t.Rsh(t, 2)
		r += 2
	}
	if t.Cmp(new(big.Int).Lsh(big.NewInt(1), 1)) >= 0 {
		r += 1
	}
	return r
}

// LeastSignificantBit returns the index of the least significant bit of the number,
// where the least significant bit is at index 0 and the most significant bit is at index 255
func LeastSignificantBit(x *big.Int) uint8 {
	bigZero := big.NewInt(0)

	if x.Cmp(bigZero) <= 0 {
		panic("value must be greater than 0")
	}

	var r uint8 = 255
	t := new(big.Int).Set(x)

	if t.And(t, util.BitMaxNumDecimal(128).BigInt()).Cmp(bigZero) > 0 {
		r -= 128
	} else {
		t.Rsh(t, 128)
	}
	if t.And(t, new(big.Int).SetUint64(^uint64(0))).Cmp(bigZero) > 0 {
		r -= 64
	} else {
		t.Rsh(t, 64)
	}
	if t.And(t, new(big.Int).SetUint64(uint64(^uint32(0)))).Cmp(bigZero) > 0 {
		r -= 32
	} else {
		t.Rsh(t, 32)
	}
	if t.And(t, new(big.Int).SetUint64(uint64(^uint16(0)))).Cmp(bigZero) > 0 {
		r -= 16
	} else {
		t.Rsh(t, 16)
	}
	if t.And(t, new(big.Int).SetUint64(uint64(^uint8(0)))).Cmp(bigZero) > 0 {
		r -= 8
	} else {
		t.Rsh(t, 8)
	}
	if t.And(t, big.NewInt(15)).Cmp(bigZero) > 0 {
		r -= 4
	} else {
		t.Rsh(t, 4)
	}
	if t.And(t, big.NewInt(3)).Cmp(bigZero) > 0 {
		r -= 2
	} else {
		t.Rsh(t, 2)
	}
	if t.And(t, big.NewInt(1)).Cmp(bigZero) > 0 {
		r -= 1
	}
	return r
}
