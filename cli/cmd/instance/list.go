package instance

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

func NewListCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "List the instances locally saved",
		Args:  cobra.NoArgs,
		Use:   "list",
		Run: func(cmd *cobra.Command, _ []string) {
			instSvc := tb.InstanceStorage()

			res, err := instSvc.List()
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			for _, inst := range res {
				fmt.Println(inst.Name())
			}
		},
	}

	return &cmd
}
