package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/google/wire"
	"uniswap/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPoolRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	toml := `
[database.default]
link = "%v"
`
	adapter, err := gcfg.NewAdapterContent(fmt.Sprintf(toml, c.Database.Driver+":"+c.Database.Source))
	if err != nil {
		return nil, nil, err
	}
	gcfg.Instance().SetAdapter(adapter)

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	return &Data{}, cleanup, nil
}
