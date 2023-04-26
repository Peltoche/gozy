package toolbox

import (
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
	"github.com/Peltoche/gozy/sdk/instance"
)

// Prod is a [Toolbox] implementation used for the prod.
//
// It contains real implementations and so have access
// real data. It should never be used in tests.
type Prod struct {
	cfgSvc          ConfigService
	instanceStorage InstanceStorageService
}

func NewProd() *Prod {
	cfgSvc := config.NewXDGConfig(appName)
	instanceStorage := instance.NewStorage(appName)

	return &Prod{cfgSvc, instanceStorage}
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

func (p *Prod) InstanceStorage() InstanceStorageService {
	return p.instanceStorage
}

func (p *Prod) Client(inst *instance.Instance) ClientService {
	return client.NewHTTPClient(inst)
}

func (p *Prod) ClientStorage(inst *instance.Instance) ClientStorageService {
	return client.NewStorage(appName, inst)
}
