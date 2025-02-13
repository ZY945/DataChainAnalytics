package api

import "github.com/gin-gonic/gin"

// RegisterRoutes 注册所有路由
func (h *Handlers) RegisterRoutes(router *gin.Engine) {
	// 健康检查接口
	router.GET("/health", h.HealthCheck)

	// API v1 组
	v1 := router.Group("/api/v1")
	{
		// 数据采集相关接口
		collector := v1.Group("/collector")
		{
			collector.POST("/blocks", h.CollectBlockData)
			collector.GET("/blocks/:number", h.GetBlockByNumber)
			collector.POST("/transactions", h.CollectTransactionData)
			collector.GET("/transactions/:hash", h.GetTransactionByHash)
			collector.GET("/status", h.GetStatus)
			collector.GET("/config", h.GetConfig)
		}
		// 数据采集相关接口
		alert := v1.Group("/alert")
		{
			alert.POST("/feishu/send", h.SendFeishuMessage)
		}
	}
}
