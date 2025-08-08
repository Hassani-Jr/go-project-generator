package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Hassani-Jr/go-project-generator/internal/generator"
	"github.com/spf13/cobra"
)

var microserviceCmd = &cobra.Command{
	Use:     "microservice [project-name]",
	Aliases: []string{"Microservice", "MICROSERVICE", "micro"},
	Short:   "Generate a microservice",
	Long:    "Generate a microservice with gRPC support and standard structure",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		projectPath := filepath.Join(outputDir, projectName)

		if verbose {
			fmt.Printf("Creating microservice project: %s\n", projectName)
			fmt.Printf("Output directory: %s\n", projectPath)
		}

		config := generator.ProjectConfig{
			ProjectName: projectName,
			ProjectPath: projectPath,
			ProjectType: "microservice",
			GitInit:     gitInit,
		}

		gen := generator.New(config)
		if err := gen.Generate(); err != nil {
			log.Fatalf("Failed to generate microservice project: %v", err)
		}

		fmt.Printf("✅ Microservice project '%s' created successfully!\n", projectName)

		if gitInit {
			initGitRepo(projectPath)
		}

		printNextSteps(projectName)
	},
}

func init() {
	rootCmd.AddCommand(microserviceCmd)
}
