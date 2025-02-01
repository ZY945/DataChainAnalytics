package types

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// CollectorConfig 采集器配置
type CollectorConfig struct {
	ID            uint   `gorm:"primarykey"`
	Name          string `gorm:"unique;size:100"`
	ChainType     string `gorm:"size:50"`
	RPCURL        string `gorm:"column:rpc_url;type:text"`
	APIKey        string `gorm:"column:api_key;size:255"`
	RetryTimes    int    `gorm:"default:3"`
	RetryInterval int    `gorm:"default:1000"`
	Status        int    `gorm:"default:1"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// CollectionField JSON 字段结构
type CollectionField struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Index int    `json:"index"`
}

// CollectionFilter 过滤器配置
type CollectionFilter struct {
	Address []string `json:"address,omitempty"`
	Topics  []string `json:"topics,omitempty"`
}

// CollectionFields 采集字段配置
type CollectionFields struct {
	EventSignature string            `json:"event_signature"`
	Fields         []CollectionField `json:"fields"`
	Filters        CollectionFilter  `json:"filters"`
}

// Value 实现 driver.Valuer 接口
func (c CollectionFields) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan 实现 sql.Scanner 接口
func (c *CollectionFields) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, &c)
}

// CollectorTask 采集任务
type CollectorTask struct {
	ID               uint   `gorm:"primarykey"`
	ConfigID         uint   `gorm:"not null"`
	Name             string `gorm:"unique;size:100"`
	Description      string `gorm:"type:text"`
	ContractAddress  string `gorm:"size:42"`
	StartBlock       uint64
	EndBlock         uint64
	CollectionFields CollectionFields `gorm:"type:json"`
	TaskInterval     int              `gorm:"default:5000"`
	BatchSize        int              `gorm:"default:100"`
	Status           int              `gorm:"default:1"`
	LastBlock        uint64
	ErrorCount       int    `gorm:"default:0"`
	LastError        string `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Config           CollectorConfig `gorm:"foreignKey:ConfigID"`
}

// CollectedData 采集的数据
type CollectedData struct {
	ID              uint   `gorm:"primarykey"`
	TaskID          uint   `gorm:"not null"`
	BlockNumber     uint64 `gorm:"not null"`
	BlockHash       string `gorm:"size:66"`
	TransactionHash string `gorm:"size:66"`
	CollectedData   string `gorm:"type:json"`
	CreatedAt       time.Time
	Task            CollectorTask `gorm:"foreignKey:TaskID"`
}
