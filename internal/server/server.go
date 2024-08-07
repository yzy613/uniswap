package server

import (
	"github.com/google/wire"
)

// SrvProviderSet is server providers.
var SrvProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
)
