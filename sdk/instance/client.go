package instance

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPClient struct {
	instance *Instance
}

func NewHTTPClientFromString(instanceURL string) (*HTTPClient, error) {
	instance, err := NewFromStr(instanceURL)
	if err != nil {
		return nil, err
	}

	return NewHTTPClient(instance), nil
}

func NewHTTPClient(instance *Instance) *HTTPClient {
	return &HTTPClient{instance: instance}
}

func (c *HTTPClient) Status() (*Status, error) {
	statusURL := c.instance.URL().JoinPath("/status").String()

	res, err := http.Get(statusURL)
	if err != nil {
		return nil, fmt.Errorf("failed to check the status: %w", err)
	}
	defer res.Body.Close()

	var status Status
	err = json.NewDecoder(res.Body).Decode(&status)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the response status: %w", err)
	}

	return &status, nil
}
