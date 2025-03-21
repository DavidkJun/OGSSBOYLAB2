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
		// Прості вирази з 2-3 операндами
		{
			name:     "simple addition",
			input:    "+ 2 4",
			expected: "(2 + 5)",
			err:      nil,
		},
		{
			name:     "simple subtraction",
			input:    "- 5 2",
			expected: "(5 - 2)",
			err:      nil,
		},
		{
			name:     "simple multiplication",
			input:    "* 3 4",
			expected: "(3 * 4)",
			err:      nil,
		},
		{
			name:     "simple division",
			input:    "/ 10 2",
			expected: "(10 / 2)",
			err:      nil,
		},
		{
			name:     "simple exponentiation",
			input:    "^ 2 3",
			expected: "(2 ^ 3)",
			err:      nil,
		},
		{
			name:     "complex expression",
			input:    "+ * - 5 2 3 / ^ 2 3 2",
			expected: "(((5 - 2) * 3) + ((2 ^ 3) / 2))",
			err:      nil,
		},
		{
			name:     "complex expression with 9 operands",
			input:    "+ * - 1 2 3 + / - 4 5 6 * 7 * 8 9",
			expected: "(((1 - 2) * 3) + (((4 - 5) / 6) + (7 * (8 * 9))))",
			err:      nil,
		},
		{
			name:     "invalid expression: insufficient operands",
			input:    "+ 2",
			expected: "",
			err:      fmt.Errorf("invalid expression"),
		},
		{
			name:     "invalid expression: empty input",
			input:    "",
			expected: "",
			err:      fmt.Errorf("empty input"),
		},
		{
			name:     "invalid expression: invalid token",
			input:    "+ 2 a",
			expected: "",
			err:      fmt.Errorf("invalid token: a"),
		},
		{
			name:     "invalid expression: invalid operator",
			input:    "~ 2 3",
			expected: "",
			err:      fmt.Errorf("invalid token: ~"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := PrefixToInfix(tt.input)
			if tt.err != nil {
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
		{"+ 2 4", "(2 + 4)"},
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
	// (2 + 4)
	// ((2 + 3) * (4 - 1))
	//
	//
	//
}
