package main

import (
	"fmt"
	"github.com/cozyo/gs/dispatcher"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "gs",
		Short: "GS CLI Tool - Hot reload and code generation",
	}

	// gs run
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Hot reload and run the project",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running with hot reload...")
			dispatcher.RunAir(args)
		},
	}

	// gs gen model / handler / service
	var genCmd = &cobra.Command{
		Use:   "gen",
		Short: "Generate code (model/entity)",
	}

	var genModelCmd = &cobra.Command{
		Use:   "model",
		Short: "Generate model code",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generating model...")
			dispatcher.RunGen(args)
		},
	}

	var genHandlerCmd = &cobra.Command{
		Use:   "handler",
		Short: "Generate handler code",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generating handler...")
			// 实现 handler 生成逻辑
		},
	}

	var genServiceCmd = &cobra.Command{
		Use:   "service",
		Short: "Generate service code",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generating service...")
			// 实现 service 生成逻辑
		},
	}

	// gs make command
	var makeCmd = &cobra.Command{
		Use:   "make",
		Short: "Make scaffold files",
	}

	var makeCommandCmd = &cobra.Command{
		Use:   "command",
		Short: "Generate a custom command scaffold",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generating custom command scaffold...")
			// 自定义 command 生成逻辑
		},
	}

	// gs db migrate
	var dbCmd = &cobra.Command{
		Use:   "db",
		Short: "Database-related commands",
	}

	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running DB migrations...")
			// 执行数据库迁移逻辑
		},
	}

	// 组装子命令结构
	genCmd.AddCommand(genModelCmd, genHandlerCmd, genServiceCmd)
	makeCmd.AddCommand(makeCommandCmd)
	dbCmd.AddCommand(migrateCmd)

	rootCmd.AddCommand(runCmd, genCmd, makeCmd, dbCmd)

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
