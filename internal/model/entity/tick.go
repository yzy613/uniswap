// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/shopspring/decimal"
)

// Tick is the golang structure for table tick.
type Tick struct {
	PoolId                         int64           `json:"pool_id"                             orm:"pool_id"                            ` //
	TickIndex                      int             `json:"tick_index"                          orm:"tick_index"                         ` //
	LiquidityGross                 decimal.Decimal `json:"liquidity_gross"                     orm:"liquidity_gross"                    ` //
	LiquidityNet                   decimal.Decimal `json:"liquidity_net"                       orm:"liquidity_net"                      ` //
	FeeGrowthOutside0X128          decimal.Decimal `json:"fee_growth_outside_0_x_128"          orm:"fee_growth_outside0_x128"           ` //
	FeeGrowthOutside1X128          decimal.Decimal `json:"fee_growth_outside_1_x_128"          orm:"fee_growth_outside1_x128"           ` //
	SecondsPerLiquidityOutsideX128 decimal.Decimal `json:"seconds_per_liquidity_outside_x_128" orm:"seconds_per_liquidity_outside_x128" ` //
	TickCumulativeOutside          decimal.Decimal `json:"tick_cumulative_outside"             orm:"tick_cumulative_outside"            ` //
	SecondsOutside                 decimal.Decimal `json:"seconds_outside"                     orm:"seconds_outside"                    ` //
	Initialized                    int             `json:"initialized"                         orm:"initialized"                        ` //
}
