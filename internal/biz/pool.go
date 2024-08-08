package biz

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/shopspring/decimal"
	"time"
	"uniswap/internal/lib/pricemath"
	"uniswap/internal/lib/swapmath"
	"uniswap/internal/lib/tickmath"
	"uniswap/internal/model/entity"
	"uniswap/internal/util"
)

type Slot0 struct {
	entity.Slot0
}

type Pool struct {
	entity.Pool
	slot0 *Slot0
}

type Position struct {
	entity.Position
}

type PoolRepo interface {
	GetPool(token0, token1 string, fee uint32) (*Pool, error)
	CreatePool(token0, token1 string, fee uint32, tickSpacing int32,
		tickSpacingToMaxLiquidityPerTick decimal.Decimal) (*Pool, error)
	FeeAmountTickSpacing(fee uint32) (tickSpacing int32)
	GetPosition(poolId int64, owner string, tickLower, tickUpper int32) (*Position, error)
	UpdatePosition(position Position,
		liquidityDelta, feeGrowthInside0, feeGrowthInside1 decimal.Decimal) error
	GetSlot0(poolId int64) (*Slot0, error)
	SaveSlot0(slot0 Slot0) error
	TryLockSlot0(poolId int64) error
	UnlockSlot0(poolId int64) error
	GetFeeGrowthGlobal(poolId int64) (feeGrowthGlobal0, feeGrowthGlobal1 decimal.Decimal, err error)
	SaveFeeGrowthGlobal0(poolId int64, feeGrowthGlobal0 decimal.Decimal) error
	SaveFeeGrowthGlobal1(poolId int64, feeGrowthGlobal1 decimal.Decimal) error
	GetProtocolFee(poolId int64) (token0, token1 decimal.Decimal, err error)
	SaveProtocolFeeToken0(poolId int64, token0 decimal.Decimal) error
	SaveProtocolFeeToken1(poolId int64, token1 decimal.Decimal) error
}

type PoolUsecase struct {
	repo PoolRepo
	log  *log.Helper

	tick        *TickUsecase
	observation *ObservationUsecase
	tickBitmap  *TickBitmapUsecase
	liquidity   *LiquidityUsecase
}

func NewPoolUsecase(repo PoolRepo, logger log.Logger, tick *TickUsecase,
	observation *ObservationUsecase, tickBitmap *TickBitmapUsecase, liquidity *LiquidityUsecase) *PoolUsecase {
	return &PoolUsecase{
		repo:        repo,
		log:         log.NewHelper(logger),
		tick:        tick,
		observation: observation,
		tickBitmap:  tickBitmap,
		liquidity:   liquidity,
	}
}

func (uc *PoolUsecase) GetPool(token0, token1 string, fee uint32) (*Pool, error) {
	return uc.repo.GetPool(token0, token1, fee)
}

func (uc *PoolUsecase) initialize(pool Pool, price decimal.Decimal) error {
	slot0, err := uc.repo.GetSlot0(pool.Id)
	if err != nil {
		return err
	}
	if slot0 != nil && !slot0.Price.IsZero() {
		return errors.BadRequest("SLOT0_EXISTS", "slot0 already exists")
	}
	tick, err := tickmath.GetTickAtRatio(price)
	if err != nil {
		return err
	}
	// TODO: _blockTimestamp()
	cardinality, cardinalityNext, err := uc.observation.Initialize(pool, uint32(time.Now().Unix()))
	if err != nil {
		return err
	}

	return uc.repo.SaveSlot0(Slot0{
		Slot0: entity.Slot0{
			PoolId:                     pool.Id,
			Price:                      price,
			CurrentTick:                int(tick),
			ObservationIndex:           0,
			ObservationCardinality:     int(cardinality),
			ObservationCardinalityNext: int(cardinalityNext),
			FeeProtocol0:               0,
			FeeProtocol1:               0,
			Unlocked:                   1,
		},
	})
}

func (uc *PoolUsecase) CreatePool(token0, token1 string, fee uint32) (*Pool, error) {
	if token0 == token1 {
		return nil, errors.BadRequest("INVALID_TOKEN", "token0 and token1 must be different")
	}
	if token0 > token1 {
		token0, token1 = token1, token0
	}
	if token0 == "" || token0 == "0" {
		return nil, errors.BadRequest("INVALID_TOKEN", "token0 is invalid")
	}
	tickSpacing := uc.repo.FeeAmountTickSpacing(fee)
	if tickSpacing == 0 {
		return nil, errors.BadRequest("INVALID_FEE", "fee is invalid")
	}
	if p, e := uc.repo.GetPool(token0, token1, fee); e == nil && p != nil {
		return nil, errors.BadRequest("POOL_EXISTS", "pool already exists")
	}

	return uc.repo.CreatePool(token0, token1, fee, tickSpacing, uc.tickSpacingToMaxLiquidityPerTick(tickSpacing))
}

func (uc *PoolUsecase) tickSpacingToMaxLiquidityPerTick(tickSpacing int32) decimal.Decimal {
	minTick := tickmath.MinTick / tickSpacing * tickSpacing
	maxTick := tickmath.MaxTick / tickSpacing * tickSpacing
	numTicks := uint32((maxTick-minTick)/tickSpacing) + 1

	return util.BitMaxNumDecimal(128).DivRound(decimal.NewFromInt(int64(numTicks)), 0)
}

func (uc *PoolUsecase) checkTick(tickLower, tickUpper int32) error {
	if tickLower >= tickUpper {
		return errors.BadRequest("INVALID_TICK", "tickLower must be less than tickUpper")
	}
	if tickLower < tickmath.MinTick || tickUpper > tickmath.MaxTick {
		return errors.BadRequest("INVALID_TICK", "tickLower and tickUpper must be within the range")
	}
	return nil
}

func (uc *PoolUsecase) Mint(pool Pool, tickLower, tickUpper int32, amount decimal.Decimal,
) (amount0, amount1 decimal.Decimal, err error) {
	// TODO: Mint
	return decimal.Zero, decimal.Zero, nil
}

func (uc *PoolUsecase) modifyPosition(poolId int64, owner string, tickLower, tickUpper int32, liquidityDelta decimal.Decimal,
) (position *Position, amount0, amount1 decimal.Decimal, err error) {
	err = uc.checkTick(tickLower, tickUpper)
	if err != nil {
		return
	}

	slot0, err := uc.repo.GetSlot0(poolId)
	if err != nil {
		return nil, decimal.Zero, decimal.Zero, err
	}

	position, err = uc.updatePosition(
		Pool{Pool: entity.Pool{Id: poolId}, slot0: slot0},
		owner,
		tickLower,
		tickUpper,
		liquidityDelta,
		int32(slot0.CurrentTick),
	)
	if err != nil {
		return nil, decimal.Zero, decimal.Zero, err
	}

	if !liquidityDelta.IsZero() {
		if slot0.CurrentTick < int(tickLower) {
			amount0 = liquidityDelta
		} else if slot0.CurrentTick < int(tickUpper) {
			liquidity, err := uc.liquidity.Get(poolId)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}

			indexUpdated, cardinalityUpdated, err := uc.observation.Write(
				poolId,
				uint16(slot0.ObservationIndex),
				uint32(time.Now().Unix()),
				int32(slot0.CurrentTick),
				liquidity,
				uint16(slot0.ObservationCardinality),
				uint16(slot0.ObservationCardinalityNext),
			)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}
			slot0.ObservationIndex, slot0.ObservationCardinality = int(indexUpdated), int(cardinalityUpdated)
			err = uc.repo.SaveSlot0(*slot0)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}

			ratio, err := tickmath.GetRatioAtTick(tickUpper)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}
			amount0, err = pricemath.GetAmount0Delta(slot0.Price, ratio, liquidityDelta)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}

			ratio, err = tickmath.GetRatioAtTick(tickLower)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}
			amount1, err = pricemath.GetAmount1Delta(ratio, slot0.Price, liquidityDelta)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}

			err = uc.liquidity.Save(poolId, liquidity.Add(liquidityDelta))
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}
		} else {
			ratioLower, err := tickmath.GetRatioAtTick(tickLower)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}
			ratioUpper, err := tickmath.GetRatioAtTick(tickUpper)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}

			amount1, err = pricemath.GetAmount1Delta(ratioLower, ratioUpper, liquidityDelta)
			if err != nil {
				return nil, decimal.Decimal{}, decimal.Decimal{}, err
			}
		}
	}

	return
}

func (uc *PoolUsecase) updatePosition(pool Pool, owner string, tickLower, tickUpper int32, liquidityDelta decimal.Decimal,
	tick int32) (*Position, error) {
	position, err := uc.repo.GetPosition(pool.Id, owner, tickLower, tickUpper)
	if err != nil {
		return nil, err
	}

	feeGrowthGlobal0, feeGrowthGlobal1, err := uc.repo.GetFeeGrowthGlobal(pool.Id)
	if err != nil {
		return nil, err
	}

	var flippedLower, flippedUpper bool
	if !liquidityDelta.IsZero() {
		// TODO: _blockTimestamp()
		time_ := uint32(time.Now().Unix())
		tickCumulative, secondsPerLiquidityCumulative, err := uc.observation.ObserveSingle(
			pool.Id,
			time_,
			0,
			tick,
			uint16(pool.slot0.ObservationIndex),
			liquidityDelta,
			uint16(pool.slot0.ObservationCardinality),
		)
		if err != nil {
			return nil, err
		}
		flippedLower, err = uc.tick.Update(
			pool.Id,
			tick,
			tickLower,
			liquidityDelta,
			feeGrowthGlobal0,
			feeGrowthGlobal1,
			secondsPerLiquidityCumulative,
			tickCumulative,
			time_,
			false,
			pool.MaxLiquidityPerTick,
		)
		if err != nil {
			return nil, err
		}
		flippedUpper, err = uc.tick.Update(
			pool.Id,
			tick,
			tickUpper,
			liquidityDelta,
			feeGrowthGlobal0,
			feeGrowthGlobal1,
			secondsPerLiquidityCumulative,
			tickCumulative,
			time_,
			true,
			pool.MaxLiquidityPerTick,
		)
		if err != nil {
			return nil, err
		}

		if flippedLower {
			err = uc.tickBitmap.FlipTick(pool.Id, tickLower, int32(pool.TickSpacing))
			if err != nil {
				return nil, err
			}
		}
		if flippedUpper {
			err = uc.tickBitmap.FlipTick(pool.Id, tickUpper, int32(pool.TickSpacing))
			if err != nil {
				return nil, err
			}
		}
	}

	feeGrowthInside0, feeGrowthInside1, err := uc.tick.GetFeeGrowthInside(
		pool.Id,
		tickLower,
		tickUpper,
		tick,
		feeGrowthGlobal0,
		feeGrowthGlobal1,
	)
	if err != nil {
		return nil, err
	}

	err = uc.repo.UpdatePosition(*position, liquidityDelta, feeGrowthInside0, feeGrowthInside1)
	if err != nil {
		return nil, err
	}

	if liquidityDelta.LessThan(decimal.Zero) {
		if flippedLower {
			err = uc.tick.Clear(pool.Id, tickLower)
			return nil, err
		}
		if flippedUpper {
			err = uc.tick.Clear(pool.Id, tickUpper)
			return nil, err
		}
	}

	return position, nil
}

func (uc *PoolUsecase) Swap(
	pool Pool,
	recipient string,
	zeroForOne bool,
	amountSpecified,
	priceLimit decimal.Decimal,
	data SwapCallbackData,
) (amount0, amount1 decimal.Decimal, err error) {
	if amountSpecified.IsZero() {
		return decimal.Decimal{}, decimal.Decimal{},
			errors.BadRequest("INVALID_AMOUNT", "amountSpecified is zero")
	}

	err = uc.repo.TryLockSlot0(pool.Id)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}
	defer uc.repo.UnlockSlot0(pool.Id)

	slot0Start, err := uc.repo.GetSlot0(pool.Id)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}

	if zeroForOne {
		if priceLimit.GreaterThanOrEqual(slot0Start.Price) {
			return decimal.Decimal{}, decimal.Decimal{},
				errors.BadRequest("INVALID_PRICE", "priceLimit is invalid")
		}
	} else {
		if priceLimit.LessThanOrEqual(slot0Start.Price) {
			return decimal.Decimal{}, decimal.Decimal{},
				errors.BadRequest("INVALID_PRICE", "priceLimit is invalid")
		}
	}

	cache := struct {
		feeProtocol                   uint8
		liquidityStart                decimal.Decimal
		blockTimestamp                uint32
		tickCumulative                int64
		secondsPerLiquidityCumulative decimal.Decimal
		computedLatestObservation     bool
	}{}
	cache.liquidityStart, err = uc.liquidity.Get(pool.Id)
	if err != nil {
		return decimal.Decimal{}, decimal.Decimal{}, err
	}
	cache.blockTimestamp = uint32(time.Now().Unix())
	if zeroForOne {
		cache.feeProtocol = uint8(slot0Start.FeeProtocol0)
	} else {
		cache.feeProtocol = uint8(slot0Start.FeeProtocol1)
	}

	exactInput := amountSpecified.GreaterThan(decimal.Zero)

	state := struct {
		amountSpecifiedRemaining decimal.Decimal
		amountCalculated         decimal.Decimal
		price                    decimal.Decimal
		tick                     int32
		feeGrowthGlobal          decimal.Decimal
		protocolFee              decimal.Decimal
		liquidity                decimal.Decimal
	}{}
	state.amountSpecifiedRemaining = amountSpecified
	state.price = slot0Start.Price
	state.tick = int32(slot0Start.CurrentTick)
	if zeroForOne {
		state.feeGrowthGlobal, _, err = uc.repo.GetFeeGrowthGlobal(pool.Id)
		if err != nil {
			return decimal.Decimal{}, decimal.Decimal{}, err
		}
	} else {
		_, state.feeGrowthGlobal, err = uc.repo.GetFeeGrowthGlobal(pool.Id)
		if err != nil {
			return decimal.Decimal{}, decimal.Decimal{}, err
		}
	}
	state.liquidity = cache.liquidityStart

	for !state.amountSpecifiedRemaining.IsZero() && !state.price.Equal(priceLimit) {
		step := struct {
			priceStart  decimal.Decimal
			tickNext    int32
			initialized bool
			priceNext   decimal.Decimal
			amountIn    decimal.Decimal
			amountOut   decimal.Decimal
			feeAmount   decimal.Decimal
		}{}

		step.priceStart = state.price

		step.tickNext, step.initialized, err = uc.tickBitmap.NextInitializedTickWithinOneWord(
			pool.Id,
			state.tick,
			int32(pool.TickSpacing),
			zeroForOne,
		)

		if step.tickNext < tickmath.MinTick {
			step.tickNext = tickmath.MinTick
		} else if step.tickNext > tickmath.MaxTick {
			step.tickNext = tickmath.MaxTick
		}

		step.priceNext, err = tickmath.GetRatioAtTick(step.tickNext)
		if err != nil {
			return decimal.Decimal{}, decimal.Decimal{}, err
		}

		state.price, step.amountIn, step.amountOut, step.feeAmount, err = swapmath.ComputeSwapStep(
			state.price,
			func() decimal.Decimal {
				if zeroForOne {
					if step.priceNext.LessThan(priceLimit) {
						return priceLimit
					}
					return step.priceNext
				}
				if step.priceNext.GreaterThan(priceLimit) {
					return priceLimit
				}
				return step.priceNext
			}(),
			state.liquidity,
			state.amountSpecifiedRemaining,
			uint32(pool.Fee),
		)

		if exactInput {
			state.amountSpecifiedRemaining =
				state.amountSpecifiedRemaining.Sub(step.amountIn.Add(step.feeAmount))
			state.amountCalculated = state.amountCalculated.Sub(step.amountOut)
		} else {
			state.amountSpecifiedRemaining = state.amountSpecifiedRemaining.Add(step.amountOut)
			state.amountCalculated = state.amountCalculated.Add(step.amountIn.Add(step.feeAmount))
		}

		if cache.feeProtocol > 0 {
			delta := step.feeAmount.Div(decimal.NewFromInt(int64(cache.feeProtocol)))
			step.feeAmount = step.feeAmount.Sub(delta)
			state.protocolFee = state.protocolFee.Add(delta)
		}

		if state.price.Equal(step.priceNext) {
			if step.initialized {
				if !cache.computedLatestObservation {
					cache.tickCumulative, cache.secondsPerLiquidityCumulative, err = uc.observation.ObserveSingle(
						pool.Id,
						cache.blockTimestamp,
						0,
						int32(slot0Start.CurrentTick),
						uint16(slot0Start.ObservationIndex),
						cache.liquidityStart,
						uint16(slot0Start.ObservationCardinality),
					)
					if err != nil {
						return
					}
					cache.computedLatestObservation = true
				}
				var liquidityNet decimal.Decimal
				var feeGrowthGlobal0, feeGrowthGlobal1 decimal.Decimal
				feeGrowthGlobal0, feeGrowthGlobal1, err = uc.repo.GetFeeGrowthGlobal(pool.Id)
				liquidityNet, err = uc.tick.Cross(
					pool.Id,
					step.tickNext,
					func() decimal.Decimal {
						if zeroForOne {
							return state.feeGrowthGlobal
						}
						return feeGrowthGlobal0
					}(),
					func() decimal.Decimal {
						if zeroForOne {
							return feeGrowthGlobal1
						}
						return state.feeGrowthGlobal
					}(),
					cache.secondsPerLiquidityCumulative,
					cache.tickCumulative,
					cache.blockTimestamp,
				)
				if err != nil {
					return
				}
				if zeroForOne {
					liquidityNet = liquidityNet.Neg()
				}

				state.liquidity = state.liquidity.Add(liquidityNet)
			}

			if zeroForOne {
				state.tick = step.tickNext - 1
			} else {
				state.tick = step.tickNext
			}
		} else if !state.price.Equal(step.priceStart) {
			state.tick, err = tickmath.GetTickAtRatio(state.price)
			if err != nil {
				return
			}
		}
	}

	if state.tick != int32(slot0Start.CurrentTick) {
		var observationIndex, observationCardinality uint16
		observationIndex, observationCardinality, err = uc.observation.Write(
			pool.Id,
			uint16(slot0Start.ObservationIndex),
			cache.blockTimestamp,
			int32(slot0Start.CurrentTick),
			cache.liquidityStart,
			uint16(slot0Start.ObservationCardinality),
			uint16(slot0Start.ObservationCardinalityNext),
		)
		if err != nil {
			return
		}
		err = uc.repo.SaveSlot0(Slot0{
			Slot0: entity.Slot0{
				PoolId:                 pool.Id,
				Price:                  state.price,
				CurrentTick:            int(state.tick),
				ObservationIndex:       int(observationIndex),
				ObservationCardinality: int(observationCardinality),
			},
		})
		if err != nil {
			return
		}
	} else {
		err = uc.repo.SaveSlot0(Slot0{
			Slot0: entity.Slot0{
				PoolId: pool.Id,
				Price:  state.price,
			},
		})
		if err != nil {
			return
		}
	}

	if !cache.liquidityStart.Equal(state.liquidity) {
		err = uc.liquidity.Save(pool.Id, state.liquidity)
		if err != nil {
			return
		}
	}

	if zeroForOne {
		err = uc.repo.SaveFeeGrowthGlobal0(pool.Id, state.feeGrowthGlobal)
		if err != nil {
			return
		}
		if state.protocolFee.GreaterThan(decimal.Zero) {
			var token0 decimal.Decimal
			token0, _, err = uc.repo.GetProtocolFee(pool.Id)
			if err != nil {
				return
			}
			err = uc.repo.SaveProtocolFeeToken0(pool.Id, token0.Add(state.protocolFee))
			if err != nil {
				return
			}
		}
	} else {
		err = uc.repo.SaveFeeGrowthGlobal1(pool.Id, state.feeGrowthGlobal)
		if err != nil {
			return
		}
		if state.protocolFee.GreaterThan(decimal.Zero) {
			var token1 decimal.Decimal
			_, token1, err = uc.repo.GetProtocolFee(pool.Id)
			if err != nil {
				return
			}
			err = uc.repo.SaveProtocolFeeToken1(pool.Id, token1.Add(state.protocolFee))
			if err != nil {
				return
			}
		}
	}

	if zeroForOne == exactInput {
		amount0, amount1 = amountSpecified.Sub(state.amountSpecifiedRemaining), state.amountCalculated
	} else {
		amount0, amount1 = state.amountCalculated, amountSpecified.Sub(state.amountSpecifiedRemaining)
	}

	// TODO: do the transfers and collect payment
	//if (zeroForOne) {
	//	if (amount1 < 0) TransferHelper.safeTransfer(token1, recipient, uint256(-amount1));
	//
	//	uint256 balance0Before = balance0();
	//	IUniswapV3SwapCallback(msg.sender).uniswapV3SwapCallback(amount0, amount1, data);
	//	require(balance0Before.add(uint256(amount0)) <= balance0(), 'IIA');
	//} else {
	//	if (amount0 < 0) TransferHelper.safeTransfer(token0, recipient, uint256(-amount0));
	//
	//	uint256 balance1Before = balance1();
	//	IUniswapV3SwapCallback(msg.sender).uniswapV3SwapCallback(amount0, amount1, data);
	//	require(balance1Before.add(uint256(amount1)) <= balance1(), 'IIA');
	//}
	return
}
