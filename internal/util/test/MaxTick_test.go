package feat

import (
	"github.com/shopspring/decimal"
	"math/big"
	"testing"
)

var (
	// TODO: 精度问题
	MinTick int32 = -887272
	MaxTick       = -MinTick
)

func delta(d int32) {
	MinTick += d
	MaxTick -= d
}

func BitMaxNumDecimal(bit uint) decimal.Decimal {
	oneInt := big.NewInt(1)

	return decimal.NewFromBigInt(new(big.Int).Sub(new(big.Int).Lsh(oneInt, bit), oneInt), 0)
}

func tickSpacingToMaxLiquidityPerTick(tickSpacing int32) decimal.Decimal {
	minTick := MinTick / tickSpacing * tickSpacing
	maxTick := MaxTick / tickSpacing * tickSpacing
	numTicks := uint32((maxTick-minTick)/tickSpacing) + 1

	return BitMaxNumDecimal(128).DivRound(decimal.NewFromInt(int64(numTicks)), 0)
}

func maxLiquidityPerTickToTickSpacing(maxLiquidityPerTick decimal.Decimal) int32 {
	// 定义一些常量
	minTick := MinTick
	maxTick := MaxTick

	// 计算numTicks
	numTicks := BitMaxNumDecimal(128).DivRound(maxLiquidityPerTick, 0).IntPart()

	// 使用numTicks计算tickSpacing
	tickSpacing := (maxTick - minTick) / int32(numTicks-1)

	return tickSpacing
}

// MaxTick_test.go:62: MinTick: -886463, MaxTick: 886463
// MaxTick_test.go:63: want 972 got 973
// MaxTick_test.go:64: maxValidTickSpacing: 971
func TestPoolUsecase_tickSpacingToMaxLiquidityPerTick(t *testing.T) {
	maxValidTickSpacing := int32(1)

	lastTickSpacing := int32(1)
	failedCount := 0

	for {
		tickSpacing := int32(1)
		for ; tickSpacing < 1<<31-1; tickSpacing++ {
			deci := tickSpacingToMaxLiquidityPerTick(tickSpacing)

			if val := maxLiquidityPerTickToTickSpacing(deci); val != tickSpacing {
				failedCount++

				if tickSpacing > maxValidTickSpacing && tickSpacing > lastTickSpacing {
					lastTickSpacing = tickSpacing

					t.Logf("MinTick: %v, MaxTick: %v", MinTick, MaxTick)
					t.Errorf("want %v got %v", tickSpacing, val)
					t.Logf("maxValidTickSpacing: %v", maxValidTickSpacing)
					t.Log()
				}

				break
			}

			if tickSpacing > maxValidTickSpacing {
				maxValidTickSpacing = tickSpacing
			}
		}

		if failedCount > 1e4 || MinTick == -1 || MaxTick == 1 {
			break
		}

		delta(1)
	}
}
