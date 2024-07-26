// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/shopspring/decimal"
)

// Liquidity is the golang structure for table liquidity.
type Liquidity struct {
	PoolId    int64           `json:"pool_id"   orm:"pool_id"   ` //
	Liquidity decimal.Decimal `json:"liquidity" orm:"liquidity" ` //
}
