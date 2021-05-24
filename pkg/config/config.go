package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken string
}

func Init() (c *Config, err error) {
	var cfg Config
	if err = parseEnv(c); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func parseEnv(c *Config) error {
	if err := viper.BindEnv("token"); err != nil {
		return err
	}
	c.TelegramToken = viper.GetString("token")
	return nil
}
