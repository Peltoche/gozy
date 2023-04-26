package instance

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

func NewSetDefaultCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "Set the instance as the default one",
		Args:  cobra.ExactArgs(1),
		Use:   "set-default [instance-name]",
		Run: func(cmd *cobra.Command, args []string) {
			inst, err := tb.InstanceStorage().Load(args[0])
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			if inst == nil {
				cmd.PrintErrf("Instance %q not found, run \"gozy instance list\" to see the available instances\n", args[0])
				os.Exit(1)
			}

			cfg, err := utils.GetConfigOrBootstrap(cmd.Context(), tb, inst)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			cfg.Instance = inst.Name()

			err = tb.Config().Save(cfg)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			fmt.Printf("Instance %q set as default\n", inst.Name())
		},
	}

	return &cmd
}
