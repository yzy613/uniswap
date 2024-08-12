// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/shopspring/decimal"
)

// Pool is the golang structure for table pool.
type Pool struct {
	Id                  int64           `json:"id"                     orm:"id"                     ` //
	Token0Address       string          `json:"token_0_address"        orm:"token0_address"         ` //
	Token1Address       string          `json:"token_1_address"        orm:"token1_address"         ` //
	Fee                 int             `json:"fee"                    orm:"fee"                    ` //
	TickSpacing         int             `json:"tick_spacing"           orm:"tick_spacing"           ` //
	MaxLiquidityPerTick decimal.Decimal `json:"max_liquidity_per_tick" orm:"max_liquidity_per_tick" ` //
	Balance0            decimal.Decimal `json:"balance_0"              orm:"balance0"               ` //
	Balance1            decimal.Decimal `json:"balance_1"              orm:"balance1"               ` //
}
