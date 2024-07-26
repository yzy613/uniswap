// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeeAmountDao is the data access object for table fee_amount.
type FeeAmountDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns FeeAmountColumns // columns contains all the column names of Table for convenient usage.
}

// FeeAmountColumns defines and stores column names for table fee_amount.
type FeeAmountColumns struct {
	Fee         string //
	TickSpacing string //
}

// feeAmountColumns holds the columns for table fee_amount.
var feeAmountColumns = FeeAmountColumns{
	Fee:         "fee",
	TickSpacing: "tick_spacing",
}

// NewFeeAmountDao creates and returns a new DAO object for table data access.
func NewFeeAmountDao() *FeeAmountDao {
	return &FeeAmountDao{
		group:   "default",
		table:   "fee_amount",
		columns: feeAmountColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeeAmountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeeAmountDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeeAmountDao) Columns() FeeAmountColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeeAmountDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeeAmountDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeeAmountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
