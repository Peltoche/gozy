package client

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/cobra"
)

func NewClientCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "client <command>",
		Short: "Manage clients",
		Long:  "Work with Cozy clients.",
		Example: heredoc.Doc(`
  $ cozy client register
  `),
	}

	cmd.AddCommand(NewRegisterCmd())
	cmd.AddCommand(NewListCmd())

	return cmd
}
