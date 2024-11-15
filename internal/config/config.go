package config

type Config struct {
	IP   string
	Port string `json:"port"`
}

var DefaultConfig = Config{
	IP:   "127.0.01",
	Port: "8080",
}

func LoadConfig() (cfg *Config, err error) {

	err = nil
	cfg = &DefaultConfig

	return cfg, err
}
