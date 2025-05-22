package gen

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Options struct {
	OutPath        string
	ModelPkgPath   string
	FieldNullable  bool
	FieldCoverable bool
	FieldSignable  bool
	DBKey          string
}

type DefaultConfig struct {
	DSN string `yaml:"dsn"`
}

type DatabaseConfig struct {
	MySQL map[string]DefaultConfig `yaml:"mysql"` // 支持多个数据库配置项
}

type MySQLConfig struct {
	Default DefaultConfig `yaml:"default"`
}

type LogConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type AppConfig struct {
	Database DatabaseConfig `yaml:"database"`
	Log      LogConfig      `yaml:"log"`
}

func LoadDSNFromYAML(configPath, dbKey string) (string, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", fmt.Errorf("failed to read config: %w", err)
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return "", fmt.Errorf("failed to unmarshal config: %w", err)
	}

	dbConfig, ok := cfg.Database.MySQL[dbKey]
	if !ok {
		return "", fmt.Errorf("database config not found for key: %s", dbKey)
	}

	return dbConfig.DSN, nil
}

func ParseFlag(args []string) Options {
	var opts Options

	fs := flag.NewFlagSet("gorm-gen", flag.ExitOnError)

	fs.StringVar(&opts.OutPath, "out", "./internal", "输出目录路径")
	fs.StringVar(&opts.ModelPkgPath, "model-pkg", "model/entity", "Model 包路径")
	fs.BoolVar(&opts.FieldNullable, "nullable", true, "是否生成可空字段指针")
	// 如果列在数据库中有默认值，则生成字段类型的指针, 避免零值问题 https://gorm.io/docs/create.html#Default-Values
	fs.BoolVar(&opts.FieldCoverable, "coverable", true, "是否生成默认值字段为指针")
	fs.BoolVar(&opts.FieldSignable, "signable", true, "是否根据类型使用有符号类型")
	fs.StringVar(&opts.DBKey, "db", "default", "读取 config.yaml 中 database.mysql 下的配置名")

	_ = fs.Parse(args)

	return opts
}

func Run(opts Options, tableName string) error {
	dsn, err := LoadDSNFromYAML("config/config.yaml", opts.DBKey)
	if err != nil {
		log.Fatalf("加载 DSN 出错: %v", err)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 初始化gen
	g := NewGenerator(Config{
		OutPath:        opts.OutPath,
		ModelPkgPath:   opts.ModelPkgPath,
		FieldNullable:  opts.FieldNullable,
		FieldCoverable: opts.FieldCoverable,
		FieldSignable:  opts.FieldSignable,
	})

	// 连接数据库
	g.UseDB(db)

	// 自定义模板
	g.ApplyBasic(
		// 这里用你的表名
		g.GenerateModel(tableName),
	)

	g.ExecuteModel()
	return nil
}
