package auth

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
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

			cfg, err := utils.GetConfigOrBootstrap(cmd.Context(), tb, inst)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			fmt.Printf("client: %+v\n", cfg)
		},
	}

	return &cmd
}
