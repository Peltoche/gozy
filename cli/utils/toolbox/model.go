package toolbox

import (
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
)

// Toolbox give access to all the required tools.
//
// It allow to easily inject all the required tool and be
// able to replace them for testing purprose.
//
// Several implementations exists:
// - [Prod] With the real tools. Use for prod.
// - [Mock] With only mocked tools. Used for testing.
type Toolbox interface {
	Client(domain string) client.Service
	Config() config.Service
}
