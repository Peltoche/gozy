package config

import "github.com/Peltoche/gozy/sdk/client"

type Service interface {
	SaveClient(client *client.Client) error
	LoadClient(name string) (*client.Client, error)
	ListClients() ([]client.Client, error)
	DeleteClient(name string) error
}
