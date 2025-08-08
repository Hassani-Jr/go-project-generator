package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Hassani-Jr/go-project-generator/internal/generator"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:     "web [project-name]",
	Aliases: []string{"Web", "WEB", "webservice"},
	Short:   "Generate a web service",
	Long:    "Generate a web service with HTTP handlers, middleware, and standard structure",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectPath := filepath.Join(outputDir, projectName)

		if verbose {
			fmt.Printf("Creating web service project: %s\n", projectName)
			fmt.Printf("Output directory: %s\n", projectPath)
		}

		config := generator.ProjectConfig{
			ProjectName: projectName,
			ProjectPath: projectPath,
			ProjectType: "web",
			GitInit:     gitInit,
		}

		gen := generator.New(config)
		if err := gen.Generate(); err != nil {
			log.Fatalf("Failed to generate web service project: %v", err)
		}

		fmt.Printf("✅ Web service project '%s' created successfully!\n", projectName)

		if gitInit {
			initGitRepo(projectPath)
		}

		printNextSteps(projectName)
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
