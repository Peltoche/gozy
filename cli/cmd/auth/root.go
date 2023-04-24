package auth

import (
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

func NewAuthCmd(tb toolbox.Toolbox) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "auth <command>",
		Short: "Authenticate to cozy",
	}

	cmd.AddCommand(NewLoginCmd(tb))

	return cmd
}
