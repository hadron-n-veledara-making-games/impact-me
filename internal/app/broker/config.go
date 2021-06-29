package broker

import (
	"fmt"

	"github.com/hadron-n-veledara-making-games/impact-me/pkg/configreader"
)

type BrokerConfig struct {
	Login    string `toml:"login"`
	Pass     string `toml:"pass"`
	Post     string `toml:"host"`
	Port     int    `toml:"port"`
	LogLevel string `toml:"log_level"`
	Debug    bool   `toml:"debug"`
	URL      string
}

func NewConfig() *BrokerConfig {
	return &BrokerConfig{}
}

func ReadConfig(path string) *BrokerConfig {
	c := &BrokerConfig{}
	configreader.Read(path, c)
	return c
}

func (c *BrokerConfig) BuildAmqpURL() {
	c.URL = fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Login, c.Pass, c.Post, c.Port)
}
