package model

import "time"

// BlockData 区块数据模型
type BlockData struct {
	ID                uint64    `json:"id" gorm:"primarykey"`
	BlockNumber       uint64    `json:"block_number" gorm:"uniqueIndex"`
	BlockHash         string    `json:"block_hash" gorm:"uniqueIndex;size:66"`
	ParentHash        string    `json:"parent_hash" gorm:"size:66"`
	Timestamp         uint64    `json:"timestamp"`
	TransactionsCount uint      `json:"transactions_count"`
	GasUsed           uint64    `json:"gas_used"`
	GasLimit          uint64    `json:"gas_limit"`
	RawData           string    `json:"raw_data" gorm:"type:json"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// TransactionData 交易数据模型
type TransactionData struct {
	ID          uint64    `json:"id" gorm:"primarykey"`
	TxHash      string    `json:"tx_hash" gorm:"uniqueIndex;size:66"`
	BlockNumber uint64    `json:"block_number" gorm:"index"`
	FromAddress string    `json:"from_address" gorm:"index;size:42"`
	ToAddress   string    `json:"to_address" gorm:"index;size:42"`
	Value       string    `json:"value" gorm:"size:78"`
	GasPrice    uint64    `json:"gas_price"`
	GasUsed     uint64    `json:"gas_used"`
	Status      uint8     `json:"status"`
	RawData     string    `json:"raw_data" gorm:"type:json"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
