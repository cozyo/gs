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

	// gs gen model / ctrl / service
	var genCmd = &cobra.Command{
		Use: "gen",
		Short: "The \"gen\" command is designed for multiple generating purposes.\n" +
			"It's currently supporting generating go files for ORM models, protobuf and protobuf entity files.\n" +
			"Please use \"gf gen model -h\" for specified type help.",
	}

	var genModelCmd = &cobra.Command{
		Use:     "model",
		Short:   "automatically generate go files for (model/entity)",
		Args:    cobra.MaximumNArgs(1),
		Example: "gs gen model page",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generating model...")
			dispatcher.RunGen(args)
		},
	}

	var genCtrlCmd = &cobra.Command{
		Use:     "ctrl",
		Short:   "parse api definitions to generate controller/sdk go files",
		Example: "gs gen ctrl page",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generating ctrl...")
			// 实现 ctrl 生成逻辑
		},
	}

	var genServiceCmd = &cobra.Command{
		Use:     "service",
		Short:   "parse struct and associated functions from packages to generate service go file",
		Example: "gs gen service page",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Generating service...")
			// 实现 service 生成逻辑
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
	genCmd.AddCommand(genModelCmd, genCtrlCmd, genServiceCmd)
	dbCmd.AddCommand(migrateCmd)

	rootCmd.AddCommand(runCmd, genCmd, dbCmd)

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
