// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ProtocolFee is the golang structure of table protocol_fee for DAO operations like Where/Data.
type ProtocolFee struct {
	g.Meta     `orm:"table:protocol_fee, do:true"`
	PoolId     interface{} //
	Token0Fees interface{} //
	Token1Fees interface{} //
}
