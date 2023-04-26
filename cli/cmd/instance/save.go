package instance

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/Peltoche/gozy/sdk/instance"
	"github.com/spf13/cobra"
)

func NewSaveCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "List the clients locally saved",
		Args:  cobra.ExactArgs(1),
		Use:   "save [<instance-url>]",
		Run: func(cmd *cobra.Command, args []string) {
			instStorageSvc := tb.InstanceStorage()

			inst, err := instance.NewFromStr(args[0])
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			err = instStorageSvc.Save(inst)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			fmt.Printf("Instance %q saved locally\n", inst.Name())

		},
	}

	return &cmd
}
