package cmd

import (
	"github.com/spf13/cobra"
)

var cliCmd = &cobra.Command{
	Use:     "cli",
	Aliases: []string{"CLI", "Cli"},
	Short:   "Subcommand for generating a cli tool",
	Long:    "Subcommand for generating a basic CLI in GO",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
