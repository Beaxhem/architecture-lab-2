package lab2

import (
	"bufio"
	"fmt"
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

// Read the expression from input and write the computed result to the output.
func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Reader)
	for scanner.Scan() {
		input := scanner.Text()
		result, err := PostfixToPrefix(input)
		if err != nil {
			return err
		}
		fmt.Fprintln(ch.Writer, result)
	}

	return nil
}
