package main

import (
	"flag"
	"log"

	"backend/internal/collector/config"
	"backend/internal/collector/service"
	"backend/pkg/persistence/mysql"
)

func main() {
	// 命令行参数
	configPath := flag.String("config", "configs/collector.yaml", "path to config file")
	flag.Parse()

	// 加载配置
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库连接
	db, err := mysql.NewClient(&cfg.Database.CollectorDB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化采集器服务
	collectorService := service.NewCollectorService(cfg, db)

	// 启动服务
	log.Printf("Starting collector service with config: %s", *configPath)
	if err := collectorService.Start(); err != nil {
		log.Fatalf("Failed to start collector service: %v", err)
	}
}
