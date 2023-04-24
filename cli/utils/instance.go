package utils

import (
	"os"

	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/Peltoche/gozy/sdk/instance"
	"github.com/spf13/cobra"
)

func GetInstance(cmd *cobra.Command, tb toolbox.Toolbox) *instance.Instance {
	instanceFlg := cmd.Flag("instance").Value.String()

	cfg, err := tb.Config().Read()
	if (err != nil || cfg.Instance == "") && instanceFlg == "" {
		cmd.Println("no instance specified, please use the --instance flag")
		os.Exit(1)
	}

	if instanceFlg == "" {
		instanceFlg = cfg.Instance
	}

	res, err := instance.NewFromStr(instanceFlg)
	if err != nil {
		cmd.Printf("invalid instance: %s\n", err)
		os.Exit(1)
	}

	return res
}
