// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Observation is the golang structure of table observation for DAO operations like Where/Data.
type Observation struct {
	g.Meta                            `orm:"table:observation, do:true"`
	PoolId                            interface{} //
	ObservationIndex                  interface{} //
	BlockTimestamp                    interface{} //
	TickCumulative                    interface{} //
	SecondsPerLiquidityCumulativeX128 interface{} //
	Initialized                       interface{} //
}
