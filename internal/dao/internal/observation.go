// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ObservationDao is the data access object for table observation.
type ObservationDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ObservationColumns // columns contains all the column names of Table for convenient usage.
}

// ObservationColumns defines and stores column names for table observation.
type ObservationColumns struct {
	PoolId                            string //
	ObservationIndex                  string //
	BlockTimestamp                    string //
	TickCumulative                    string //
	SecondsPerLiquidityCumulativeX128 string //
	Initialized                       string //
}

// observationColumns holds the columns for table observation.
var observationColumns = ObservationColumns{
	PoolId:                            "pool_id",
	ObservationIndex:                  "observation_index",
	BlockTimestamp:                    "block_timestamp",
	TickCumulative:                    "tick_cumulative",
	SecondsPerLiquidityCumulativeX128: "seconds_per_liquidity_cumulative_x128",
	Initialized:                       "initialized",
}

// NewObservationDao creates and returns a new DAO object for table data access.
func NewObservationDao() *ObservationDao {
	return &ObservationDao{
		group:   "default",
		table:   "observation",
		columns: observationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ObservationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ObservationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ObservationDao) Columns() ObservationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ObservationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ObservationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ObservationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
