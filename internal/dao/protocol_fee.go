// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"uniswap/internal/dao/internal"
)

// internalProtocolFeeDao is internal type for wrapping internal DAO implements.
type internalProtocolFeeDao = *internal.ProtocolFeeDao

// protocolFeeDao is the data access object for table protocol_fee.
// You can define custom methods on it to extend its functionality as you wish.
type protocolFeeDao struct {
	internalProtocolFeeDao
}

var (
	// ProtocolFee is globally public accessible object for table protocol_fee operations.
	ProtocolFee = protocolFeeDao{
		internal.NewProtocolFeeDao(),
	}
)

// Fill with you ideas below.