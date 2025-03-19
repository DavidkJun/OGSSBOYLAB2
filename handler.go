package lab2

import (
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := io.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	expression := string(data)
	result, err := PrefixToInfix(expression)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(result))
	return err
}
