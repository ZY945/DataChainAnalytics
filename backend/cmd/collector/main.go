package main

import (
	"log"

	"github.com/yourusername/projectname/internal/collector/config"
	"github.com/yourusername/projectname/internal/collector/service"
	"github.com/yourusername/projectname/pkg/persistence/mysql"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化数据库连接
	db, err := mysql.NewClient(&cfg.Database.CollectorDB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化采集器服务
	collectorService := service.NewCollectorService(db)

	// 启动服务
	if err := collectorService.Start(); err != nil {
		log.Fatalf("Failed to start collector service: %v", err)
	}

	log.Println("Collector service started successfully")
}
