package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type Client struct {
	db *gorm.DB
}

func NewClient(cfg *Config) (*Client, error) {
	if cfg.Host == "" {
		return nil, fmt.Errorf("database host cannot be empty")
	}
	if cfg.Port == 0 {
		cfg.Port = 3306
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Client{db: db}, nil
}

func (c *Client) DB() *gorm.DB {
	return c.db
}
