package air

import (
	"flag"
	"fmt"
	"github.com/cozyo/gs/cmd/air/runner"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
)

var cmdArgs map[string]runner.TomlInfo

func helpMessage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n\n", os.Args[0])
	fmt.Printf("If no command is provided %s will start the runner with the provided flags\n\n", os.Args[0])
	fmt.Println("Commands:")
	fmt.Print("  init	creates a .air.toml file with default settings to the current directory\n\n")

	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func ParseFlag(args []string) Options {
	flag.Usage = helpMessage

	flag.StringVar(&opts.ConfigPath, "c", "", "config path")
	flag.BoolVar(&opts.DebugMode, "d", false, "debug mode")
	flag.BoolVar(&opts.ShowSplash, "v", false, "show version")

	cmd := flag.CommandLine
	cmdArgs = runner.ParseConfigFlag(cmd)

	if err := flag.CommandLine.Parse(args); err != nil {
		log.Fatal(err)
	}

	return opts
}

type VersionInfo struct {
	airVersion string
	goVersion  string
}

func GetVersionInfo() VersionInfo { //revive:disable:unexported-return
	if len(airVersion) != 0 && len(goVersion) != 0 {
		return VersionInfo{
			airVersion: airVersion,
			goVersion:  goVersion,
		}
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		return VersionInfo{
			airVersion: info.Main.Version,
			goVersion:  runtime.Version(),
		}
	}
	return VersionInfo{
		airVersion: "(unknown)",
		goVersion:  runtime.Version(),
	}
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

type Options struct {
	ConfigPath string
	DebugMode  bool
	ShowSplash bool
}

var opts Options

func Run(opts Options) error {
	if opts.ShowSplash {
		printSplash()
		return nil
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	cfg, err := runner.InitConfig(opts.ConfigPath, cmdArgs)
	if err != nil {
		return err
	}

	if !cfg.Log.Silent {
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
