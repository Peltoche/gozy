package client

import (
	"fmt"

	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

func NewForgetCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "Forget a client on this machine.",
		Args:  cobra.ExactArgs(1),
		Use:   "forget [<name>]",
		Run: func(_ *cobra.Command, args []string) {
			tb.Config().DeleteClient(args[0])

			fmt.Printf("The client %q have been forgotten\n", args[0])

		},
	}

	return &cmd
}
