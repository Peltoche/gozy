package utils

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
	"github.com/Peltoche/gozy/sdk/instance"
)

func GetConfigOrBootstrap(ctx context.Context, tb toolbox.Toolbox, inst *instance.Instance) (*config.Config, error) {
	cfgSvc := tb.Config()

	cfg, err := cfgSvc.Read()
	if errors.Is(err, os.ErrNotExist) {
		err = registerClient(ctx, tb, inst)
		if err != nil {
			return nil, err
		}

		cfg, err = cfgSvc.Read()
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func registerClient(ctx context.Context, tb toolbox.Toolbox, inst *instance.Instance) error {
	cfgSvc := tb.Config()

	err := cfgSvc.EnsureDirExists()
	if err != nil {
		return err
	}

	clientName := tb.AppName()

	hostname, _ := os.Hostname()
	if hostname != "" {
		clientName += "-" + hostname
	}

	res, err := tb.Client(inst).Register(ctx, &client.RegisterCmd{
		ClientName:   clientName,
		SoftwareID:   "github.com/Peltoche/gozy",
		ClientKind:   "CLI",
		ClientURI:    "github.com/Peltoche/gozy",
		RedirectURIs: []string{"http://localhost"},
	})
	if err != nil {
		return fmt.Errorf("failed to register the gozy client: %w", err)
	}

	err = tb.ClientStorage(inst).Save(res)
	if err != nil {
		return fmt.Errorf("failed to save the client config: %w", err)
	}

	err = cfgSvc.Save(&config.Config{
		Instance: inst.Name(),
		Client:   res.ClientName,
	})
	if err != nil {
		return fmt.Errorf("failed to save the config: %w", err)
	}

	return nil
}
