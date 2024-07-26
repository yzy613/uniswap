// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Position is the golang structure of table position for DAO operations like Where/Data.
type Position struct {
	g.Meta                   `orm:"table:position, do:true"`
	PoolId                   interface{} //
	OwnerAddress             interface{} //
	TickLower                interface{} //
	TickUpper                interface{} //
	Liquidity                interface{} //
	FeeGrowthInside0LastX128 interface{} //
	FeeGrowthInside1LastX128 interface{} //
	TokensOwed0              interface{} //
	TokensOwed1              interface{} //
}
