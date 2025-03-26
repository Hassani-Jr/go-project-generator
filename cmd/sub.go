package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:     "subtract",
	Short:   "Subtracts 2 numbers from each other",
	Long:    "Subtracts 2 integers",
	Aliases: []string{"sub"},
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Subtraction of %s from %s = %s.\n\n", args[1], args[0], Subtract(args[0], args[1]))
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
