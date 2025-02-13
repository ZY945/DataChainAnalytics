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
		AnalyzerDB mysql.Config `yaml:"analyzer_db"`
	} `yaml:"database"`

	Analyzer struct {
		Interval  int `yaml:"interval"`
		BatchSize int `yaml:"batch_size"`
	} `yaml:"analyzer"`
}

func LoadConfig(configPath string) (*Config, error) {
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
