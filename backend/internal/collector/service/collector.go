package service

import (
	"fmt"
	"log"

	"backend/internal/collector/api"
	"backend/internal/collector/config"
	"backend/internal/notify"
	"backend/pkg/persistence/mysql"

	"github.com/gin-gonic/gin"
)

type CollectorService struct {
	config   *config.Config
	db       *mysql.Client
	router   *gin.Engine
	handlers *api.Handlers
	notifier notify.NotifyService
}

func NewCollectorService(cfg *config.Config, db *mysql.Client) *CollectorService {
	router := gin.Default()

	// 创建飞书通知服务
	feishuNotifier := notify.NewFeishuNotifyService(
		cfg.Alert.FeishuConfig.Webhook,
		cfg.Alert.FeishuConfig.AppSecret,
	)

	// 创建默认通知服务，包含飞书通知
	defaultNotifier := notify.NewDefaultNotifyService(feishuNotifier)

	// 创建处理器
	handlers := api.NewHandlers(db, cfg, defaultNotifier)

	service := &CollectorService{
		config:   cfg,
		db:       db,
		router:   router,
		handlers: handlers,
		notifier: defaultNotifier,
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

func (s *CollectorService) handleError(err error) {
	s.notifier.Send(notify.TypeError, "Collector Error", err.Error())
}

func (s *CollectorService) reportStatus(status string) {
	s.notifier.Send(notify.TypeInfo, "Collector Status", status)
}
