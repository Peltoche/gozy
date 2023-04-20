package sdk

import (
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
)

type SDK struct {
	domain string
	client client.Service
	config config.Service
}

// NewHTTPClient instantiate a new [HTTPClient] for the given domain.
//
// The domain is an base url to your instance. For example it would
// be "https://foobar.mycozy.cloud" for the cozy hosted client with
// the "foobar" account.
func NewSDK(appName, domain string) *SDK {
	return &SDK{
		domain: domain,
		client: client.NewHTTPClient(domain),
		config: config.NewXDG(appName),
	}
}

// Domain return the domain linked to this client.
func (s *SDK) Domain() string {
	return s.domain
}

// Client return the Client Service used to manipulate the client resource.
func (s *SDK) Client() client.Service {
	return s.client
}

func (s *SDK) Config() config.Service {
	return s.config
}
