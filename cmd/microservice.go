package cmd

import (
	"github.com/spf13/cobra"
)

var microServiceCmd = &cobra.Command{
	Use:     "microservice",
	Aliases: []string{"Microservice", "MICROSERVICE"},
	Short:   "Subcommand for generating a microservice in go",
	Long:    "Subcommand for generating a basic microservice in GO",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(microServiceCmd)
}
