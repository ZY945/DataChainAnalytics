package config

import (
	"fmt"
	"os"
	"path/filepath"

	"backend/pkg/persistence/mysql"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Database struct {
		// MySQL mysql.Config `yaml:"mysql"`
		CollectorDB mysql.Config `yaml:"collector_db"`
	} `yaml:"database"`

	Collector struct {
		Interval  int `yaml:"interval"`
		BatchSize int `yaml:"batch_size"`
	} `yaml:"collector"`

	Alert struct {
		FeishuConfig struct {
			AppID     string `yaml:"app_id"`
			AppSecret string `yaml:"app_secret"`
			Webhook   string `yaml:"webhook"`
		} `yaml:"feishu"`
	} `yaml:"alert"`
}

// LoadConfig 加载配置文件
func LoadConfig(configPath string) (*Config, error) {
	// 如果传入的是相对路径，转换为绝对路径
	if !filepath.IsAbs(configPath) {
		pwd, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("get working directory failed: %v", err)
		}
		configPath = filepath.Join(pwd, configPath)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read config file failed: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshal config failed: %v", err)
	}

	return &config, nil
}
