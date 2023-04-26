package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Peltoche/gozy/sdk/instance"
	"github.com/adrg/xdg"
)

const clientDir = "clients"

type Storage struct {
	dir string
}

func NewStorage(appName string, inst *instance.Instance) *Storage {
	return &Storage{dir: path.Join(xdg.ConfigHome, appName, inst.Name(), clientDir)}
}

func (s *Storage) Save(client *Client) error {
	err := os.MkdirAll(s.dir, 0o755)
	if err != nil {
		return err
	}

	raw, _ := json.MarshalIndent(client, "  ", "  ")

	filePath := path.Join(s.dir, client.ClientName+".json")

	err = os.WriteFile(filePath, raw, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write into %s: %w", filePath, err)
	}

	return nil
}

func (s *Storage) List() ([]Client, error) {
	entries, err := os.ReadDir(s.dir)
	if errors.Is(err, os.ErrNotExist) {
		return []Client{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open the dir %q: %w", s.dir, err)
	}

	res := make([]Client, len(entries))

	for i, entry := range entries {
		clientName := strings.TrimSuffix(entry.Name(), ".json")

		client, err := s.Load(clientName)
		if err != nil {
			return nil, fmt.Errorf("failed to load %q: %w", entry.Name(), err)
		}

		res[i] = *client
	}

	return res, nil
}

func (s *Storage) Load(client string) (*Client, error) {
	clientPath := path.Join(s.dir, client+".json")

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

func (s *Storage) Delete(client string) error {
	err := os.Remove(path.Join(s.dir, client+".json"))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return err
}
