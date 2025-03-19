package lab2

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestComputeHandler_Compute tests the Compute method of ComputeHandler.
func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr error
	}{
		{
			name:        "valid expression: simple addition",
			input:       "+ 2 3",
			expected:    "(2 + 3)",
			expectedErr: nil,
		},
		{
			name:        "valid expression: complex expression",
			input:       "* + 2 3 - 4 1",
			expected:    "((2 + 3) * (4 - 1))",
			expectedErr: nil,
		},
		{
			name:        "invalid expression: insufficient operands",
			input:       "+ 2",
			expected:    "",
			expectedErr: errors.New("invalid expression"),
		},
		{
			name:        "invalid expression: empty input",
			input:       "",
			expected:    "",
			expectedErr: errors.New("empty input"),
		},
		{
			name:        "invalid expression: invalid token",
			input:       "+ 2 a",
			expected:    "",
			expectedErr: errors.New("invalid token: a"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			output := &bytes.Buffer{}

			handler := &ComputeHandler{
				Input:  input,
				Output: output,
			}

			err := handler.Compute()

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error(), "Error should match expected error")
			} else {
				assert.Nil(t, err, "Error should be nil for valid input")
			}

			assert.Equal(t, tt.expected, output.String(), "Output should match expected result")
		})
	}
}
