package main

import (
	"flag"
	"log"

	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/telegrambot"
)

var confgiPath string

func init() {
	flag.StringVar(
		&confgiPath,
		"config-path",
		"configs/telegrambot.toml",
		"path to config file")
}

func main() {
	flag.Parse()

	bot := telegrambot.New(telegrambot.ReadConfig(confgiPath))
	if err := bot.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
