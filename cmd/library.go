package cmd

import (
	"github.com/spf13/cobra"
)

var libraryCmd = &cobra.Command{
	Use:     "library",
	Aliases: []string{"LIB", "lib", "Library", "LIBRARY"},
	Short:   "Subcommand for generating a library in GO",
	Long:    "Subcommand for generating a basic library in GO",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(libraryCmd)
}
