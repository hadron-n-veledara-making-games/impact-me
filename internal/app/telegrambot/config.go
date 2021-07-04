package telegrambot

import (
	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/broker"
	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/store"
	"github.com/hadron-n-veledara-making-games/impact-me/pkg/configreader"
)

type BotConfig struct {
	Debug    bool   `toml:"debug"`
	Token    string `toml:"token"`
	LogLevel string `toml:"log_level"`
	Broker   *broker.BrokerConfig
	Store    *store.StoreConfig
}

func ReadConfig(path string) *BotConfig {
	c := &BotConfig{
		LogLevel: "debug",
		Debug:    true,
		Broker:   broker.NewConfig(),
		Store:    store.NewConfig(),
	}
	configreader.Read(path, c)
	c.Broker.BuildAmqpURL()
	c.Store.BuildDBComplexFields()
	return c
}
