package dispatcher

import (
	"fmt"
	"github.com/cozyo/gs/cmd/air"
	"os"
)

func RunAir(args []string) {
	opts := air.ParseFlag(args)
	if err := air.Run(opts); err != nil {
		fmt.Fprintf(os.Stderr, "Error running air: %v\n", err)
		os.Exit(1)
	}
}
