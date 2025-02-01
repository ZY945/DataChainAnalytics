package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/projectname/internal/collector/model"
	"github.com/yourusername/projectname/pkg/persistence/mysql"
)

type Handlers struct {
	db     *mysql.Client
	config interface{} // 根据需要定义配置类型
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
		"service": "collector",
	})
}

// GetConfig 获取配置信息
func (h *Handlers) GetConfig(c *gin.Context) {
	c.JSON(200, gin.H{
		"config": h.config,
	})
}

// CollectBlockData 采集区块数据
func (h *Handlers) CollectBlockData(c *gin.Context) {
	var block model.BlockData
	if err := c.ShouldBindJSON(&block); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.DB().Create(&block).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Block data collected successfully",
		"data":    block,
	})
}

// GetBlockByNumber 根据区块号获取区块数据
func (h *Handlers) GetBlockByNumber(c *gin.Context) {
	number := c.Param("number")
	var block model.BlockData

	if err := h.db.DB().Where("block_number = ?", number).First(&block).Error; err != nil {
		c.JSON(404, gin.H{"error": "Block not found"})
		return
	}

	c.JSON(200, block)
}

// CollectTransactionData 采集交易数据
func (h *Handlers) CollectTransactionData(c *gin.Context) {
	var tx model.TransactionData
	if err := c.ShouldBindJSON(&tx); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.DB().Create(&tx).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Transaction data collected successfully",
		"data":    tx,
	})
}

// GetTransactionByHash 根据哈希获取交易数据
func (h *Handlers) GetTransactionByHash(c *gin.Context) {
	hash := c.Param("hash")
	var tx model.TransactionData

	if err := h.db.DB().Where("tx_hash = ?", hash).First(&tx).Error; err != nil {
		c.JSON(404, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(200, tx)
}

// GetStatus 获取服务状态
func (h *Handlers) GetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":    "running",
		"timestamp": time.Now().Unix(),
	})
}
