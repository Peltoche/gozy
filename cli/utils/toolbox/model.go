package toolbox

import (
	"context"

	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
	"github.com/Peltoche/gozy/sdk/instance"
)

const appName = "gozy"
const clientDir = "clients"

type ClientService interface {
	Register(ctx context.Context, cmd *client.RegisterCmd) (*client.Client, error)
	Delete(ctx context.Context, cmd *client.DeleteCmd) error
}

type ConfigService interface {
	EnsureDirExists() error
	Save(cfg *config.Config) error
	Read() (*config.Config, error)
}

type ClientStorageService interface {
	Save(inst *instance.Instance, client *client.Client) error
	List(inst *instance.Instance) ([]client.Client, error)
	Load(inst *instance.Instance, client string) (*client.Client, error)
	Delete(inst *instance.Instance, client string) error
}

type InstanceService interface {
}

// Toolbox give access to all the required tools.
//
// It allow to easily inject all the required tool and be
// able to replace them for testing purprose.
//
// Several implementations exists:
// - [Prod] With the real tools. Use for prod.
// - [Mock] With only mocked tools. Used for testing.
type Toolbox interface {
	Instance(instance *instance.Instance) InstanceService
	Client(instance *instance.Instance) ClientService
	ClientStorage() ClientStorageService
	Config() ConfigService
	AppName() string
}
