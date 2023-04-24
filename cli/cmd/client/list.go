package client

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

func NewListCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "List the clients locally saved",
		Args:  cobra.NoArgs,
		Use:   "list",
		Run: func(cmd *cobra.Command, _ []string) {
			inst := utils.GetInstance(cmd, tb)

			res, err := tb.ClientStorage().List(inst)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			for _, client := range res {
				fmt.Println(client.ClientName)
			}
		},
	}

	return &cmd
}
