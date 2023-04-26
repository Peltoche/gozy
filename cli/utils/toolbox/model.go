package toolbox

import (
	"context"

	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
	"github.com/Peltoche/gozy/sdk/instance"
)

const appName = "gozy"

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
	Save(client *client.Client) error
	List() ([]client.Client, error)
	Load(client string) (*client.Client, error)
	Delete(client string) error
}

type InstanceService interface {
	Status() (*instance.Status, error)
	Version() (*instance.Version, error)
}

type InstanceStorageService interface {
	List() ([]instance.Instance, error)
	Forget(inst *instance.Instance) error
	Save(inst *instance.Instance) error
	Load(instance string) (*instance.Instance, error)
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
	InstanceStorage() InstanceStorageService
	Client(instance *instance.Instance) ClientService
	ClientStorage(inst *instance.Instance) ClientStorageService
	Config() ConfigService
	AppName() string
}
