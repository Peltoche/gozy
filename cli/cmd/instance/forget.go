package instance

import (
	"fmt"

	"github.com/Peltoche/gozy/cli/utils"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

func NewForgetCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "Forget a client on this machine.",
		Args:  cobra.NoArgs,
		Use:   "forget",
		Run: func(cmd *cobra.Command, args []string) {
			inst := utils.GetInstance(cmd, tb)

			tb.InstanceStorage().Forget(inst)

			fmt.Printf("The %q have been forgotten\n", args[0])
		},
	}

	return &cmd
}
