package telegrambot

import "github.com/sirupsen/logrus"

type TelegramBot struct {
	config *BotConfig
	logger *logrus.Logger
}

func New(config *BotConfig) *TelegramBot {
	return &TelegramBot{
		config: config,
		logger: logrus.New(),
	}
}

func (b *TelegramBot) Start() error {
	if err := b.configureLogger(); err != nil {
		return err
	}

	b.logger.Info("starting telegram bot")
	return nil
}

func (b *TelegramBot) configureLogger() error {
	level, err := logrus.ParseLevel(b.config.LogLevel)
	if err != nil {
		return err
	}

	b.logger.SetLevel(level)
	return nil
}
