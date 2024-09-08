package main_test

import (
    "testing"
	"os/exec"
	"regexp"
	"strings"
)

func TestRunWithValidUsername(t *testing.T) {
	cmd := exec.Command("go", "run", "..", "selop")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed to run: %v\nOutput: %s", err, output)
	}

	// Check if the output contains expected activity types
	expectedPatterns := []string{
		"🔨 Pushed commits to",
		"🐛 .* an issue in",
		"⭐ Starred",
		"🔔 .* on",
	}

	for _, pattern := range expectedPatterns {
		matched, err := regexp.Match(pattern, output)
		if err != nil {
			t.Fatalf("Error matching pattern '%s': %v", pattern, err)
		}
		if !matched {
			t.Errorf("Expected output to match pattern '%s', but it didn't.\nOutput: %s", pattern, output)
		}
	}
}

func TestRunWithInvalidUsername(t *testing.T) {
	cmd := exec.Command("go", "run", "..", "invalid-user")
	output, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("Expected command to fail, but it didn't.\nOutput: %s", output)
	}

	// Check if the error message is as expected
	expectedError := "No recent activity found for user 'invalid-user'"
	if !strings.Contains(string(output), expectedError) {
		t.Errorf("Expected error message '%s', but got '%s'", expectedError, string(output))
	}
}
