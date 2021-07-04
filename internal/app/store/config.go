package store

import (
	"fmt"

	"github.com/hadron-n-veledara-making-games/impact-me/pkg/configreader"
)

type StoreConfig struct {
	DBParams   string
	DBURL      string
	DBHostname string `toml:"db_hostname"`
	DBPort     int    `toml:"db_port"`
	DBName     string `toml:"db_name"`
	DBUsername string `toml:"db_user"`
	DBPassword string `toml:"db_pass"`
	DBSSLMode  bool   `toml:"db_sslmode"`
	LogLevel   string `toml:"log_level"`
}

func NewConfig() *StoreConfig {
	// connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
	return &StoreConfig{
		DBHostname: "localhost",
		DBName:     "postgres",
		DBUsername: "postgres",
		DBPassword: "postgres",
		DBSSLMode:  false,
		LogLevel:   "debug",
	}
}

func ReadConfig(path string) *StoreConfig {
	c := &StoreConfig{}
	configreader.Read(path, c)
	return c
}

func (c *StoreConfig) BuildDBComplexFields() {
	sslmode := "disable"
	if c.DBSSLMode {
		sslmode = "enable"
	}

	c.DBParams = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DBHostname,
		c.DBPort,
		c.DBUsername,
		c.DBPassword,
		c.DBName,
		sslmode)

	c.DBURL = fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		c.DBUsername,
		c.DBPassword,
		c.DBHostname,
		c.DBPort,
		c.DBName,
		sslmode)
}
