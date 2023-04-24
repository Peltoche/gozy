package toolbox

import (
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
)

const appName = "gozy"

// Prod is a [Toolbox] implementation used for the prod.
//
// It contains real implementations and so have access
// real data. It should never be used in tests.
type Prod struct{}

func NewProd() *Prod {
	return &Prod{}
}

func (p *Prod) Config() config.Service {
	return config.NewXDGConfig(appName)
}

func (p *Prod) Client(domain string) client.Service {
	return client.NewHTTPClient(domain)
}
