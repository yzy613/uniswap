// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/shopspring/decimal"
)

// ProtocolFee is the golang structure for table protocol_fee.
type ProtocolFee struct {
	PoolId     int64           `json:"pool_id"      orm:"pool_id"     ` //
	Token0Fees decimal.Decimal `json:"token_0_fees" orm:"token0_fees" ` //
	Token1Fees decimal.Decimal `json:"token_1_fees" orm:"token1_fees" ` //
}
