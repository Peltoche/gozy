package instance

import (
	"fmt"
	"os"
	"path"

	"github.com/adrg/xdg"
)

type Storage struct {
	dir string
}

func NewStorage(appName string) *Storage {
	return &Storage{dir: path.Join(xdg.ConfigHome, appName)}
}

func (s *Storage) Dir() string {
	return s.dir
}

func (s *Storage) Load(instanceName string) (*Instance, error) {
	inst, err := NewFromStr(instanceName)
	if err != nil {
		return nil, fmt.Errorf("invalide instance name: %w", err)
	}

	instances, err := s.List()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the existing instances: %w", err)
	}

	for _, i := range instances {
		if i.Name() == inst.Name() {
			return &i, nil
		}
	}

	return nil, nil
}

func (s *Storage) List() ([]Instance, error) {
	entries, err := os.ReadDir(s.dir)
	if err != nil {
		return nil, fmt.Errorf("failed to open the config dir: %w", err)
	}

	res := []Instance{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		inst, err := NewFromStr(entry.Name())
		if err != nil {
			return nil, fmt.Errorf("found an invalid instance name: %q: %w", entry.Name(), err)
		}

		res = append(res, *inst)
	}

	return res, nil
}

func (s *Storage) Save(inst *Instance) error {
	return os.MkdirAll(path.Join(s.dir, inst.Name()), 0o755)
}

func (s *Storage) Forget(inst *Instance) error {
	return os.RemoveAll(path.Join(s.dir, inst.Name()))
}
