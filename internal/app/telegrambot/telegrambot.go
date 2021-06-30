package telegrambot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/broker"
	"github.com/sirupsen/logrus"
)

type TelegramBot struct {
	config *BotConfig
	logger *logrus.Logger
	Broker *broker.Broker
	API    *tgbotapi.BotAPI
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

	return &TelegramBot{
		config: config,
		logger: logrus.New(),
		API:    _api,
		Broker: _broker,
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

		b.logger.Info("recieved message ", update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Принято")
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := b.API.Send(msg); err != nil {
			b.logger.Fatal(err.Error())
		}
		if err := b.Broker.Send([]byte(update.Message.Text)); err != nil {
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
