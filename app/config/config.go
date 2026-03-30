package config

import (
	"os"
)

type Config struct {
	DB   DBConfig
	HTTP HTTPConfig
}

type DBConfig struct {
	Driver string // sqlite, postgres, mysql
	DSN    string // Data Source Name
}

type HTTPConfig struct {
	Port string
}

func LoadConfig() *Config {
	// Default configuration
	cfg := &Config{
		DB: DBConfig{
			Driver: getEnv("DB_DRIVER", "sqlite"),
			DSN:    getEnv("DB_DSN", "./database.sqlite"),
		},
		HTTP: HTTPConfig{
			Port: getEnv("HTTP_PORT", "3000"),
		},
	}
	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
