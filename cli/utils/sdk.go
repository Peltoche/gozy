package utils

import (
	"os"

	"github.com/Peltoche/gozy/sdk"
	"github.com/spf13/cobra"
)

func NewClient(cmd *cobra.Command) *sdk.SDK {
	domain := cmd.Flag("domain").Value.String()
	if domain == "" {
		cmd.PrintErrln("--domain missing")
		os.Exit(1)
	}

	return sdk.NewSDK(domain)
}
