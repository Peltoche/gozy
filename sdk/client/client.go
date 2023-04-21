package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrRequestFormat    = errors.New("failed to generate the request")
	ErrRequestSend      = errors.New("request failed")
	ErrReadBody         = errors.New("failed to read the response body")
	ErrUnexpectedStatus = errors.New("unexpected response status")
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
		return nil, fmt.Errorf("%w: %w", ErrRequestFormat, err)
	}
	defer res.Body.Close()

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrReadBody, err)
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		return nil, fmt.Errorf("%w: %q : %s", ErrUnexpectedStatus, res.Status, string(raw))
	}

	// First unmarshal for the client
	var resBody Client
	err = json.Unmarshal(raw, &resBody)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}

	return &resBody, nil
}

func (s *HTTPClient) Delete(ctx context.Context, cmd *DeleteCmd) error {
	req, err := http.NewRequest(http.MethodDelete, "https://jeanbon.mycozy.cloud/auth/register/"+cmd.ClientID, nil)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRequestFormat, err)
	}

	res, err := s.client.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("%w: %w", ErrRequestSend, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		raw, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrReadBody, err)
		}

		return fmt.Errorf("%w: %q : %s", ErrUnexpectedStatus, res.Status, string(raw))
	}

	return nil
}
