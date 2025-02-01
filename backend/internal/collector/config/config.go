package config

import (
	"os"

	"github.com/yourusername/projectname/pkg/persistence/mysql"
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
		MasterDB    mysql.Config `yaml:"master_db"`
		ReadOnlyDB  mysql.Config `yaml:"readonly_db"`
	} `yaml:"database"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
