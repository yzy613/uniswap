// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiquidityDao is the data access object for table liquidity.
type LiquidityDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns LiquidityColumns // columns contains all the column names of Table for convenient usage.
}

// LiquidityColumns defines and stores column names for table liquidity.
type LiquidityColumns struct {
	PoolId    string //
	Liquidity string //
}

// liquidityColumns holds the columns for table liquidity.
var liquidityColumns = LiquidityColumns{
	PoolId:    "pool_id",
	Liquidity: "liquidity",
}

// NewLiquidityDao creates and returns a new DAO object for table data access.
func NewLiquidityDao() *LiquidityDao {
	return &LiquidityDao{
		group:   "default",
		table:   "liquidity",
		columns: liquidityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LiquidityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LiquidityDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LiquidityDao) Columns() LiquidityColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LiquidityDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LiquidityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LiquidityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
