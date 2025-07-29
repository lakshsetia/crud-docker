package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address" env-required:"true"`
}

type Postgresql struct {
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DBName   string `yaml:"dbname" env-required:"true"`
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
}

type Database struct {
	Postgresql Postgresql `yaml:"postgresql" env-required:"true"`
}

type Config struct {
	Env        string     `yaml:"env" ENV:"env" env-required:"true"`
	HTTPServer HTTPServer `yaml:"http_server" env-required:"true"`
	Database   Database   `yaml:"database" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		configPath = *flags
		if configPath == "" {
			return nil, fmt.Errorf("config path not specified")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file not found at %v", configPath)
	}
	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return nil, fmt.Errorf("failed to read config file at %v: %w", configPath, err)
	}
	return &config, nil
}