// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeeGrowthGlobalDao is the data access object for table fee_growth_global.
type FeeGrowthGlobalDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns FeeGrowthGlobalColumns // columns contains all the column names of Table for convenient usage.
}

// FeeGrowthGlobalColumns defines and stores column names for table fee_growth_global.
type FeeGrowthGlobalColumns struct {
	PoolId               string //
	FeeGrowthGlobal0X128 string //
	FeeGrowthGlobal1X128 string //
}

// feeGrowthGlobalColumns holds the columns for table fee_growth_global.
var feeGrowthGlobalColumns = FeeGrowthGlobalColumns{
	PoolId:               "pool_id",
	FeeGrowthGlobal0X128: "fee_growth_global0_x128",
	FeeGrowthGlobal1X128: "fee_growth_global1_x128",
}

// NewFeeGrowthGlobalDao creates and returns a new DAO object for table data access.
func NewFeeGrowthGlobalDao() *FeeGrowthGlobalDao {
	return &FeeGrowthGlobalDao{
		group:   "default",
		table:   "fee_growth_global",
		columns: feeGrowthGlobalColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FeeGrowthGlobalDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FeeGrowthGlobalDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FeeGrowthGlobalDao) Columns() FeeGrowthGlobalColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FeeGrowthGlobalDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FeeGrowthGlobalDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FeeGrowthGlobalDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
