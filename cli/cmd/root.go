/*
Copyright © 2023 Peltoche pierre.peltier@protonmail.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"os"

	"github.com/Peltoche/gozy/cli/cmd/auth"
	"github.com/Peltoche/gozy/cli/cmd/client"
	"github.com/Peltoche/gozy/cli/cmd/instance"
	"github.com/Peltoche/gozy/cli/utils/toolbox"
	"github.com/spf13/cobra"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cmd := &cobra.Command{
		Use:   "gozy",
		Short: "Manage your cozy instance in your terminal.",
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	// Generic flags

	tb := toolbox.NewProd()

	// Subcommands
	cmd.AddCommand(auth.NewAuthCmd(tb))
	cmd.AddCommand(client.NewClientCmd(tb))
	cmd.AddCommand(instance.NewInstanceCmd(tb))

	cmd.PersistentFlags().StringP("instance", "I", "", "Instance to contact (example: \"foobar.mycozy.cloud\")")

	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
