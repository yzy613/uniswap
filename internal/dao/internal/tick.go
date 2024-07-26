// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TickDao is the data access object for table tick.
type TickDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TickColumns // columns contains all the column names of Table for convenient usage.
}

// TickColumns defines and stores column names for table tick.
type TickColumns struct {
	PoolId                         string //
	TickIndex                      string //
	LiquidityGross                 string //
	LiquidityNet                   string //
	FeeGrowthOutside0X128          string //
	FeeGrowthOutside1X128          string //
	SecondsPerLiquidityOutsideX128 string //
	TickCumulativeOutside          string //
	SecondsOutside                 string //
	Initialized                    string //
}

// tickColumns holds the columns for table tick.
var tickColumns = TickColumns{
	PoolId:                         "pool_id",
	TickIndex:                      "tick_index",
	LiquidityGross:                 "liquidity_gross",
	LiquidityNet:                   "liquidity_net",
	FeeGrowthOutside0X128:          "fee_growth_outside0_x128",
	FeeGrowthOutside1X128:          "fee_growth_outside1_x128",
	SecondsPerLiquidityOutsideX128: "seconds_per_liquidity_outside_x128",
	TickCumulativeOutside:          "tick_cumulative_outside",
	SecondsOutside:                 "seconds_outside",
	Initialized:                    "initialized",
}

// NewTickDao creates and returns a new DAO object for table data access.
func NewTickDao() *TickDao {
	return &TickDao{
		group:   "default",
		table:   "tick",
		columns: tickColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TickDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TickDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TickDao) Columns() TickColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TickDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TickDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TickDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
