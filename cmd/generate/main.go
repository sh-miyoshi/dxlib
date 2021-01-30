package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sh-miyoshi/dxlib/cmd/generate/output"
	"github.com/sh-miyoshi/dxlib/cmd/generate/parse"
)

// func parseComment(line string) error {
// 	if !strings.HasPrefix(line, "//comment") {
// 		return nil
// 	}

// 	// valid format
// 	//comment; <FuncName>; <comments>
// 	data := strings.Split(line, ";")
// 	if len(data) != 3 {
// 		return fmt.Errorf("invalid comment format")
// 	}

// 	comments = append(comments, comment{
// 		funcName: strings.TrimSpace(data[1]),
// 		comments: strings.Split(strings.TrimSpace(data[2]), "\\n"),
// 	})

// 	return nil
// }

// func getComment(funcName string) string {
// 	for _, c := range comments {
// 		if c.funcName == funcName {
// 			res := "// " + funcName + " "
// 			for _, msg := range c.comments {
// 				res += msg + "\n// "
// 			}
// 			return strings.TrimSuffix(res, "// ")
// 		}
// 	}
// 	return ""
// }

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

	// 	// Parse function
	// 	d, err := parse(line)
	// 	if err != nil {
	// 		fmt.Printf("Failed to parse line %s: %v\n", line, err)
	// 		os.Exit(1)
	// 	}
	// 	if d != nil {
	// 		data = append(data, *d)
	// 		continue
	// 	}

	// 	// Parse extra function
	// 	if d := parseExtraFunc(line); d != nil {
	// 		extra = append(extra, *d)
	// 		continue
	// 	}

	// 	// Parse command
	// 	if err := parseComment(line); err != nil {
	// 		fmt.Printf("Failed to parse line %s: %v\n", line, err)
	// 		os.Exit(1)
	// 	}
	// }
}
