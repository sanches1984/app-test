package config

import (
	"github.com/rs/zerolog"
)

type BotConfig struct {
	Name  string `yaml:"name"`
	Title string `yaml:"title"`
	Token string `yaml:"token"`
}

type Config struct {
	Bot           BotConfig     `yaml:"bot"`
	DBCredentials string        `yaml:"dbCredentials"`
	LogLevel      zerolog.Level `yaml:"logLevel"`
}

func (c *Config) Validate() error {
	// todo
	return nil
}

func NewDefaultConfig() *Config {
	return &Config{
		Bot: BotConfig{
			Name:  "referral-bot",
			Title: "Referral Bot",
		},
		DBCredentials: "db_credentials.json",
		LogLevel:      zerolog.DebugLevel,
	}
}
