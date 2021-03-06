package main

import (
	"flag"
	"log"

	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/telegrambot"
)

var configPath string

func init() {
	flag.StringVar(
		&configPath,
		"config-path",
		"configs/telegrambot.toml",
		"path to config file")
}

func main() {
	flag.Parse()

	bot := telegrambot.New(telegrambot.ReadConfig(configPath))
	if err := bot.Listen(); err != nil {
		log.Fatal(err.Error())
	}
}
