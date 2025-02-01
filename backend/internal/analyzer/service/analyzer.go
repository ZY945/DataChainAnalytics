package service

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/projectname/internal/analyzer/api"
	"github.com/yourusername/projectname/internal/analyzer/config"
	"github.com/yourusername/projectname/pkg/persistence/mysql"
)

type AnalyzerService struct {
	config   *config.Config
	db       *mysql.Client
	router   *gin.Engine
	handlers *api.Handlers
}

func NewAnalyzerService(cfg *config.Config, db *mysql.Client) *AnalyzerService {
	router := gin.Default()
	handlers := api.NewHandlers(db, cfg)

	service := &AnalyzerService{
		config:   cfg,
		db:       db,
		router:   router,
		handlers: handlers,
	}

	// 注册路由
	handlers.RegisterRoutes(router)
	return service
}

func (s *AnalyzerService) Start() error {
	addr := fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port)
	log.Printf("Analyzer service starting on %s...", addr)
	return s.router.Run(addr)
}
