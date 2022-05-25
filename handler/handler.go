package handler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Beaxhem/architecture-lab-2/implementation"
)

// Struct incupsulating
type ComputeHandler struct {
	Reader io.Reader
	Writer io.Writer
}

func NewComputeHandler(inputExpression *string, inputPath *string, outputPath *string) (*ComputeHandler, error) {
	reader, err := reader(inputExpression, inputPath)
	if err != nil {
		return nil, err
	}
	writer, err := writer(outputPath)
	if err != nil {
		return nil, err
	}
	return &ComputeHandler{
		Reader: reader,
		Writer: writer,
	}, nil
}

// Read the expression from input and write the computed result to the output.
func (ch *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(ch.Reader)
	for scanner.Scan() {
		input := scanner.Text()
		result, err := lab2.PostfixToPrefix(input)
		if err != nil {
			return err
		}
		fmt.Fprintln(ch.Writer, result)
	}

	return nil
}

// Returns io.Reader depending on inputExpression and inputPath.
// If inputExpression is not empty, returns strings.NewReader(inputExpression).
// If inputPath is not empty, returns os.Open(inputPath).
// If both inputExpression and inputPath are empty, returns os.Stdin.
func reader(inputExpression *string, inputPath *string) (io.Reader, error) {
	if inputExpression != nil && *inputExpression != "" {
		return strings.NewReader(*inputExpression), nil
	}
	if inputPath != nil && *inputPath != "" {
		reader, err := os.Open(*inputPath)
		if err != nil {
			return nil, err
		}
		return reader, nil
	}
	return os.Stdin, nil
}

// Returns io.Writer depending on outputPath. If outputPath is empty, returns os.Stdout.
func writer(outputPath *string) (io.Writer, error) {
	if outputPath != nil && *outputPath != "" {
		writer, err := os.Create(*outputPath)
		if err != nil {
			return nil, err
		}
		return writer, nil
	}
	return os.Stdout, nil
}
