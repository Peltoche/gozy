package client

import (
	"fmt"
	"os"

	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/Peltoche/gozy/sdk/client"
	"github.com/spf13/cobra"
)

func NewRegisterCmd(tb toolbox.Toolbox) *cobra.Command {
	var opt client.RegisterCmd
	var domain string

	cmd := cobra.Command{
		Short: "Register a new application client.",
		Args:  cobra.ExactArgs(1),
		Use:   "register [<name>]",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.PrintErrln("must provide <name> when not running interactively")
				cmd.Usage()
				return
			}

			if domain == "" {
				cmd.PrintErrln("--domain missing")
				os.Exit(1)
			}

			opt.ClientName = args[0]

			if len(opt.RedirectURIs) == 0 || opt.SoftwareID == "" {
				cmd.PrintErrln("must provide --redirect-uris and --software-id when not running interactively")
				cmd.Usage()
				return
			}

			res, err := tb.Client(domain).Register(cmd.Context(), &opt)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			fmt.Printf("Client %q created in %s\n", res.ClientName, domain)
			err = tb.Config().SaveClient(res)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}
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

	cmd.Flags().StringVar(&domain, "domain", "", "Domain to contact (example: \"foobar.mycozy.cloud\")")

	return &cmd
}
