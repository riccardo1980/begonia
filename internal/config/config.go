package config

import (
	"fmt"
	"os"

	// this will automatically load your .env file:
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var defaultEnvFile = "./.env"

type Config struct {
	IP   string
	Port string `json:"port"`
}

var defaultConfig = Config{
	IP:   "127.0.0.1",
	Port: "8080",
}

func (c *Config) toDotenv(filename string) (int, error) {
	f, err := os.Create(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	msg := ""
	msg += fmt.Sprintf("PORT=%s\n", c.Port)
	msg += fmt.Sprintf("IP=%s\n", c.IP)

	return f.WriteString(msg)
}

func (c *Config) fromEnvOrDefault() {
	c.IP = os.Getenv("IP")
	if c.IP == "" {
		c.IP = defaultConfig.IP
	}

	c.Port = os.Getenv("PORT")
	if c.Port == "" {
		c.Port = defaultConfig.Port
	}
}

func loadOrDefault(envFile string) (*Config, error) {
	// load env file

	godotenv.Load(envFile)
	cfg := &Config{}
	cfg.fromEnvOrDefault()

	return cfg, nil
}

func LoadConfig() (*Config, error) {

	return loadOrDefault(defaultEnvFile)
}
