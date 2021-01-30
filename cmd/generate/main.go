package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sh-miyoshi/dxlib/cmd/generate/output"
	"github.com/sh-miyoshi/dxlib/cmd/generate/parse"
)

func main() {
	var input, out string
	flag.StringVar(&input, "input", "", "input file name")
	flag.StringVar(&input, "i", "", "input file name")
	flag.StringVar(&out, "output", "", "output file name")
	flag.StringVar(&out, "o", "", "output file name")

	flag.Parse()

	if input == "" || out == "" {
		fmt.Printf("Please set input and output\n")
		os.Exit(1)
	}

	parser, err := parse.New(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer parser.End()

	data, err := parser.Parse()
	if err != nil {
		fmt.Printf("Failed to parse data: %v\n", err)
		os.Exit(1)
	}

	output.Output(out, *data)
}
