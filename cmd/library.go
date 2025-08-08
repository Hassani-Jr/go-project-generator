package cmd

import (
	"fmt"
	"github.com/Hassani-Jr/go-project-generator/internal/generator"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var libraryCmd = &cobra.Command{
	Use:     "library [project-name]",
	Aliases: []string{"LIB", "lib", "Library", "LIBRARY"},
	Short:   "Generate a Go library",
	Long:    "Generate a Go library with examples, tests, and standard structure",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectPath := filepath.Join(outputDir, projectName)

		if verbose {
			fmt.Printf("Creating library project: %s\n", projectName)
			fmt.Printf("Output directory: %s\n", projectPath)
		}

		config := generator.ProjectConfig{
			ProjectName: projectName,
			ProjectPath: projectPath,
			ProjectType: "library",
			GitInit:     gitInit,
		}

		gen := generator.New(config)
		if err := gen.Generate(); err != nil {
			log.Fatalf("Failed to generate library project: %v", err)
		}

		fmt.Printf("✅ Library project '%s' created successfully!\n", projectName)

		if gitInit {
			initGitRepo(projectPath)
		}

		printNextSteps(projectName)
	},
}

func init() {
	rootCmd.AddCommand(libraryCmd)
}
