package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var shouldRoundUp bool
var multiplyCmd = &cobra.Command{
	Use:     "multiply",
	Aliases: []string{"mult"},
	Short:   "Multiplies 2 ints",
	Long:    "Multiplies 2 ints together",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Multiplication of %s and %s = %s.\n\n", args[0], args[1], Multiply(args[0], args[1], shouldRoundUp))
	},
}

func init() {
	multiplyCmd.Flags().BoolVarP(&shouldRoundUp, "round", "r", false, "Round results up to 2 decimal places")
	rootCmd.AddCommand(multiplyCmd)
}
