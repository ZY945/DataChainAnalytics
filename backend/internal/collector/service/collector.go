package service

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/projectname/internal/collector/api"
	"github.com/yourusername/projectname/internal/collector/config"
	"github.com/yourusername/projectname/pkg/persistence/mysql"
)

type CollectorService struct {
	config   *config.Config
	db       *mysql.Client
	router   *gin.Engine
	handlers *api.Handlers
}

func NewCollectorService(cfg *config.Config, db *mysql.Client) *CollectorService {
	router := gin.Default()
	handlers := api.NewHandlers(db, cfg)

	service := &CollectorService{
		config:   cfg,
		db:       db,
		router:   router,
		handlers: handlers,
	}

	// 注册路由
	handlers.RegisterRoutes(router)
	return service
}

func (s *CollectorService) Start() error {
	addr := fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port)
	log.Printf("Collector service starting on %s...", addr)
	return s.router.Run(addr)
}
