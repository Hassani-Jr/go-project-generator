package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	outputDir string
	gitInit   bool
	verbose   bool
)

var rootCmd = &cobra.Command{
	Use:   "go-project-generator",
	Short: "Generate Go project structures quickly",
	Long: `Go Project Generator is a powerful CLI tool designed to streamline 
the process of creating new Go projects. It provides quick, consistent 
project scaffolding for various project types.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", ".", "Output directory for the project")
	rootCmd.PersistentFlags().BoolVarP(&gitInit, "git", "g", true, "Initialize git repository")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}
