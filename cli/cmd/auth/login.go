package auth

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/Peltoche/gozy/sdk/config"
	"github.com/Peltoche/gozy/sdk/instance"
	"github.com/spf13/cobra"
)

func NewLoginCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Args: cobra.ExactArgs(1),
		Use:  "login [<instance-url>]",
		Run: func(cmd *cobra.Command, args []string) {
			inst, err := instance.NewFromStr(args[0])
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			cfg, err := getConfigOrBootstrap(cmd.Context(), tb, inst)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			fmt.Printf("client: %+v\n", cfg)
		},
	}

	return &cmd
}

func getConfigOrBootstrap(ctx context.Context, tb toolbox.Toolbox, inst *instance.Instance) (*config.Config, error) {
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
