package dispatcher

import (
	"fmt"
	"github.com/cozyo/gs/pkg/air"
	"github.com/cozyo/gs/pkg/gen"
	"os"
)

func RunAir(args []string) {
	opts := air.ParseFlag(args)
	if err := air.Run(opts); err != nil {
		fmt.Fprintf(os.Stderr, "Error running air: %v\n", err)
		os.Exit(1)
	}
}

func RunGen(args []string) {
	tableName := "page" // 默认值
	if len(args) > 0 {
		tableName = args[0]
	}
	opts := gen.ParseFlag(args)
	if err := gen.Run(opts, tableName); err != nil {
		fmt.Fprintf(os.Stderr, "Error running air: %v\n", err)
		os.Exit(1)
	}
}
