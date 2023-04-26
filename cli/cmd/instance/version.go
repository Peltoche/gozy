package instance

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func NewVersionCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "Ask the version to the server",
		Args:  cobra.NoArgs,
		Use:   "version",
		Run: func(cmd *cobra.Command, _ []string) {
			inst := utils.GetInstance(cmd, tb)

			res, err := tb.Instance(inst).Version()
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			fmt.Printf("The version for %q is %q\n\n", inst.Name(), res.Version)

			t := tabby.New()
			t.AddLine("Mode", res.BuildMode)
			t.AddLine("Build Time", res.BuildTime)
			t.AddLine("Runtime", res.RuntimeVersion)
			t.AddLine("Version", res.Version)
			t.Print()

		},
	}

	return &cmd
}
