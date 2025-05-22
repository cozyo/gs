package air

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"

	"github.com/cozyo/gs/cmd/air/runner"
)

type Options struct {
	ConfigPath string
	DebugMode  bool
	ShowSplash bool
}

func Run(opts Options) error {
	if opts.ShowSplash {
		printSplash()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	cmdArgs := runner.ParseConfigFlag(nil)
	cfg, err := runner.InitConfig(opts.ConfigPath, cmdArgs)
	if err != nil {
		return err
	}
	if !cfg.Log.Silent && opts.ShowSplash {
		printSplash()
	}
	if opts.DebugMode && !cfg.Log.Silent {
		fmt.Println("[debug] mode")
	}

	r, err := runner.NewEngineWithConfig(cfg, opts.DebugMode)
	if err != nil {
		return err
	}

	go func() {
		<-sigs
		r.Stop()
	}()

	defer func() {
		if e := recover(); e != nil {
			log.Fatalf("PANIC: %+v", e)
		}
	}()

	r.Run()
	return nil
}

func printSplash() {
	version := GetVersionInfo()
	fmt.Printf(`
   ____   ____  
  / ___| / ___| 
 | |  _  \___ \ 
 | |_| |  ___) |
  \____| |____/   %s, built with Go %s

`, version.airVersion, version.goVersion)
}

type versionInfo struct {
	airVersion string
	goVersion  string
}

func GetVersionInfo() versionInfo {
	if info, ok := debug.ReadBuildInfo(); ok {
		return versionInfo{
			airVersion: info.Main.Version,
			goVersion:  runtime.Version(),
		}
	}
	return versionInfo{
		airVersion: "(unknown)",
		goVersion:  runtime.Version(),
	}
}
