package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {
	tests := []struct {
		name        string
		projectType string
		wantErr     bool
		checkFiles  []string
	}{
		{
			name:        "CLI project",
			projectType: "cli",
			wantErr:     false,
			checkFiles: []string{
				"main.go",
				"cmd/root.go",
				"internal/commands/example.go",
				"go.mod",
				"README.md",
			},
		},
		{
			name:        "Web service project",
			projectType: "web",
			wantErr:     false,
			checkFiles: []string{
				"main.go",
				"internal/handlers/handlers.go",
				"internal/middleware/middleware.go",
				"configs/config.yaml",
				"go.mod",
				"README.md",
			},
		},
		{
			name:        "Microservice project",
			projectType: "microservice",
			wantErr:     false,
			checkFiles: []string{
				"main.go",
				"internal/service/service.go",
				"internal/proto/service.proto",
				"scripts/proto-gen.sh",
				"go.mod",
				"README.md",
			},
		},
		{
			name:        "Library project",
			projectType: "library",
			wantErr:     false,
			checkFiles: []string{
				"pkg/test-project/test-project.go",
				"pkg/test-project/test-project_test.go",
				"examples/example.go",
				"LICENSE",
				"go.mod",
				"README.md",
			},
		},
		{
			name:        "Tool project",
			projectType: "tool",
			wantErr:     false,
			checkFiles: []string{
				"main.go",
				"cmd/commands/process.go",
				"cmd/commands/analyze.go",
				"internal/utils/utils.go",
				"go.mod",
				"README.md",
			},
		},
		{
			name:        "Invalid project type",
			projectType: "invalid",
			wantErr:     true,
			checkFiles:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary directory for test
			tempDir, err := os.MkdirTemp("", "generator-test-*")
			if err != nil {
				t.Fatalf("Failed to create temp dir: %v", err)
			}
			defer func(path string) {
				err := os.RemoveAll(path)
				if err != nil {
					fmt.Println(err)
				}
			}(tempDir)

			projectPath := filepath.Join(tempDir, "test-project")

			config := ProjectConfig{
				ProjectName: "test-project",
				ProjectPath: projectPath,
				ProjectType: tt.projectType,
				GitInit:     false,
			}

			gen := New(config)
			err = gen.Generate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Check if expected files were created
				for _, file := range tt.checkFiles {
					filePath := filepath.Join(projectPath, file)
					if _, err := os.Stat(filePath); os.IsNotExist(err) {
						t.Errorf("Expected file %s was not created", file)
					}
				}
			}
		})
	}
}

func TestGenerator_createDir(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "generator-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
		}
	}(tempDir)

	gen := &Generator{
		Config: ProjectConfig{
			ProjectPath: tempDir,
		},
	}

	testDir := "test/nested/directory"
	err = gen.createDir(testDir)
	if err != nil {
		t.Errorf("createDir() error = %v", err)
		return
	}

	fullPath := filepath.Join(tempDir, testDir)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		t.Errorf("Directory %s was not created", testDir)
	}
}

func TestGenerator_createFile(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "generator-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
		}
	}(tempDir)

	gen := &Generator{
		Config: ProjectConfig{
			ProjectPath: tempDir,
		},
	}

	testFile := "test/file.txt"
	testContent := "test content"

	err = gen.createFile(testFile, testContent)
	if err != nil {
		t.Errorf("createFile() error = %v", err)
		return
	}

	fullPath := filepath.Join(tempDir, testFile)
	content, err := os.ReadFile(fullPath)
	if err != nil {
		t.Errorf("Failed to read created file: %v", err)
		return
	}

	if string(content) != testContent {
		t.Errorf("File content = %v, want %v", string(content), testContent)
	}
}
