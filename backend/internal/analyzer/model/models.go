package model

import "time"

// AnalysisTask 分析任务模型
type AnalysisTask struct {
	ID         uint64    `json:"id" gorm:"primarykey"`
	TaskType   string    `json:"task_type" gorm:"size:50"`
	Status     string    `json:"status" gorm:"size:20"`
	StartBlock uint64    `json:"start_block"`
	EndBlock   uint64    `json:"end_block"`
	Parameters string    `json:"parameters" gorm:"type:json"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// AnalysisResult 分析结果模型
type AnalysisResult struct {
	ID         uint64    `json:"id" gorm:"primarykey"`
	TaskID     uint64    `json:"task_id" gorm:"index"`
	ResultType string    `json:"result_type" gorm:"size:50"`
	Data       string    `json:"data" gorm:"type:json"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
