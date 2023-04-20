package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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

func (s *HTTPClient) Register(ctx context.Context, cmd *RegisterCmd) (*Client, error) {
	rawBody, _ := json.Marshal(cmd)

	req, err := http.NewRequest(http.MethodPost, "https://jeanbon.mycozy.cloud/auth/register", bytes.NewReader(rawBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create the register request: %w", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := s.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to register the application: %w", err)
	}
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response body: %w", err)
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		return nil, fmt.Errorf("unexpected response %q status for register: %s", res.Status, string(raw))
	}

	// First unmarshal for the client
	var resBody Client
	err = json.Unmarshal(raw, &resBody)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the register response: %s", err)
	}

	return &resBody, nil
}
