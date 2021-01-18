package config

import (
	"encoding/json"
	"os"
)

type Database struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
}

type Config struct {
	AllowedOrigins  string   `json:"allowed-origins"`
	ApplicationPort string   `json:"application-port"`
	Database        Database `json:"database"`
}

func ParseConfig() (c *Config, err error) {
	f, err := os.Open("/root/automatisation/InstallationService/config/config.json")
	if err != nil {
		return
	}
	c = new(Config)
	err = json.NewDecoder(f).Decode(c)
	return
}
