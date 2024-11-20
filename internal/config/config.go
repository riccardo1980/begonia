package config

import (
	"os"
	// this will automatically load your .env file:
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	IP   string
	Port string `json:"port"`
}

var DefaultConfig = Config{
	IP:   "127.0.0.1",
	Port: "8080",
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	cfg.IP = os.Getenv("IP")
	if cfg.IP == "" {
		cfg.IP = DefaultConfig.IP
	}

	cfg.Port = os.Getenv("PORT")
	if cfg.Port == "" {
		cfg.Port = DefaultConfig.Port
	}

	return cfg, nil
}
