package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Database struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Dbname   string `toml:"dbname"`
		Sslmode  string `toml:"sslmode"`
		RedisAddr  string `toml:"redis_addr"`
	} `toml:"database"`

	Email struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		Password string `toml:"password"`
		From string `toml:"from"`
	} `toml:"email"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
		return nil, err
	}
	return &config, nil
}