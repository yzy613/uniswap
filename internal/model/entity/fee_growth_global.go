// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/shopspring/decimal"
)

// FeeGrowthGlobal is the golang structure for table fee_growth_global.
type FeeGrowthGlobal struct {
	PoolId           int64           `json:"pool_id"             orm:"pool_id"            ` //
	FeeGrowthGlobal0 decimal.Decimal `json:"fee_growth_global_0" orm:"fee_growth_global0" ` //
	FeeGrowthGlobal1 decimal.Decimal `json:"fee_growth_global_1" orm:"fee_growth_global1" ` //
}
