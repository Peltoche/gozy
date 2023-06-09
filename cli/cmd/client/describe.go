package client

import (
	"os"

	"github.com/Peltoche/gozy/cli/utils"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
)

func NewDescribeCmd(tb toolbox.Toolbox) *cobra.Command {
	cmd := cobra.Command{
		Short: "Describe a client content.",
		Args:  cobra.ExactArgs(1),
		Use:   "describe [<name>]",
		Run: func(cmd *cobra.Command, args []string) {
			inst := utils.GetInstance(cmd, tb)

			res, err := tb.ClientStorage(inst).Load(args[0])
			if err != nil {
				cmd.Printf("client %q not found: run \"%s client list\" to have the available clients\n", args[0], tb.AppName())
				os.Exit(1)
			}

			t := tabby.New()
			t.AddLine("client ID", res.ClientID)
			t.AddLine("client secret", res.ClientSecret)
			t.AddLine("name", res.ClientName)
			t.AddLine("kind", res.ClientKind)
			t.AddLine("uri", res.ClientURI)
			t.AddLine("logo URI", res.LogoURI)
			t.AddLine("policy URI", res.PolicyURI)
			t.AddLine("redirect URIs", res.RedirectURIs)
			t.AddLine("registration token", res.RegistrationToken)
			t.AddLine("secret expires at", res.SecretExpiresAt)
			t.AddLine("software ID", res.SoftwareID)
			t.AddLine("software version", res.SoftwareVersion)
			t.Print()
		},
	}

	return &cmd
}
