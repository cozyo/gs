package dispatcher

import (
	"fmt"
	"github.com/cozyo/gs/cmd/air"
)

func runAir(args []string) {
	opts := air.Options{
		ConfigPath: "",    // 可从 args 解析
		DebugMode:  false, // 可从 args 解析
		ShowSplash: true,
	}
	if err := air.Run(opts); err != nil {
		fmt.Println("Air run failed:", err)
	}
}
