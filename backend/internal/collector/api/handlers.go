package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"backend/internal/collector/config"
	"backend/internal/collector/model"
	"backend/internal/notify"
	"backend/pkg/persistence/mysql"
	alltick "backend/pkg/utils/alltick"

	"github.com/gin-gonic/gin"
)

// GoldPriceResponse 黄金价格响应结构
type GoldPriceResponse struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		Code      string `json:"code"`
		KlineType int    `json:"kline_type"`
		KlineList []struct {
			Timestamp  string `json:"timestamp"`
			OpenPrice  string `json:"open_price"`
			ClosePrice string `json:"close_price"`
			HighPrice  string `json:"high_price"`
			LowPrice   string `json:"low_price"`
			Volume     string `json:"volume"`
			Turnover   string `json:"turnover"`
		} `json:"kline_list"`
	} `json:"data"`
}

type Handlers struct {
	db       *mysql.Client
	config   interface{}
	notifier notify.NotifyService
}

func NewHandlers(db *mysql.Client, config interface{}, notifier notify.NotifyService) *Handlers {
	return &Handlers{
		db:       db,
		config:   config,
		notifier: notifier,
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

// SendFeishuMessage 发送飞书消息
func (h *Handlers) SendFeishuMessage(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Type    string `json:"type" binding:"required,oneof=info warning error"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 转换消息类型
	var msgType notify.NotifyType
	switch req.Type {
	case "info":
		msgType = notify.TypeInfo
	case "warning":
		msgType = notify.TypeWarning
	case "error":
		msgType = notify.TypeError
	}

	// 发送消息
	if err := h.notifier.Send(msgType, req.Title, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message sent successfully",
	})
}

// SendFeishuMessage 发送飞书消息
func (h *Handlers) SendFeishuMessageGold(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
		Type    string `json:"type" binding:"required,oneof=info warning error"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 转换消息类型
	var msgType notify.NotifyType
	switch req.Type {
	case "info":
		msgType = notify.TypeInfo
	case "warning":
		msgType = notify.TypeWarning
	case "error":
		msgType = notify.TypeError
	}

	// 发送消息
	if err := h.notifier.Send(msgType, req.Title, req.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message sent successfully",
	})
}

// GetGoldPrice 获取黄金价格并发送飞书通知
func (h *Handlers) GetGoldPriceToFeishuText(c *gin.Context) {
	cfg, ok := h.config.(*config.Config)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid config type"})
		return
	}

	// 获取黄金价格
	goldPrice, err := alltick.HttpGoldPrice(cfg.Gold.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 格式化消息内容
	content, err := alltick.FormatGoldPrice(goldPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("黄金价格：", content)

	// 发送飞书卡片消息
	feishuSvc, ok := h.notifier.(*notify.FeishuNotifyService)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid notifier type"})
		return
	}

	if err := feishuSvc.Send(notify.TypeInfo, "黄金价格更新", content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Gold price fetched and notification sent",
		"data":    goldPrice,
	})
}

// GetGoldPriceToFeishuCard 发送卡片消息
func (h *Handlers) GetGoldPriceToFeishuCard(c *gin.Context) {
	cfg, ok := h.config.(*config.Config)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid config type"})
		return
	}

	// 获取黄金价格
	goldPriceJson, err := alltick.HttpGoldPrice(cfg.Gold.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 解析出goldPriceJson的json需要的参数
	var goldResp GoldPriceResponse
	if err := json.Unmarshal([]byte(goldPriceJson), &goldResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 转换时间戳为标准格式
	timestamp, err := strconv.ParseInt(goldResp.Data.KlineList[0].Timestamp, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	timeStr := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")

	// 构造模板变量
	variables := map[string]interface{}{
		"opening_price":      goldResp.Data.KlineList[0].OpenPrice,
		"closing_price":      goldResp.Data.KlineList[0].ClosePrice,
		"maximum_price":      goldResp.Data.KlineList[0].HighPrice,
		"bottom_price":       goldResp.Data.KlineList[0].LowPrice,
		"time":               timeStr,
		"turnover":           goldResp.Data.KlineList[0].Volume,
		"transaction_volume": goldResp.Data.KlineList[0].Turnover,
	}

	// 发送飞书卡片消息
	templateID := "templateID"
	templateVersion := "1.0.0"
	if err := h.notifier.SendCard("黄金价格更新", variables, templateID, templateVersion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Gold price fetched and notification sent",
		"data":    goldPriceJson,
	})
}
