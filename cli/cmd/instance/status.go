package instance

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func NewStatusCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "Give the instance healthcheck status",
		Args:  cobra.NoArgs,
		Use:   "status",
		Run: func(cmd *cobra.Command, _ []string) {
			inst := utils.GetInstance(cmd, tb)

			res, err := tb.Instance(inst).Status()
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			fmt.Printf("The status for %q is %q\n\n", inst.Name(), res.Status)

			t := tabby.New()
			t.AddHeader("Name", "Status", "Latency")
			t.AddLine("cache", res.Cache, res.Latency["cache"])
			t.AddLine("couchdb", res.Couchdb, res.Latency["couchdb"])
			t.AddLine("FS", res.FS, res.Latency["fs"])
			t.Print()

		},
	}

	return &cmd
}
