package echo_test

import (
	"testing"

	"github.com/mvndaai/go-integration-tests/echo"
	"gotest.tools/assert"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "empty string",
			in:       "",
			expected: "",
		},
		{
			name:     "emoji",
			in:       "👍",
			expected: "👍",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acutal := echo.Echo(tt.in)
			assert.Equal(t, tt.expected, acutal)
		})
	}
}
