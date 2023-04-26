package instance

import (
	"fmt"
	"net/url"
)

type Version struct {
	BuildMode      string `json:"build_mode"`
	BuildTime      string `json:"build_time"`
	RuntimeVersion string `json:"runtime_version"`
	Version        string `json:"version"`
}

type Status struct {
	Cache   string            `json:"cache"`
	Couchdb string            `json:"couchdb"`
	FS      string            `json:"fs"`
	Latency map[string]string `json:"latency"`
	Message string            `json:"message"`
	Status  string            `json:"status"`
}

type Instance struct {
	u *url.URL
}

func NewFromStr(in string) (*Instance, error) {
	u, err := url.Parse(in)
	if err != nil {
		return nil, fmt.Errorf("invalid instance url: %w", err)
	}

	if u.Host == "" {
		return NewFromStr("https://" + in)
	}

	// Keep only the host and the port and for the use of https
	instance := url.URL{
		Scheme: "https",
		Host:   u.Host,
	}

	return &Instance{u: &instance}, nil
}

func (i *Instance) Name() string {
	return i.u.Hostname()
}

func (i *Instance) URL() *url.URL {
	return &url.URL{
		Scheme: i.u.Scheme,
		Host:   i.u.Host,
	}
}
