// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Slot0 is the golang structure of table slot0 for DAO operations like Where/Data.
type Slot0 struct {
	g.Meta                     `orm:"table:slot0, do:true"`
	PoolId                     interface{} //
	Price                      interface{} //
	CurrentTick                interface{} //
	ObservationIndex           interface{} //
	ObservationCardinality     interface{} //
	ObservationCardinalityNext interface{} //
	FeeProtocol0               interface{} //
	FeeProtocol1               interface{} //
	Unlocked                   interface{} //
}
