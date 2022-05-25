package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/Beaxhem/architecture-lab-2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputPath       = flag.String("f", "", "Path to file with expression to compute")
	outputPath      = flag.String("o", "", "Path to file to write the result")
)

func main() {
	flag.Parse()
	handler := lab2.ComputeHandler{
		Reader: reader(inputExpression, inputPath),
		Writer: writer(outputPath),
	}

	err := handler.Compute()
	if err != nil {
		fmt.Println(err)
	}
}

// Returns io.Reader depending on inputExpression and inputPath.
// If inputExpression is not empty, returns strings.NewReader(inputExpression).
// If inputPath is not empty, returns os.Open(inputPath).
// If both inputExpression and inputPath are empty, returns os.Stdin.
func reader(inputExpression *string, inputPath *string) io.Reader {
	if inputExpression != nil && *inputExpression != "" {
		return strings.NewReader(*inputExpression)
	}
	if inputPath != nil && *inputPath != "" {
		reader, err := os.Open(*inputPath)
		if err != nil {
			return nil
		}
		return reader
	}
	return os.Stdin
}

// Returns io.Writer depending on outputPath. If outputPath is empty, returns os.Stdout.
func writer(outputPath *string) io.Writer {
	if outputPath != nil && *outputPath != "" {
		writer, err := os.Create(*outputPath)
		if err != nil {
			return nil
		}
		return writer
	}
	return os.Stdout
}
