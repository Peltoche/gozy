package instance

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

func NewInstanceCmd(tb toolbox.Toolbox) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "instance <command>",
		Short: "Manage instances",
		Long:  "Work the the connected instances",
		Example: heredoc.Doc(`
  $ cozy instance list
  `),
	}

	cmd.AddCommand(NewListCmd(tb))
	cmd.AddCommand(NewForgetCmd(tb))
	cmd.AddCommand(NewSaveCmd(tb))
	cmd.AddCommand(NewSetDefaultCmd(tb))

	return cmd
}
