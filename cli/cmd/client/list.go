package client

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/sdk/config"
	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	cmd := cobra.Command{
		Short: "List the clients locally saved",
		Args:  cobra.NoArgs,
		Use:   "list",
		Run: func(cmd *cobra.Command, _ []string) {
			res, err := config.NewXDG("gozy").ListClients()
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
