package cmd

import (
	"github.com/spf13/cobra"
)

var toolCmd = &cobra.Command{
	Use:     "tool",
	Aliases: []string{"TOOL", "Tool"},
	Short:   "Subcommand for generating a tool in GO",
	Long:    "Subcommand for generating a basic GO tool",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(toolCmd)
}
