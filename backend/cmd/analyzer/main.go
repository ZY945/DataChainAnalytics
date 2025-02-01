package main

import (
	"flag"
	"log"

	"github.com/yourusername/projectname/internal/analyzer/config"
	"github.com/yourusername/projectname/internal/analyzer/service"
	"github.com/yourusername/projectname/pkg/persistence/mysql"
)

func main() {
	configPath := flag.String("config", "configs/analyzer.yaml", "path to config file")
	flag.Parse()

	// 加载配置
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库连接
	db, err := mysql.NewClient(&cfg.Database.AnalyzerDB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化分析器服务
	analyzerService := service.NewAnalyzerService(cfg, db)

	// 启动服务
	log.Printf("Starting analyzer service with config: %s", *configPath)
	if err := analyzerService.Start(); err != nil {
		log.Fatalf("Failed to start analyzer service: %v", err)
	}
}
