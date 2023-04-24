package sdk

import (
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
)

type SDK struct {
	domain       string
	unauthClient client.UnauthenticatedService
	config       config.Service
}

// NewHTTPClient instantiate a new [HTTPClient] for the given domain.
//
// The domain is an base url to your instance. For example it would
// be "https://foobar.mycozy.cloud" for the cozy hosted client with
// the "foobar" account.
func NewSDK(appName, domain string) *SDK {
	return &SDK{
		domain:       domain,
		unauthClient: client.NewUnauthenticatedHTTPClient(domain),
		config:       config.NewXDGConfig(appName),
	}
}

// Domain return the domain linked to this client.
func (s *SDK) Domain() string {
	return s.domain
}

// UnauthClient return the Client Service used to manipulate the client resource.
func (s *SDK) UnauthClient() client.UnauthenticatedService {
	return s.unauthClient
}

func (s *SDK) Config() config.Service {
	return s.config
}
