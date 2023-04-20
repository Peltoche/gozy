package client

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

func NewListCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "List the clients locally saved",
		Args:  cobra.NoArgs,
		Use:   "list",
		Run: func(cmd *cobra.Command, _ []string) {
			res, err := tb.Config().ListClients()
			if err != nil {
				cmd.PrintErr(err)
				os.Exit(1)
			}

			for _, client := range res {
				fmt.Println(client.ClientName)
			}
		},
	}

	return &cmd
}
