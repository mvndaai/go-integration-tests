package integration_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func runIntegration(args ...string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	mainPath := filepath.Join(wd, "..", "main.go")
	args = append([]string{"run", mainPath}, args...)
	cmd := exec.Command("go", args...)
	b, err := cmd.CombinedOutput()
	return string(b), err
}

func TestIntegration(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "empty string",
			in:       "",
			expected: "\n",
		},
		{
			name:     "emoji",
			in:       "👍",
			expected: "👍\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := runIntegration(tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
