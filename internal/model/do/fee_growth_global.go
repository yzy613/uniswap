// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// FeeGrowthGlobal is the golang structure of table fee_growth_global for DAO operations like Where/Data.
type FeeGrowthGlobal struct {
	g.Meta               `orm:"table:fee_growth_global, do:true"`
	PoolId               interface{} //
	FeeGrowthGlobal0X128 interface{} //
	FeeGrowthGlobal1X128 interface{} //
}
