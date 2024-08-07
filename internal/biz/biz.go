package biz

import "github.com/google/wire"

// BizProviderSet is biz providers.
var BizProviderSet = wire.NewSet(
	NewLiquidityUsecase,
	NewObservationUsecase,
	NewPoolUsecase,
	NewTickUsecase,
	NewTickBitMapUsecase,
)
