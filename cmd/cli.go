package cmd

import (
	"fmt"
	"github.com/Hassani-Jr/go-project-generator/internal/generator"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var cliCmd = &cobra.Command{
	Use:     "cli [project-name]",
	Aliases: []string{"CLI", "Cli"},
	Short:   "Generate a CLI application",
	Long:    "Generate a CLI application with cobra integration and standard structure",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectPath := filepath.Join(outputDir, projectName)

		if verbose {
			fmt.Printf("Creating CLI project: %s\n", projectName)
			fmt.Printf("Output directory: %s\n", projectPath)
		}

		config := generator.ProjectConfig{
			ProjectName: projectName,
			ProjectPath: projectPath,
			ProjectType: "cli",
			GitInit:     gitInit,
		}

		gen := generator.New(config)
		if err := gen.Generate(); err != nil {
			log.Fatalf("Failed to generate CLI project: %v", err)
		}

		fmt.Printf("✅ CLI project '%s' created successfully!\n", projectName)

		if gitInit {
			initGitRepo(projectPath)
		}

		printNextSteps(projectName)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
}
