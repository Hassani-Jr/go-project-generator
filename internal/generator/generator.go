package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type ProjectConfig struct {
	ProjectName string
	ProjectPath string
	ProjectType string
	GitInit     bool
}

type Generator struct {
	Config ProjectConfig
}

func New(config ProjectConfig) *Generator {
	return &Generator{Config: config}
}

func (g *Generator) Generate() error {
	switch g.Config.ProjectType {
	case "cli":
		return g.generateCLI()
	case "web":
		return g.generateWeb()
	case "microservice":
		return g.generateMicroservice()
	case "library":
		return g.generateLibrary()
	case "tool":
		return g.generateTool()
	default:
		return fmt.Errorf("unknown project type: %s", g.Config.ProjectType)
	}
}

func (g *Generator) createDir(path string) error {
	return os.MkdirAll(filepath.Join(g.Config.ProjectPath, path), 0755)
}

func (g *Generator) createFile(path, content string) error {
	fullPath := filepath.Join(g.Config.ProjectPath, path)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(fullPath, []byte(content), 0644)
}

func (g *Generator) generateCLI() error {
	// Create directory structure
	dirs := []string{
		"cmd",
		"internal/commands",
		"pkg",
		"configs",
		"scripts",
	}

	for _, dir := range dirs {
		if err := g.createDir(dir); err != nil {
			return err
		}
	}

	// Create main.go
	mainContent := `package main

import (
	"fmt"
	"os"

	"{{.ProjectName}}/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
`
	if err := g.createFileFromTemplate("main.go", mainContent, g.Config); err != nil {
		return err
	}

	// Create cmd/root.go
	rootContent := `package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "{{.ProjectName}}",
	Short: "A brief description of your CLI application",
	Long:  "A longer description of your CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to {{.ProjectName}}!")
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
}
`
	if err := g.createFileFromTemplate("cmd/root.go", rootContent, g.Config); err != nil {
		return err
	}

	// Create internal/commands/example.go
	exampleCommand := `package commands

import (
	"fmt"
)

type ExampleCommand struct {
	Name string
}

func NewExampleCommand(name string) *ExampleCommand {
	return &ExampleCommand{Name: name}
}

func (c *ExampleCommand) Run() error {
	fmt.Printf("Running example command with name: %s\n", c.Name)
	return nil
}
`
	if err := g.createFile("internal/commands/example.go", exampleCommand); err != nil {
		return err
	}

	// Create go.mod
	if err := g.createGoMod(); err != nil {
		return err
	}

	// Create README
	if err := g.createReadme("CLI"); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateWeb() error {
	// Create directory structure
	dirs := []string{
		"cmd/server",
		"internal/handlers",
		"internal/middleware",
		"internal/models",
		"internal/services",
		"pkg/database",
		"pkg/config",
		"configs",
		"scripts/migrations",
		"api",
	}

	for _, dir := range dirs {
		if err := g.createDir(dir); err != nil {
			return err
		}
	}

	// Create main.go
	mainContent := `package main

import (
	"log"
	"net/http"
	"os"

	"{{.ProjectName}}/internal/handlers"
	"{{.ProjectName}}/internal/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	
	// Setup routes
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/api/v1/", handlers.APIHandler)
	
	// Apply middleware
	handler := middleware.Logging(middleware.CORS(mux))
	
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
`
	if err := g.createFileFromTemplate("main.go", mainContent, g.Config); err != nil {
		return err
	}

	// Create handlers
	healthHandler := `package handlers

import (
	"encoding/json"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "API endpoint",
		"version": "v1",
	})
}
`
	if err := g.createFile("internal/handlers/handlers.go", healthHandler); err != nil {
		return err
	}

	// Create middleware
	middlewareContent := `package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}
`
	if err := g.createFile("internal/middleware/middleware.go", middlewareContent); err != nil {
		return err
	}

	// Create config.yaml
	configContent := `server:
  port: 8080
  host: localhost

database:
  host: localhost
  port: 5432
  name: {{.ProjectName}}_db
  user: postgres
  password: password

app:
  name: {{.ProjectName}}
  version: 1.0.0
  environment: development
`
	if err := g.createFileFromTemplate("configs/config.yaml", configContent, g.Config); err != nil {
		return err
	}

	// Create go.mod
	if err := g.createGoMod(); err != nil {
		return err
	}

	// Create README
	if err := g.createReadme("Web Service"); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateMicroservice() error {
	// Create directory structure
	dirs := []string{
		"cmd/server",
		"internal/service",
		"internal/proto",
		"pkg/interceptors",
		"pkg/client",
		"configs",
		"scripts",
		"api",
	}

	for _, dir := range dirs {
		if err := g.createDir(dir); err != nil {
			return err
		}
	}

	// Create main.go
	mainContent := `package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"{{.ProjectName}}/internal/service"
)

func main() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	
	// Register your services here
	svc := service.NewService()
	// Example: pb.RegisterYourServiceServer(grpcServer, svc)
	
	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		log.Println("Shutting down gRPC server...")
		grpcServer.GracefulStop()
	}()

	log.Printf("gRPC server listening on port %s", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
`
	if err := g.createFileFromTemplate("main.go", mainContent, g.Config); err != nil {
		return err
	}

	// Create service implementation
	serviceContent := `package service

import (
	"context"
	"log"
)

type Service struct {
	// Add your service fields here
}

func NewService() *Service {
	return &Service{}
}

// Implement your gRPC service methods here
func (s *Service) ExampleMethod(ctx context.Context) error {
	log.Println("ExampleMethod called")
	return nil
}
`
	if err := g.createFile("internal/service/service.go", serviceContent); err != nil {
		return err
	}

	// Create proto file
	protoContent := `syntax = "proto3";

package {{.ProjectName}};

option go_package = "{{.ProjectName}}/internal/proto";

service ExampleService {
    rpc ExampleMethod(ExampleRequest) returns (ExampleResponse);
}

message ExampleRequest {
    string id = 1;
    string data = 2;
}

message ExampleResponse {
    string result = 1;
    bool success = 2;
}
`
	if err := g.createFileFromTemplate("internal/proto/service.proto", protoContent, g.Config); err != nil {
		return err
	}

	// Create proto generation script
	protoGenScript := `#!/bin/bash

# Generate Go code from proto files
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/proto/*.proto

echo "Proto files generated successfully"
`
	if err := g.createFile("scripts/proto-gen.sh", protoGenScript); err != nil {
		err := os.Chmod(filepath.Join(g.Config.ProjectPath, "scripts/proto-gen.sh"), 0755)
		if err != nil {
			return err
		}
		return err
	}

	// Make the script executable
	scriptPath := filepath.Join(g.Config.ProjectPath, "scripts/proto-gen.sh")
	if err := os.Chmod(scriptPath, 0755); err != nil {
		// Log warning but don't fail - chmod might not work on all systems
		fmt.Printf("Warning: Could not make proto-gen.sh executable: %v\n", err)
	}

	// Create go.mod with gRPC dependencies
	if err := g.createGoModWithDeps([]string{
		"google.golang.org/grpc",
		"google.golang.org/protobuf",
	}); err != nil {
		return err
	}

	// Create README
	if err := g.createReadme("Microservice"); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateLibrary() error {
	// Create directory structure
	dirs := []string{
		"pkg/" + g.Config.ProjectName,
		"examples",
		"internal/helpers",
		"scripts",
		"docs",
	}

	for _, dir := range dirs {
		if err := g.createDir(dir); err != nil {
			return err
		}
	}

	// Create main library file
	libContent := `package {{.ProjectName}}

import (
	"fmt"
)

// Version represents the library version
const Version = "1.0.0"

// Config holds the library configuration
type Config struct {
	// Add your configuration fields here
	Debug bool
}

// Client represents the main library client
type Client struct {
	config *Config
}

// New creates a new instance of the library client
func New(config *Config) *Client {
	if config == nil {
		config = &Config{}
	}
	return &Client{config: config}
}

// ExampleMethod is an example public method
func (c *Client) ExampleMethod(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input cannot be empty")
	}
	return fmt.Sprintf("Processed: %s", input), nil
}
`
	libPath := fmt.Sprintf("pkg/%s/%s.go", g.Config.ProjectName, g.Config.ProjectName)
	if err := g.createFileFromTemplate(libPath, libContent, g.Config); err != nil {
		return err
	}

	// Create example usage
	exampleContent := `package main

import (
	"fmt"
	"log"

	"{{.ProjectName}}/pkg/{{.ProjectName}}"
)

func main() {
	// Create a new client
	client := {{.ProjectName}}.New(&{{.ProjectName}}.Config{
		Debug: true,
	})

	// Use the library
	result, err := client.ExampleMethod("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
`
	if err := g.createFileFromTemplate("examples/example.go", exampleContent, g.Config); err != nil {
		return err
	}

	// Create test file
	testContent := `package {{.ProjectName}}_test

import (
	"testing"

	"{{.ProjectName}}/pkg/{{.ProjectName}}"
)

func TestExampleMethod(t *testing.T) {
	client := {{.ProjectName}}.New(nil)

	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:    "valid input",
			input:   "test",
			want:    "Processed: test",
			wantErr: false,
		},
		{
			name:    "empty input",
			input:   "",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.ExampleMethod(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExampleMethod() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExampleMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}
`
	testPath := fmt.Sprintf("pkg/%s/%s_test.go", g.Config.ProjectName, g.Config.ProjectName)
	if err := g.createFileFromTemplate(testPath, testContent, g.Config); err != nil {
		return err
	}

	// Create LICENSE
	licenseContent := `MIT License

Copyright (c) 2024 {{.ProjectName}}

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
`
	if err := g.createFileFromTemplate("LICENSE", licenseContent, g.Config); err != nil {
		return err
	}

	// Create go.mod
	if err := g.createGoMod(); err != nil {
		return err
	}

	// Create README
	if err := g.createReadme("Library"); err != nil {
		return err
	}

	return nil
}

func (g *Generator) generateTool() error {
	// Create directory structure
	dirs := []string{
		"cmd/commands",
		"internal/utils",
		"pkg",
		"configs",
		"scripts",
	}

	for _, dir := range dirs {
		if err := g.createDir(dir); err != nil {
			return err
		}
	}

	// Create main.go
	mainContent := `package main

import (
	"flag"
	"fmt"
	"os"

	"{{.ProjectName}}/cmd/commands"
)

func main() {
	var (
		verbose = flag.Bool("v", false, "verbose output")
		help    = flag.Bool("h", false, "show help")
	)

	flag.Parse()

	if *help || len(flag.Args()) == 0 {
		printUsage()
		os.Exit(0)
	}

	command := flag.Arg(0)
	args := flag.Args()[1:]

	switch command {
	case "process":
		cmd := commands.NewProcessCommand(*verbose)
		if err := cmd.Run(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	case "analyze":
		cmd := commands.NewAnalyzeCommand(*verbose)
		if err := cmd.Run(args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: {{.ProjectName}} [options] <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  process    Process input data")
	fmt.Println("  analyze    Analyze input data")
	fmt.Println("\nOptions:")
	fmt.Println("  -v         Verbose output")
	fmt.Println("  -h         Show this help message")
}
`
	if err := g.createFileFromTemplate("main.go", mainContent, g.Config); err != nil {
		return err
	}

	// Create process command
	processCmd := `package commands

import (
	"fmt"
)

type ProcessCommand struct {
	verbose bool
}

func NewProcessCommand(verbose bool) *ProcessCommand {
	return &ProcessCommand{verbose: verbose}
}

func (c *ProcessCommand) Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no input provided")
	}

	if c.verbose {
		fmt.Println("Processing input...")
	}

	for _, arg := range args {
		fmt.Printf("Processing: %s\n", arg)
		// Add your processing logic here
	}

	return nil
}
`
	if err := g.createFile("cmd/commands/process.go", processCmd); err != nil {
		return err
	}

	// Create analyze command
	analyzeCmd := `package commands

import (
	"fmt"
)

type AnalyzeCommand struct {
	verbose bool
}

func NewAnalyzeCommand(verbose bool) *AnalyzeCommand {
	return &AnalyzeCommand{verbose: verbose}
}

func (c *AnalyzeCommand) Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no input provided for analysis")
	}

	if c.verbose {
		fmt.Println("Analyzing input...")
	}

	for _, arg := range args {
		fmt.Printf("Analyzing: %s\n", arg)
		// Add your analysis logic here
	}

	return nil
}
`
	if err := g.createFile("cmd/commands/analyze.go", analyzeCmd); err != nil {
		return err
	}

	// Create utils
	utilsContent := `package utils

import (
	"os"
	"path/filepath"
)

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// EnsureDir creates a directory if it doesn't exist
func EnsureDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// GetExecutablePath returns the path of the current executable
func GetExecutablePath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(ex), nil
}
`
	if err := g.createFile("internal/utils/utils.go", utilsContent); err != nil {
		return err
	}

	// Create go.mod
	if err := g.createGoMod(); err != nil {
		return err
	}

	// Create README
	if err := g.createReadme("Tool"); err != nil {
		return err
	}

	return nil
}

func (g *Generator) createGoMod() error {
	content := `module {{.ProjectName}}

go 1.21

require (
	github.com/spf13/cobra v1.8.0
)
`
	return g.createFileFromTemplate("go.mod", content, g.Config)
}

func (g *Generator) createGoModWithDeps(deps []string) error {
	content := `module {{.ProjectName}}

go 1.21

require (
`
	for _, dep := range deps {
		content += fmt.Sprintf("\t%s v1.65.0\n", dep)
	}
	content += ")\n"

	return g.createFileFromTemplate("go.mod", content, g.Config)
}

func (g *Generator) createReadme(projectType string) error {
	content := `# {{.ProjectName}}

## Description

This is a {{.ProjectType}} project generated with go-project-generator.

## Installation

` + "```bash\ngo mod download\n```" + `

## Usage

### Running the application
` + "```bash\ngo run main.go\n```" + `

### Building
` + "```bash\ngo build -o {{.ProjectName}}\n```" + `

### Testing
` + "```bash\ngo test ./...\n```" + `

## Project Structure

See the project structure for {{.ProjectType}} projects in the go-project-generator documentation.

## Contributing

1. Fork the repository
2. Create your feature branch (` + "`git checkout -b feature/amazing-feature`" + `)
3. Commit your changes (` + "`git commit -m 'Add some amazing feature'`" + `)
4. Push to the branch (` + "`git push origin feature/amazing-feature`" + `)
5. Open a Pull Request

## License

This project is licensed under the MIT License.
`
	data := struct {
		ProjectName string
		ProjectType string
	}{
		ProjectName: g.Config.ProjectName,
		ProjectType: projectType,
	}

	return g.createFileFromTemplate("README.md", content, data)
}

func (g *Generator) createFileFromTemplate(path, templateContent string, data interface{}) error {
	tmpl, err := template.New(path).Parse(templateContent)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(g.Config.ProjectPath, path)
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	return tmpl.Execute(file, data)
}
