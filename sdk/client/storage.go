package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Peltoche/gozy/sdk/instance"
)

type Storage struct {
	dir string
}

func NewStorage(dir string) *Storage {
	return &Storage{dir}
}

// SaveClient in `$XDG_CONFIG_HOME/{appName}/{instance}/clients//{name}`.
//
// If $XDG_DATA_HOME doesnt exists, fallback to `$HOME/.local/share/`.
func (s *Storage) Save(inst *instance.Instance, client *Client) error {
	fileDir := path.Join(s.dir, inst.Name())

	err := os.MkdirAll(fileDir, 0o755)
	if err != nil {
		return err
	}

	raw, _ := json.MarshalIndent(client, "  ", "  ")

	filePath := path.Join(fileDir, client.ClientName+".json")

	err = os.WriteFile(filePath, raw, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write into %s: %w", filePath, err)
	}

	return nil
}

// ListClients found in `$XDG_CONFIG_HOME/{appName}/clients`.
//
// If $XDG_DATA_HOME doesnt exists, fallback to `$HOME/.local/share/`.
func (s *Storage) List(inst *instance.Instance) ([]Client, error) {
	clientsDir := path.Join(s.dir, inst.Name())

	entries, err := os.ReadDir(clientsDir)
	if errors.Is(err, os.ErrNotExist) {
		return []Client{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open the dir %q: %w", clientsDir, err)
	}

	res := make([]Client, len(entries))

	for i, entry := range entries {
		clientName := strings.TrimSuffix(entry.Name(), ".json")

		client, err := s.Load(inst, clientName)
		if err != nil {
			return nil, fmt.Errorf("failed to load %q: %w", entry.Name(), err)
		}

		res[i] = *client
	}

	return res, nil
}

func (s *Storage) Load(inst *instance.Instance, client string) (*Client, error) {
	clientPath := path.Join(s.dir, inst.Name(), client+".json")

	raw, err := os.ReadFile(clientPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open %q: %s", clientPath, err)
	}

	var res Client
	err = json.Unmarshal(raw, &res)
	if err != nil {
		return nil, fmt.Errorf("invalid content in %q: %w", clientPath, err)
	}

	return &res, nil
}

func (s *Storage) Delete(inst *instance.Instance, client string) error {
	err := os.Remove(path.Join(s.dir, inst.Name(), client+".json"))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return err
}
