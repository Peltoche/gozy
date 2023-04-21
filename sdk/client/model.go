package client

import "context"

type Service interface {
	Register(ctx context.Context, cmd *RegisterCmd) (*Client, error)
	Delete(ctx context.Context, cmd *DeleteCmd) error
}

type Client struct {
	ClientID          string   `json:"client_id"`
	ClientSecret      string   `json:"client_secret"`
	ClientName        string   `json:"client_name"`
	ClientKind        string   `json:"client_kind,omitempty"`
	ClientURI         string   `json:"client_uri,omitempty"`
	LogoURI           string   `json:"logo_uri,omitempty"`
	PolicyURI         string   `json:"policy_uri,omitempty"`
	RedirectURIs      []string `json:"redirect_uris"`
	RegistrationToken string   `json:"registration_access_token"`
	SecretExpiresAt   int      `json:"client_secret_expires_at"`
	SoftwareID        string   `json:"software_id"`
	SoftwareVersion   string   `json:"software_version,omitempty"`
}

type RegisterCmd struct {
	ClientName           string   `json:"client_name"`
	SoftwareID           string   `json:"software_id"`
	ClientKind           string   `json:"client_kind,omitempty"`
	ClientURI            string   `json:"client_uri,omitempty"`
	LogoURI              string   `json:"logo_uri,omitempty"`
	PolicyURI            string   `json:"policy_uri,omitempty"`
	RedirectURIs         []string `json:"redirect_uris"`
	SecretExpiresAt      int      `json:"client_secret_expires_at"`
	SoftwareVersion      string   `json:"software_version,omitempty"`
	NotificationPlatform string   `json:"notification_platform"`
}

type DeleteCmd struct {
	ClientName      string
	RegistrationCmd string
}
