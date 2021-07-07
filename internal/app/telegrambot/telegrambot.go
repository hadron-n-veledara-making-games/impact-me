package telegrambot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/broker"
	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/store"
	"github.com/sirupsen/logrus"
)

type TelegramBot struct {
	config *BotConfig
	logger *logrus.Logger
	Broker *broker.Broker
	API    *tgbotapi.BotAPI
	Store  *store.Store
}

func New(config *BotConfig) *TelegramBot {
	_api, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Fatal(err.Error())
	}
	_broker := broker.New(config.Broker)
	if err := _broker.Open(_api.Self.UserName); err != nil {
		log.Fatal(err.Error())
	}
	_store := store.New(config.Store)
	if err := _store.Open(); err != nil {
		log.Fatal(err.Error())
	}

	return &TelegramBot{
		config: config,
		logger: logrus.New(),
		API:    _api,
		Broker: _broker,
		Store:  _store,
	}
}

func (b *TelegramBot) Listen() error {
	defer b.Broker.Close()

	if err := b.configureLogger(); err != nil {
		return err
	}
	b.configureBotAPI()

	b.logger.Info("starting telegram bot")
	b.logger.Info("authorized on account: @", b.API.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := b.API.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		b.logger.Info(fmt.Sprintf("User @%s sent message %#v", update.Message.From.UserName, update.Message.Text))

		if err := b.Broker.Send(*update.Message); err != nil {
			b.logger.Fatal(err.Error())
		}
	}
	return nil
}

func (b *TelegramBot) configureLogger() error {
	level, err := logrus.ParseLevel(b.config.LogLevel)
	if err != nil {
		return err
	}
	b.logger.SetLevel(level)

	b.logger.Formatter = &logrus.TextFormatter{
		ForceColors: true,
	}

	return nil
}

func (b *TelegramBot) configureBotAPI() {
	b.API.Debug = b.config.Debug
}

func (b *TelegramBot) CheckCommand(c string) bool {
	commands, err := b.API.GetMyCommands()
	if err != nil {
		b.logger.Error(err.Error())
	}

	for _, i := range commands {
		if c == i.Command || c == "start" {
			return true
		}
	}

	return false
}
