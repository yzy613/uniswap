// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/shopspring/decimal"
)

// Observation is the golang structure for table observation.
type Observation struct {
	PoolId                        int64           `json:"pool_id"                          orm:"pool_id"                          ` //
	ObservationIndex              int             `json:"observation_index"                orm:"observation_index"                ` //
	BlockTimestamp                int             `json:"block_timestamp"                  orm:"block_timestamp"                  ` //
	TickCumulative                int             `json:"tick_cumulative"                  orm:"tick_cumulative"                  ` //
	SecondsPerLiquidityCumulative decimal.Decimal `json:"seconds_per_liquidity_cumulative" orm:"seconds_per_liquidity_cumulative" ` //
	Initialized                   int             `json:"initialized"                      orm:"initialized"                      ` //
}
