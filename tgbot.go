package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var responses = map[string]string{
	"Правила Школы":                    "https://applicant.21-school.ru/rules",
	"Пригласить Гостя":                 "https://forms.yandex.ru/u/62dfba5e5921e8cbb5872e28/",
	"Бронирование пространств":         "https://docs.google.com/spreadsheets/d/1Q5zNrtgrJ0Bfdil65lec2dR9DIFptmwwY8-9KqwuHoM/edit",
	"Шаблон письма продления дедлайна": "Текст шаблона письма продления дедлайна",
	"Реферальная программа":            "https://docs.google.com/spreadsheets/d/1G1yPQlZQIknS0xGEWnULyeEObu2IFpEfWbrimuhbeiQ/edit",
	"Заморозка аккаунта":               "https://docs.yandex.ru/docs/view?url=ya-disk-public%3A%2F%2FyrxeeHGbVHBe8FaEiKlRWrT%2BAyWTCIFR30jRDdl8BCKJfHC5VOBZCCsR0JqfiOUqRmR%2F0fePyGwwW%2FWKW0%2FCEA%3D%3D&name=%D0%9F%D1%80%D0%B8%D0%BE%D1%81%D1%82%D0%B0%D0%BD%D0%BE%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D0%BE%D0%B1%D1%83%D1%87%D0%B5%D0%BD%D0%B8%D1%8F.docx",
}

func main() {
	bot, err := tgbotapi.NewBotAPI("7257975100:AAGqXX_TsAHAigII")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			var msg tgbotapi.MessageConfig

			if responses, ok := responses[update.Message.Text]; ok {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, responses)
			} else {

				switch update.Message.Text {
				case "/start":
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome! I am your helper bot.")

					keyboard := tgbotapi.NewReplyKeyboard(
						tgbotapi.NewKeyboardButtonRow(
							tgbotapi.NewKeyboardButton("Правила Школы"),
							tgbotapi.NewKeyboardButton("Пригласить Гостя")),
						tgbotapi.NewKeyboardButtonRow(
							tgbotapi.NewKeyboardButton("Бронирование пространств"),
							tgbotapi.NewKeyboardButton("Шаблон письма продления дедлайна")),
						tgbotapi.NewKeyboardButtonRow(
							tgbotapi.NewKeyboardButton("Реферальная программа"),
							tgbotapi.NewKeyboardButton("Заморозка аккаунта")),
					)
					msg.ReplyMarkup = keyboard
				case "/help":
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "I can help you with the following commands:\n/start - Start the bot\n/help - Display this help message")
				default:
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command")
				}
			}

			bot.Send(msg)
		}
	}
}

//t.me/assistant215252_bot
//"7257975100:AAGqXX_TsJNXqFII"
