package client

import "net/http"

// HTTPClient is the main interface used to interact with an the [Client] resource.
type HTTPClient struct {
	domain string
	client *http.Client
}

// NewHTTPClient instantiate a new [Service] for the given domain.
//
// The domain is an base url to your instance. For example it would
// be "https://foobar.mycozy.cloud" for the cozy hosted client with
// the "foobar" account.
func NewHTTPClient(domain string) *HTTPClient {
	return &HTTPClient{
		domain: domain,
		client: http.DefaultClient,
	}
}
