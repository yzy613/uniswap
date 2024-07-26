// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Tick is the golang structure of table tick for DAO operations like Where/Data.
type Tick struct {
	g.Meta                         `orm:"table:tick, do:true"`
	PoolId                         interface{} //
	TickIndex                      interface{} //
	LiquidityGross                 interface{} //
	LiquidityNet                   interface{} //
	FeeGrowthOutside0X128          interface{} //
	FeeGrowthOutside1X128          interface{} //
	SecondsPerLiquidityOutsideX128 interface{} //
	TickCumulativeOutside          interface{} //
	SecondsOutside                 interface{} //
	Initialized                    interface{} //
}
