package client

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(tb toolbox.Toolbox) *cobra.Command {
	var domain string

	cmd := cobra.Command{
		Short: "Delete a client.",
		Args:  cobra.ExactArgs(1),
		Use:   "delete [<name>]",
		Run: func(cmd *cobra.Command, args []string) {
			if domain == "" {
				cmd.PrintErrln("--domain missing")
				os.Exit(1)
			}

			res, err := tb.Config().LoadClient(args[0])
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			err = tb.Client(domain).Delete(cmd.Context(), &client.DeleteCmd{
				ClientID:        res.ClientID,
				RegistrationCmd: res.RegistrationToken,
			})
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			tb.Config().DeleteClient(res.ClientName)

			fmt.Printf("The client %q have been deleted\n", res.ClientName)

		},
	}

	cmd.Flags().StringVar(&domain, "domain", "", "Domain to contact (example: \"foobar.mycozy.cloud\")")

	return &cmd
}
