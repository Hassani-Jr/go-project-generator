package cmd

import (
	"testing"
)

func TestRootCommand(t *testing.T) {
	// Test that root command exists and can be executed
	if rootCmd == nil {
		t.Error("rootCmd is nil")
	}

	// Test command name
	if rootCmd.Use != "go-project-generator" {
		t.Errorf("rootCmd.Use = %v, want %v", rootCmd.Use, "go-project-generator")
	}
}

func TestSubCommands(t *testing.T) {
	subCommands := []string{"cli", "web", "microservice", "library", "tool"}

	for _, cmdName := range subCommands {
		cmd, _, err := rootCmd.Find([]string{cmdName})
		if err != nil {
			t.Errorf("Command %s not found: %v", cmdName, err)
			continue
		}

		if cmd == nil {
			t.Errorf("Command %s is nil", cmdName)
		}
	}
}
