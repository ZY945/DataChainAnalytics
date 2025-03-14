package api

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册所有路由
func (h *Handlers) RegisterRoutes(router *gin.Engine) {
	// 健康检查接口
	router.GET("/health", h.HealthCheck)

	// API v1 组
	v1 := router.Group("/api/v1")
	{
		analyzer := v1.Group("/analyzer")
		{
			// 分析任务相关接口
			analyzer.POST("/tasks", h.CreateAnalysisTask)
			analyzer.GET("/tasks/:id", h.GetAnalysisTask)

			// 服务状态接口
			analyzer.GET("/status", h.GetStatus)
			analyzer.GET("/config", h.GetConfig)
		}

		// API 配置管理接口
		config := v1.Group("/config")
		{
			config.POST("/", h.CreateConfig)
			config.PUT("/:id", h.UpdateConfig)
			config.GET("/", h.ListConfig)
			config.GET("/:id", h.GetOneConfig)
			config.DELETE("/:id", h.DeleteConfig)
			config.PATCH("/:id/status", h.UpdateStatusConfig)
		}
	}
}
