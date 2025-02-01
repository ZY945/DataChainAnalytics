package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/projectname/internal/analyzer/model"
	"github.com/yourusername/projectname/pkg/persistence/mysql"
)

type Handlers struct {
	db     *mysql.Client
	config interface{}
}

func NewHandlers(db *mysql.Client, config interface{}) *Handlers {
	return &Handlers{
		db:     db,
		config: config,
	}
}

// HealthCheck 健康检查
func (h *Handlers) HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"service": "analyzer",
	})
}

// GetConfig 获取配置信息
func (h *Handlers) GetConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"config": h.config,
	})
}

// CreateAnalysisTask 创建分析任务
func (h *Handlers) CreateAnalysisTask(c *gin.Context) {
	var task model.AnalysisTask
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task.Status = "pending"
	if err := h.db.DB().Create(&task).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Analysis task created successfully",
		"data":    task,
	})
}

// GetAnalysisTask 获取分析任务
func (h *Handlers) GetAnalysisTask(c *gin.Context) {
	id := c.Param("id")
	var task model.AnalysisTask

	if err := h.db.DB().First(&task, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(200, task)
}

// GetStatus 获取服务状态
func (h *Handlers) GetStatus(c *gin.Context) {
	var taskCount int64
	h.db.DB().Model(&model.AnalysisTask{}).Count(&taskCount)

	c.JSON(200, gin.H{
		"status":        "running",
		"timestamp":     time.Now().Unix(),
		"pending_tasks": taskCount,
	})
}
