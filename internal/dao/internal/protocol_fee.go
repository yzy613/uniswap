// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProtocolFeeDao is the data access object for table protocol_fee.
type ProtocolFeeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ProtocolFeeColumns // columns contains all the column names of Table for convenient usage.
}

// ProtocolFeeColumns defines and stores column names for table protocol_fee.
type ProtocolFeeColumns struct {
	PoolId     string //
	Token0Fees string //
	Token1Fees string //
}

// protocolFeeColumns holds the columns for table protocol_fee.
var protocolFeeColumns = ProtocolFeeColumns{
	PoolId:     "pool_id",
	Token0Fees: "token0_fees",
	Token1Fees: "token1_fees",
}

// NewProtocolFeeDao creates and returns a new DAO object for table data access.
func NewProtocolFeeDao() *ProtocolFeeDao {
	return &ProtocolFeeDao{
		group:   "default",
		table:   "protocol_fee",
		columns: protocolFeeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProtocolFeeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProtocolFeeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProtocolFeeDao) Columns() ProtocolFeeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProtocolFeeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProtocolFeeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProtocolFeeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
