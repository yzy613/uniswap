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
	Price                      decimal.Decimal `json:"price"                        orm:"price"                        ` //
	CurrentTick                int             `json:"current_tick"                 orm:"current_tick"                 ` //
	ObservationIndex           int             `json:"observation_index"            orm:"observation_index"            ` //
	ObservationCardinality     int             `json:"observation_cardinality"      orm:"observation_cardinality"      ` //
	ObservationCardinalityNext int             `json:"observation_cardinality_next" orm:"observation_cardinality_next" ` //
	FeeProtocol0               int             `json:"fee_protocol_0"               orm:"fee_protocol0"                ` //
	FeeProtocol1               int             `json:"fee_protocol_1"               orm:"fee_protocol1"                ` //
	Unlocked                   int             `json:"unlocked"                     orm:"unlocked"                     ` //
}
