package tickmath

import (
	"errors"
	"uniswap/internal/util"

	"github.com/shopspring/decimal"
)

const (
	MinTick = -887272
	MaxTick = -MinTick

	minSqrtRatioStr = "4295128739"
	maxSqrtRatioStr = "1461446703485210103287273052203988822378723970342"
)

var (
	MinSqrtRatio, _ = decimal.NewFromString(minSqrtRatioStr)
	MaxSqrtRatio, _ = decimal.NewFromString(maxSqrtRatioStr)

	MaxUint256 = util.BitMaxNumDecimal(256)
)

func GetSqrtRatioAtTick(tick int32) (decimal.Decimal, error) {
	absTick := tick
	if tick < 0 {
		absTick = -tick
	}

	if absTick > MaxTick {
		return decimal.Decimal{}, errors.New("T")
	}

	ratio := util.NewDecimalFromHex("0x100000000000000000000000000000000")
	if absTick&0x1 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xfffcb933bd6fad37aa2d162d1a594001")),
			128)
	}
	if absTick&0x2 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xfff97272373d413259a46990580e213a")),
			128)
	}
	if absTick&0x4 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xfff2e50f5f656932ef12357cf3c7fdcc")),
			128)
	}
	if absTick&0x8 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xffe5caca7e10e4e61c3624eaa0941cd0")),
			128)
	}
	if absTick&0x10 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xffcb9843d60f6159c9db58835c926644")),
			128)
	}
	if absTick&0x20 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xff973b41fa98c081472e6896dfb254c0")),
			128)
	}
	if absTick&0x40 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xff2ea16466c96a3843ec78b326b52861")),
			128)
	}
	if absTick&0x80 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xfe5dee046a99a2a811c461f1969c3053")),
			128)
	}
	if absTick&0x100 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xfcbe86c7900a88aedcffc83b479aa3a4")),
			128)
	}
	if absTick&0x200 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xf987a7253ac413176f2b074cf7815e54")),
			128)
	}
	if absTick&0x400 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xf3392b0822b70005940c7a398e4b70f3")),
			128)
	}
	if absTick&0x800 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xe7159475a2c29b7443b29c7fa6e889d9")),
			128)
	}
	if absTick&0x1000 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xd097f3bdfd2022b8845ad8f792aa5825")),
			128)
	}
	if absTick&0x2000 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0xa9f746462d870fdf8a65dc1f90e061e5")),
			128)
	}
	if absTick&0x4000 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0x70d869a156d2a1b890bb3df62baf32f7")),
			128)
	}
	if absTick&0x8000 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0x31be135f97d08fd981231505542fcfa6")),
			128)
	}
	if absTick&0x10000 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0x9aa508b5b7a84e1c677de54f3e99bc9")),
			128)
	}
	if absTick&0x20000 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0x5d6af8dedb81196699c329225ee604")),
			128)
	}
	if absTick&0x40000 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0x2216e584f5fa1ea926041bedfe98")),
			128)
	}
	if absTick&0x80000 != 0 {
		ratio = util.DecimalRsh(
			ratio.Mul(util.NewDecimalFromHex("0x48a170391f7dc42444e8fa2")),
			128)
	}

	if tick > 0 {
		ratio = MaxUint256.DivRound(ratio, 0)
	}

	price := util.DecimalRsh(ratio, 32)
	if !ratio.Mod(decimal.NewFromInt(1 << 32)).Equal(decimal.Zero) {
		price.Add(decimal.NewFromInt(1))
	}
	return price, nil
}

func GetTickAtSqrtRatio(price decimal.Decimal) (int32, error) {
	if price.Cmp(MinSqrtRatio) < 0 || price.Cmp(MaxSqrtRatio) >= 0 {
		return 0, errors.New("price out of range")
	}

	ratio := util.DecimalLsh(price, 32)

	r := ratio
	msb := uint32(0)

	if r.Cmp(util.BitMaxNumDecimal(128)) > 0 {
		f := uint32(1) << 7
		msb |= f
		r = util.DecimalRsh(r, uint(f))
	}
	if r.Cmp(util.BitMaxNumDecimal(64)) > 0 {
		f := uint32(1) << 6
		msb |= f
		r = util.DecimalRsh(r, uint(f))
	}
	if r.Cmp(util.BitMaxNumDecimal(32)) > 0 {
		f := uint32(1) << 5
		msb |= f
		r = util.DecimalRsh(r, uint(f))
	}
	if r.Cmp(util.BitMaxNumDecimal(16)) > 0 {
		f := uint32(1) << 4
		msb |= f
		r = util.DecimalRsh(r, uint(f))
	}
	if r.Cmp(util.BitMaxNumDecimal(8)) > 0 {
		f := uint32(1) << 3
		msb |= f
		r = util.DecimalRsh(r, uint(f))
	}
	if r.Cmp(util.BitMaxNumDecimal(4)) > 0 {
		f := uint32(1) << 2
		msb |= f
		r = util.DecimalRsh(r, uint(f))
	}
	if r.Cmp(util.BitMaxNumDecimal(2)) > 0 {
		f := uint32(1) << 1
		msb |= f
		r = util.DecimalRsh(r, uint(f))
	}
	if r.Cmp(util.BitMaxNumDecimal(1)) > 0 {
		msb |= 1
	}

	if msb >= 128 {
		r = util.DecimalRsh(ratio, uint(msb)-127)
	} else {
		r = util.DecimalLsh(ratio, 127-uint(msb))
	}

	log2 := util.DecimalLsh(decimal.NewFromInt(int64(msb)-128), 64)

	for i := uint32(63); i >= 50; i-- {
		r = util.DecimalRsh(r.Mul(r), 127)
		f := util.DecimalRsh(r, 128)
		log2 = util.DecimalOr(log2, util.DecimalLsh(f, uint(i)))
		r = util.DecimalRsh(r, uint(f.BigInt().Uint64()))
	}

	d1, _ := decimal.NewFromString("255738958999603826347141")
	logSqrt10001 := log2.Mul(d1)

	d2, _ := decimal.NewFromString("3402992956809132418596140100660247210")
	d3, _ := decimal.NewFromString("291339464771989622907027621153398088495")
	tickLow := int32(util.DecimalRsh(logSqrt10001.Sub(d2), 128).IntPart())
	tickHi := int32(util.DecimalRsh(logSqrt10001.Add(d3), 128).IntPart())

	if tickLow == tickHi {
		return tickLow, nil
	}

	if d, _ := GetSqrtRatioAtTick(tickHi); d.Cmp(price) <= 0 {
		return tickHi, nil
	}
	return tickLow, nil
}
