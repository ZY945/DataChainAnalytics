package models

import "time"

type APIConfig struct {
	ID        uint64    `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name;unique"`
	URL       string    `gorm:"column:url"`
	Token     string    `gorm:"column:token"`
	Secret    string    `gorm:"column:secret"`
	Type      int8      `gorm:"column:type;default:1"`
	Status    int8      `gorm:"column:status;default:1"`
	IsDeleted int8      `gorm:"column:is_deleted;default:0"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (APIConfig) TableName() string {
	return "api_config"
}
