// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/shopspring/decimal"
)

// Position is the golang structure for table position.
type Position struct {
	PoolId               int64           `json:"pool_id"                  orm:"pool_id"                 ` //
	OwnerAddress         string          `json:"owner_address"            orm:"owner_address"           ` //
	TickLower            int             `json:"tick_lower"               orm:"tick_lower"              ` //
	TickUpper            int             `json:"tick_upper"               orm:"tick_upper"              ` //
	Liquidity            decimal.Decimal `json:"liquidity"                orm:"liquidity"               ` //
	FeeGrowthInside0Last decimal.Decimal `json:"fee_growth_inside_0_last" orm:"fee_growth_inside0_last" ` //
	FeeGrowthInside1Last decimal.Decimal `json:"fee_growth_inside_1_last" orm:"fee_growth_inside1_last" ` //
	TokensOwed0          decimal.Decimal `json:"tokens_owed_0"            orm:"tokens_owed0"            ` //
	TokensOwed1          decimal.Decimal `json:"tokens_owed_1"            orm:"tokens_owed1"            ` //
}
