package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fantolla123/lzt-api-go/internal/codegen"
)

func main() {
	schema := flag.String("schema", "", "")
	output := flag.String("output", "", "")
	pkg := flag.String("package", "", "")
	flag.Parse()

	if *schema == "" || *output == "" || *pkg == "" {
		fmt.Fprintf(os.Stderr, "usage: codegen --schema <path> --output <dir> --package <name>\n")
		os.Exit(1)
	}

	result, err := codegen.ParseSchema(*schema)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parse error: %v\n", err)
		os.Exit(1)
	}

	if err := codegen.Generate(result, *output, *pkg); err != nil {
		fmt.Fprintf(os.Stderr, "generate error: %v\n", err)
		os.Exit(1)
	}
}
