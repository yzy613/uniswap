// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Pool is the golang structure of table pool for DAO operations like Where/Data.
type Pool struct {
	g.Meta              `orm:"table:pool, do:true"`
	Id                  interface{} //
	Token0Address       interface{} //
	Token1Address       interface{} //
	Fee                 interface{} //
	TickSpacing         interface{} //
	MaxLiquidityPerTick interface{} //
}
