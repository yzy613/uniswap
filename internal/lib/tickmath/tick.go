package tickmath

import (
	"errors"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
)

// 精度问题
// MinTick = -887272
const (
	MinTick = -886463
	MaxTick = -MinTick
)

func GetRatioAtTick(tick int32) (decimal.Decimal, error) {
	if tick > MaxTick || tick < MinTick {
		return decimal.Decimal{}, errors.New("T")
	}

	ratio, _ := decimal.NewFromString("1.0001")
	ratio = ratio.Pow(decimal.NewFromInt(int64(tick)))

	return ratio, nil
}

func GetTickAtRatio(ratio decimal.Decimal) (int32, error) {
	ratioStr := ratio.String()
	ratioFloat, err := strconv.ParseFloat(ratioStr, 64)
	if err != nil {
		return 0, errors.New("invalid ratio")
	}

	if ratioFloat <= 0 {
		return 0, errors.New("ratio must be greater than 0")
	}
	tick := int32(math.Round(math.Log(ratioFloat) / math.Log(1.0001)))

	if tick > MaxTick || tick < -MaxTick {
		return 0, errors.New("tick out of bounds")
	}

	return tick, nil
}
