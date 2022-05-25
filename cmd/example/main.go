package main

import (
	"flag"
	"fmt"

	handler "github.com/Beaxhem/architecture-lab-2/handler"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputPath       = flag.String("f", "", "Path to file with expression to compute")
	outputPath      = flag.String("o", "", "Path to file to write the result to")
)

func main() {
	flag.Parse()
	handler, err := handler.NewComputeHandler(inputExpression, inputPath, outputPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = handler.Compute()
	if err != nil {
		fmt.Println(err)
	}
}
