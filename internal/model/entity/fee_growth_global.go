// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/shopspring/decimal"
)

// FeeGrowthGlobal is the golang structure for table fee_growth_global.
type FeeGrowthGlobal struct {
	PoolId               int64           `json:"pool_id"                   orm:"pool_id"                 ` //
	FeeGrowthGlobal0X128 decimal.Decimal `json:"fee_growth_global_0_x_128" orm:"fee_growth_global0_x128" ` //
	FeeGrowthGlobal1X128 decimal.Decimal `json:"fee_growth_global_1_x_128" orm:"fee_growth_global1_x128" ` //
}
