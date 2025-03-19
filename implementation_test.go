package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToInfix(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			name:     "simple addition",
			input:    "+ 2 3",
			expected: "(2 + 3)",
			err:      nil,
		},
		{
			name:     "complex expression",
			input:    "* + 2 3 - 4 1",
			expected: "((2 + 3) * (4 - 1))",
			err:      nil,
		},
		{
			name:     "invalid expression",
			input:    "+ 2",
			expected: "",
			err:      fmt.Errorf("invalid expression"),
		},
		{
			name:     "invalid token",
			input:    "+ 2 a",
			expected: "",
			err:      fmt.Errorf("invalid token: a"),
		},
		{
			name:     "empty input",
			input:    "",
			expected: "",
			err:      fmt.Errorf("empty input"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := PrefixToInfix(tt.input)
			if tt.err != nil {
				fmt.Println(tt.err)
				assert.EqualError(t, err, tt.err.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.expected, res)
			}
		})
	}
}

func ExamplePrefixToInfix() {
	tests := []struct {
		input    string
		expected string
	}{
		{"+ 2 3", "(2 + 3)"},
		{"* + 2 3 - 4 1", "((2 + 3) * (4 - 1))"},
		{"+ 2", ""},
		{"+ 2 a", ""},
		{"", ""},
	}

	for _, tt := range tests {
		res, _ := PrefixToInfix(tt.input)
		fmt.Println(res)
	}

	// Output:
	// (2 + 3)
	// ((2 + 3) * (4 - 1))
	//
	//
	//
}
