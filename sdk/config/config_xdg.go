package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/adrg/xdg"
)

const configFileName = "config.json"

// XDGConfig manage the configurations by saving them into the host
// file system following the XDGConfig specification.
type XDGConfig struct {
	homeDir string
}

// NewXDGConfig instantiate a new [XDGConfig].
func NewXDGConfig(appName string) *XDGConfig {
	return &XDGConfig{homeDir: path.Join(xdg.ConfigHome, appName)}
}

func (c *XDGConfig) Dir() string {
	return c.homeDir
}

func (c *XDGConfig) EnsureDirExists() error {
	return os.MkdirAll(c.homeDir, 0o755)
}

func (c *XDGConfig) Save(cfg *Config) error {
	cfgPath := path.Join(c.homeDir, configFileName)

	rawContent, _ := json.MarshalIndent(cfg, "  ", "  ")

	err := os.WriteFile(cfgPath, rawContent, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write into %q: %w", cfgPath, err)
	}

	return nil
}

func (c *XDGConfig) Read() (*Config, error) {
	cfgPath := path.Join(c.homeDir, configFileName)

	rawFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open the file %q: %w", cfgPath, err)
	}

	var cfg Config
	err = json.Unmarshal(rawFile, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode the file %q: %w", cfgPath, err)
	}

	return &cfg, nil
}
