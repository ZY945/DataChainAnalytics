package api

import (
	"net/http"
	"time"

	"backend/internal/analyzer/model"
	"backend/pkg/persistence/mysql"
	"backend/pkg/persistence/mysql/models"

	"github.com/gin-gonic/gin"
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

// 创建 API 配置

// 创建 API 配置
func (h *Handlers) CreateConfig(c *gin.Context) {
	var config models.APIConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.DB().Create(&config).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

// 更新 API 配置
func (h *Handlers) UpdateConfig(c *gin.Context) {
	var config models.APIConfig
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.DB().Save(&config).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, config)
}

// 获取所有 API 配置
func (h *Handlers) ListConfig(c *gin.Context) {
	var configs []models.APIConfig
	if err := h.db.DB().Where("is_deleted = ?", 0).Find(&configs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, configs)
}

// 获取单个 API 配置
func (h *Handlers) GetOneConfig(c *gin.Context) {
	id := c.Param("id")
	var config models.APIConfig
	if err := h.db.DB().First(&config, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "API config not found"})
		return
	}

	c.JSON(http.StatusOK, config)
}

// 删除 API 配置（软删除）
func (h *Handlers) DeleteConfig(c *gin.Context) {
	id := c.Param("id")
	if err := h.db.DB().Model(&models.APIConfig{}).Where("id = ?", id).Update("is_deleted", 1).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "API config deleted successfully"})
}

// 更新 API 状态
func (h *Handlers) UpdateStatusConfig(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		Status int8 `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.DB().Model(&models.APIConfig{}).Where("id = ?", id).Update("status", req.Status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}
