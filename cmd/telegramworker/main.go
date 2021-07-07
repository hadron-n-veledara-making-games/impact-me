package main

import (
	"flag"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/models"
	"github.com/hadron-n-veledara-making-games/impact-me/internal/app/telegrambot"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
	workersNum int
)

var numericKeyboardLockedAge = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Да, я хочу видеть своих ровесников"),
		tgbotapi.NewKeyboardButton("Нет, мне все равно на возраст тиммейта"),
	),
)

var numericKeyboardSex = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Я девушка"),
		tgbotapi.NewKeyboardButton("Я парень"),
	),
)

var numericKeyboardInterest = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Девушки"),
		tgbotapi.NewKeyboardButton("Парни"),
		tgbotapi.NewKeyboardButton("Мне все равно"),
	),
)

var numericKeyboardServer = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Европа"),
		tgbotapi.NewKeyboardButton("Америка"),
		tgbotapi.NewKeyboardButton("Азия"),
		tgbotapi.NewKeyboardButton("SAR"),
	),
)

var numericKeyboardPhotoAttach = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Да, я хочу прикрепить фото"),
		tgbotapi.NewKeyboardButton("Нет, я не хочу прикреплять фото"),
	),
)

var numericKeyboardCheck = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Да, все хорошо"),
		tgbotapi.NewKeyboardButton("Нет, я хочу заполнить анкету заново"),
	),
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
				message, err := bot.Broker.FromGOB64(string(m.Body))
				if err != nil {
					log.Fatal(err.Error())
				}

				user, err := bot.Store.User().FindByTelegramID(int(message.From.ID))
				if err != nil {
					bot.Broker.Logger.WithFields(logrus.Fields{
						"worker": id,
					}).Error(err.Error())
				}

				if user == nil {
					user = &models.User{
						TelegramID: int(message.From.ID),
					}
					if _, err := bot.Store.User().Create(user); err != nil {
						bot.Broker.Logger.WithFields(logrus.Fields{
							"worker": id,
						}).Error(err.Error())
					}
				}
				// for true {
				// 	if reflect.TypeOf(message.Text).Kind() == reflect.String && message.Text != "/start" {
				// 		msg := tgbotapi.NewMessage(message.Chat.ID, "Для начала работы введите '/start'")
				// 		bot.API.Send(msg)
				// 	} else {
				// 		msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! В этом боте ты сможешь найти себе друга для игры в Genshin Impact. Давай создадим тебе анкету.")
				// 		bot.API.Send(msg)
				// 	}
				// }

				// msg := tgbotapi.NewMessage(message.Chat.ID, "Принято")
				// msg.ReplyToMessageID = message.MessageID
				// if _, err := bot.API.Send(msg); err != nil {
				// 	bot.Broker.Logger.Fatal(err.Error())
				// }

				// if reflect.TypeOf(message.Text).Kind() == reflect.String && message.Text != "" {

				if !bot.CheckCommand(message.Command()) {
					bot.Broker.Logger.WithFields(logrus.Fields{
						"worker": id,
					}).Info(fmt.Sprintf("Recieved a non-command message %v from user %d",
						message.Text,
						message.From.ID,
					))
				} else {
					bot.Broker.Logger.WithFields(logrus.Fields{
						"worker": id,
					}).Info(fmt.Sprintf("Recieved a command message %v from user %d",
						message.Text,
						message.From.ID,
					))
				}

				// bot.Broker.Logger.WithFields(logrus.Fields{
				// 	"worker": id,
				// }).Info(fmt.Sprintf("Received a message %v from user %d",
				// 	message.Text,
				// 	message.From.ID,
				// ))
			}
		}(i)
	}
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
