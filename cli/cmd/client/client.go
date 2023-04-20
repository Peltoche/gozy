package client

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

func NewClientCmd(tb toolbox.Toolbox) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "client <command>",
		Short: "Manage clients",
		Long:  "Work with Cozy clients.",
		Example: heredoc.Doc(`
  $ cozy client register
  `),
	}

	cmd.AddCommand(NewRegisterCmd(tb))
	cmd.AddCommand(NewListCmd(tb))

	return cmd
}
