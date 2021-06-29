package store

import "github.com/hadron-n-veledara-making-games/impact-me/pkg/configreader"

type StoreConfig struct {
	DatabaseURL string `toml:"databse_url"`
}

func ReadConfig(path string) *StoreConfig {
	c := &StoreConfig{}
	configreader.Read(path, c)
	return c
}
