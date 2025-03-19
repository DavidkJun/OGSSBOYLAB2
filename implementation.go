package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

// PrefixToInfix converts a prefix expression to an infix expression.
func PrefixToInfix(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("empty input")
	}

	tokens := strings.Fields(input)
	stack := []string{}

	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]

		if isOperator(token) {
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression")
			}

			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			infix := fmt.Sprintf("(%s %s %s)", operand1, token, operand2)
			stack = append(stack, infix)
		} else {
			// Перевірка, чи токен є числом
			_, err := strconv.Atoi(token)
			if err != nil {
				return "", fmt.Errorf("invalid token: %s", token)
			}
			stack = append(stack, token)
		}
	}

	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression")
	}

	return stack[0], nil
}

// isOperator перевіряє, чи є токен оператором.
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/" || token == "^"
}
