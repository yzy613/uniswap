// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PositionDao is the data access object for table position.
type PositionDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns PositionColumns // columns contains all the column names of Table for convenient usage.
}

// PositionColumns defines and stores column names for table position.
type PositionColumns struct {
	PoolId               string //
	OwnerAddress         string //
	TickLower            string //
	TickUpper            string //
	Liquidity            string //
	FeeGrowthInside0Last string //
	FeeGrowthInside1Last string //
	TokensOwed0          string //
	TokensOwed1          string //
}

// positionColumns holds the columns for table position.
var positionColumns = PositionColumns{
	PoolId:               "pool_id",
	OwnerAddress:         "owner_address",
	TickLower:            "tick_lower",
	TickUpper:            "tick_upper",
	Liquidity:            "liquidity",
	FeeGrowthInside0Last: "fee_growth_inside0_last",
	FeeGrowthInside1Last: "fee_growth_inside1_last",
	TokensOwed0:          "tokens_owed0",
	TokensOwed1:          "tokens_owed1",
}

// NewPositionDao creates and returns a new DAO object for table data access.
func NewPositionDao() *PositionDao {
	return &PositionDao{
		group:   "default",
		table:   "position",
		columns: positionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PositionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PositionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PositionDao) Columns() PositionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PositionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PositionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PositionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
