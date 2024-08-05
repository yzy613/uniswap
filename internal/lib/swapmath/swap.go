package swapmath

import (
	"github.com/shopspring/decimal"
	"uniswap/internal/lib/pricemath"
)

func ComputeSwapStep(ratioCurrent, ratioTarget, liquidity, amountRemaining decimal.Decimal, feePips uint32,
) (ratioNext, amountIn, amountOut, feeAmount decimal.Decimal, err error) {
	zeroForOne := ratioCurrent.GreaterThanOrEqual(ratioTarget)
	exactIn := amountRemaining.GreaterThanOrEqual(decimal.Zero)

	if exactIn {
		amountRemainingLessFee := amountRemaining.Mul(decimal.NewFromInt(1e6 - int64(feePips))).
			Div(decimal.NewFromInt(1e6))
		if zeroForOne {
			amountIn, err = pricemath.GetAmount0DeltaWithRound(ratioTarget, ratioCurrent, liquidity, true)
			if err != nil {
				return
			}
		} else {
			amountIn, err = pricemath.GetAmount1DeltaWithRound(ratioCurrent, ratioTarget, liquidity, true)
			if err != nil {
				return
			}
		}
		if amountRemainingLessFee.GreaterThanOrEqual(amountIn) {
			ratioNext = ratioTarget
		} else {
			ratioNext, err = pricemath.GetNextPriceFromInput(
				ratioCurrent,
				liquidity,
				amountRemainingLessFee,
				zeroForOne)
			if err != nil {
				return
			}
		}
	} else {
		if zeroForOne {
			amountOut, err = pricemath.GetAmount1DeltaWithRound(ratioTarget, ratioCurrent, liquidity, false)
			if err != nil {
				return
			}
		} else {
			amountOut, err = pricemath.GetAmount0DeltaWithRound(ratioCurrent, ratioTarget, liquidity, false)
			if err != nil {
				return
			}
			amountRemainingNeg := amountRemaining.Neg()
			if amountRemainingNeg.GreaterThanOrEqual(amountOut) {
				ratioNext = ratioTarget
			} else {
				ratioNext, err = pricemath.GetNextPriceFromOutput(
					ratioCurrent,
					liquidity,
					amountRemainingNeg,
					zeroForOne)
				if err != nil {
					return
				}
			}
		}
	}

	isMax := ratioTarget.Equal(ratioNext)

	if isMax {
		if zeroForOne {
			if !exactIn {
				amountIn, err = pricemath.GetAmount0DeltaWithRound(ratioNext, ratioCurrent, liquidity, true)
				if err != nil {
					return
				}
			} else {
				amountOut, err = pricemath.GetAmount1DeltaWithRound(ratioNext, ratioCurrent, liquidity, false)
				if err != nil {
					return
				}
			}
		} else {
			if !exactIn {
				amountIn, err = pricemath.GetAmount1DeltaWithRound(ratioCurrent, ratioNext, liquidity, true)
				if err != nil {
					return
				}
			} else {
				amountOut, err = pricemath.GetAmount0DeltaWithRound(ratioCurrent, ratioNext, liquidity, false)
				if err != nil {
					return
				}
			}
		}
	}

	if !exactIn && amountOut.GreaterThan(amountRemaining.Neg()) {
		amountOut = amountRemaining.Neg()
	}

	if exactIn && !ratioNext.Equal(ratioTarget) {
		feeAmount = amountRemaining.Sub(amountIn)
	} else {
		feePipsD := decimal.NewFromInt(int64(feePips))
		feeAmount = amountIn.Mul(feePipsD).Div(decimal.NewFromInt(1e6).RoundUp(0).Sub(feePipsD))
	}

	return
}
