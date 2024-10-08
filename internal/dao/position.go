// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"uniswap/internal/dao/internal"
)

// internalPositionDao is internal type for wrapping internal DAO implements.
type internalPositionDao = *internal.PositionDao

// positionDao is the data access object for table position.
// You can define custom methods on it to extend its functionality as you wish.
type positionDao struct {
	internalPositionDao
}

var (
	// Position is globally public accessible object for table position operations.
	Position = positionDao{
		internal.NewPositionDao(),
	}
)

// Fill with you ideas below.
