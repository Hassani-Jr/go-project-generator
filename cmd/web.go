package cmd

import (
	"github.com/spf13/cobra"
)

var webFrameworkCmd = &cobra.Command{
	Use:     "web",
	Aliases: []string{"Web", "WEB"},
	Short:   "Subcommand for generating a go web framework tool",
	Long:    "Subcommand for generating a basic web framework setup in GO",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(webFrameworkCmd)
}
