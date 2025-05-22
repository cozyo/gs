package main

import (
	"fmt"
	"github.com/cozyo/gs/pkg/gen"
	"os"
)

func main() {
	var arg []string
	opts := gen.ParseFlag(arg)
	tableName := "page"
	if err := gen.Run(opts, tableName); err != nil {
		fmt.Fprintf(os.Stderr, "Error running air: %v\n", err)
		os.Exit(1)
	}
}
