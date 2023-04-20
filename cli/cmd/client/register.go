package client

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils"
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/spf13/cobra"
)

func NewRegisterCmd() *cobra.Command {
	var opt client.RegisterCmd

	cmd := cobra.Command{
		Short: "Register a new application client.",
		Long: `This route is used by OAuth2 clients to dynamically register them-selves.

See OAuth 2.0 Dynamic Client Registration Protocol for the details.`,
		Args: cobra.ExactArgs(1),
		Use:  "register [<name>]",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.PrintErrln("must provide <name> when not running interactively\n")
				cmd.Usage()
				return
			}

			opt.ClientName = args[0]

			if len(opt.RedirectURIs) == 0 || opt.SoftwareID == "" {
				cmd.PrintErrln("must provide --redirect-uris and --software-id when not running interactively\n")
				cmd.Usage()
				return
			}

			sdk := utils.NewClient(cmd)

			res, err := sdk.Client().Register(cmd.Context(), &opt)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			fmt.Printf("Client %q created in %s\n", res.ClientName, sdk.Domain())
			sdk.Config().SaveClient(res)
		},
	}

	// Required fields
	cmd.Flags().StringSliceVar(&opt.RedirectURIs, "redirect-uris", []string{}, "All the available redirect URIs for this client.")
	cmd.Flags().StringVar(&opt.SoftwareID, "software-id", "", "Identifier by the client.")

	// Optionals fields
	cmd.Flags().StringVar(&opt.ClientKind, "client-kind", "", "Possible values: web, desktop, mobile, browser, etc.")
	cmd.Flags().StringVar(&opt.ClientURI, "client-uri", "", "A web page URL providing information about the client.")
	cmd.Flags().StringVar(&opt.LogoURI, "logo-uri", "", "An icon URL displayed during the authorization flow.")
	cmd.Flags().StringVar(&opt.NotificationPlatform, "notification-platform", "", "To activate notifications on the associated device.")
	cmd.Flags().StringVar(&opt.PolicyURI, "policy-uri", "", "URL string pointing to a human-readable policy.")
	cmd.Flags().StringVar(&opt.SoftwareVersion, "software-version", "", "A version identifier string for the client software")

	return &cmd
}
