package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func initGitRepo(projectPath string) {
	if verbose {
		fmt.Println("Initializing git repository...")
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = projectPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to initialize git repository: %v\n", err)
		return
	}

	// Create .gitignore
	gitignoreContent := `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# Go workspace file
go.work

# IDE directories
.idea/
.vscode/
*.swp
*.swo
*~

# OS files
.DS_Store
Thumbs.db

# Environment variables
.env
.env.local

# Build directories
dist/
build/
`
	gitignorePath := filepath.Join(projectPath, ".gitignore")
	if err := os.WriteFile(gitignorePath, []byte(gitignoreContent), 0644); err != nil {
		fmt.Printf("Warning: Failed to create .gitignore: %v\n", err)
		return
	}

	// Initial commit
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to stage files: %v\n", err)
		return
	}

	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: Failed to create initial commit: %v\n", err)
		return
	}

	if verbose {
		fmt.Println("Git repository initialized with initial commit")
	}
}

func printNextSteps(projectName string) {
	fmt.Println("\n📝 Next steps:")
	fmt.Printf("   cd %s\n", projectName)
	fmt.Println("   go mod tidy")
	fmt.Println("   go run main.go")
	fmt.Println("\n📚 For more information, check the README.md file in your project")
}
