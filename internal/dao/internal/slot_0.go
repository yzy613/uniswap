// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// Slot0Dao is the data access object for table slot0.
type Slot0Dao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns Slot0Columns // columns contains all the column names of Table for convenient usage.
}

// Slot0Columns defines and stores column names for table slot0.
type Slot0Columns struct {
	PoolId                     string //
	Price                      string //
	CurrentTick                string //
	ObservationIndex           string //
	ObservationCardinality     string //
	ObservationCardinalityNext string //
	FeeProtocol                string //
	Unlocked                   string //
}

// slot0Columns holds the columns for table slot0.
var slot0Columns = Slot0Columns{
	PoolId:                     "pool_id",
	Price:                      "price",
	CurrentTick:                "current_tick",
	ObservationIndex:           "observation_index",
	ObservationCardinality:     "observation_cardinality",
	ObservationCardinalityNext: "observation_cardinality_next",
	FeeProtocol:                "fee_protocol",
	Unlocked:                   "unlocked",
}

// NewSlot0Dao creates and returns a new DAO object for table data access.
func NewSlot0Dao() *Slot0Dao {
	return &Slot0Dao{
		group:   "default",
		table:   "slot0",
		columns: slot0Columns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *Slot0Dao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *Slot0Dao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *Slot0Dao) Columns() Slot0Columns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *Slot0Dao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *Slot0Dao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *Slot0Dao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
