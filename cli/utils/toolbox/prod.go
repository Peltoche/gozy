package toolbox

import (
	"path"

	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
	"github.com/Peltoche/gozy/sdk/instance"
)

// Prod is a [Toolbox] implementation used for the prod.
//
// It contains real implementations and so have access
// real data. It should never be used in tests.
type Prod struct {
	cfgSvc        ConfigService
	clientStorage ClientStorageService
}

func NewProd() *Prod {
	cfgSvc := config.NewXDGConfig(appName)
	clientStorage := client.NewStorage(path.Join(cfgSvc.Dir(), clientDir))

	return &Prod{cfgSvc, clientStorage}
}

func (p *Prod) AppName() string {
	return appName
}

func (p *Prod) Config() ConfigService {
	return p.cfgSvc
}

func (p *Prod) Instance(i *instance.Instance) InstanceService {
	return instance.NewHTTPClient(i)
}

func (p *Prod) Client(instance *instance.Instance) ClientService {
	return client.NewHTTPClient(instance)
}

func (p *Prod) ClientStorage() ClientStorageService {
	return p.clientStorage
}
