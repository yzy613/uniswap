// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PoolDao is the data access object for table pool.
type PoolDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns PoolColumns // columns contains all the column names of Table for convenient usage.
}

// PoolColumns defines and stores column names for table pool.
type PoolColumns struct {
	Id                  string //
	Token0Address       string //
	Token1Address       string //
	Fee                 string //
	TickSpacing         string //
	MaxLiquidityPerTick string //
}

// poolColumns holds the columns for table pool.
var poolColumns = PoolColumns{
	Id:                  "id",
	Token0Address:       "token0_address",
	Token1Address:       "token1_address",
	Fee:                 "fee",
	TickSpacing:         "tick_spacing",
	MaxLiquidityPerTick: "max_liquidity_per_tick",
}

// NewPoolDao creates and returns a new DAO object for table data access.
func NewPoolDao() *PoolDao {
	return &PoolDao{
		group:   "default",
		table:   "pool",
		columns: poolColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PoolDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PoolDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PoolDao) Columns() PoolColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PoolDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PoolDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PoolDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
