package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/telegrambot"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
	workersNum int
)

func init() {
	flag.StringVar(
		&configPath,
		"config-path",
		"configs/telegrambot.toml",
		"path to config file")

	flag.IntVar(
		&workersNum,
		"workers",
		1,
		"workers amount")
}

func main() {
	flag.Parse()

	forever := make(chan bool)

	for i := 0; i < workersNum; i++ {
		go func(id int) {
			bot := telegrambot.New(telegrambot.ReadConfig(configPath))
			defer bot.Broker.Close()

			msgs, err := bot.Broker.Recieve()
			if err != nil {
				log.Fatal(err.Error())
			}

			for m := range msgs {
				bot.Broker.Logger.WithFields(logrus.Fields{
					"worker": id,
				}).Info(fmt.Sprintf("Received a message >> %s", m.Body))
			}
		}(i)
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
