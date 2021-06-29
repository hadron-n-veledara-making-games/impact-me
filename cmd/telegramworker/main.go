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
	defer bot.Broker.Close()

	msgs, err := bot.Broker.Recieve()
	if err != nil {
		log.Fatal(err.Error())
	}

	forever := make(chan bool)

	go func() {
		for m := range msgs {
			log.Printf("Received a message >> %s", m.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
