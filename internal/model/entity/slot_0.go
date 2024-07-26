// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/shopspring/decimal"
)

// Slot0 is the golang structure for table slot0.
type Slot0 struct {
	PoolId                     int64           `json:"pool_id"                      orm:"pool_id"                      ` //
	SqrtPriceX96               decimal.Decimal `json:"sqrt_price_x_96"              orm:"sqrt_price_x96"               ` //
	CurrentTick                int             `json:"current_tick"                 orm:"current_tick"                 ` //
	ObservationIndex           int             `json:"observation_index"            orm:"observation_index"            ` //
	ObservationCardinality     int             `json:"observation_cardinality"      orm:"observation_cardinality"      ` //
	ObservationCardinalityNext int             `json:"observation_cardinality_next" orm:"observation_cardinality_next" ` //
	FeeProtocol                int             `json:"fee_protocol"                 orm:"fee_protocol"                 ` //
	Unlocked                   int             `json:"unlocked"                     orm:"unlocked"                     ` //
}
