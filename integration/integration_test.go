package integration_test

import (
	"testing"

	"github.com/mvndaai/go-integration-tests/integration"
	"github.com/stretchr/testify/assert"
)

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
			actual, err := integration.RunIntegration(tt.in)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
