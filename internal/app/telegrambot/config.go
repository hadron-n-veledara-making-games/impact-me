package telegrambot

import "github.com/hadron-n-veledara-making-games/impact-me/pkg/configreader"

type BotConfig struct {
	Token    string `toml:"token"`
	LogLevel string `toml:"log_level"`
}

func ReadConfig(path string) *BotConfig {
	c := &BotConfig{
		LogLevel: "debug",
	}
	configreader.Read(path, c)
	return c
}
