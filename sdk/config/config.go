package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Peltoche/gozy/sdk/client"
	"github.com/adrg/xdg"
)

const clientDir = "clients"

// XDG manage the configurations by saving them into the host
// file system following the XDG specification.
type XDG struct {
	appName string
}

// NewXDG instantiate a new [XDGConfig].
//
// appName describe the application name. It will be used
// for the subfolder naming. If appName is empty the function
// will panic.
func NewXDG(appName string) *XDG {
	if appName == "" {
		panic("appName must not be empty")
	}

	return &XDG{appName}
}

// SaveClient in `$XDG_DATA_HOME/{appName}/clients/{appName}/clients/{name}`.
//
// If $XDG_DATA_HOME doesnt exists, fallback to `$HOME/.local/share/`.
func (c *XDG) SaveClient(client *client.Client) error {
	dataFile, err := xdg.DataFile(path.Join(c.appName, clientDir, client.ClientName+".json"))
	if err != nil {
		return err
	}

	raw, _ := json.Marshal(client)

	err = os.WriteFile(dataFile, raw, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write into %s: %w", dataFile, err)
	}

	return nil
}

// ListClients found in `$XDG_DATA_HOME/{appName}/clients/{appName}/clients/{name}`.
//
// If $XDG_DATA_HOME doesnt exists, fallback to `$HOME/.local/share/`.
func (c *XDG) LoadClient(name string) (*client.Client, error) {
	clientFile, err := xdg.SearchDataFile(path.Join(c.appName, clientDir, name+".json"))
	if err != nil {
		return nil, err
	}

	raw, err := os.ReadFile(clientFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open %q: %s", clientFile, err)
	}

	var client client.Client
	err = json.Unmarshal(raw, &client)
	if err != nil {
		return nil, fmt.Errorf("invalid content in %q: %w", clientFile, err)
	}

	return &client, nil
}

// ListClients found in `$XDG_DATA_HOME/{appName}/clients`.
//
// If $XDG_DATA_HOME doesnt exists, fallback to `$HOME/.local/share/`.
func (c *XDG) ListClients() ([]client.Client, error) {
	clientsDir := path.Join(xdg.DataHome, c.appName, clientDir)

	entries, err := os.ReadDir(clientsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to open the dir %q: %w", clientsDir, err)
	}

	res := make([]client.Client, len(entries))

	for i, entry := range entries {
		clientName := strings.TrimSuffix(entry.Name(), ".json")

		client, err := c.LoadClient(clientName)
		if err != nil {
			return nil, fmt.Errorf("failed to load %q: %w", entry.Name(), err)
		}

		res[i] = *client
	}

	return res, nil
}
