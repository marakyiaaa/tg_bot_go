package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var currentState string
var userAnswers = make(map[string]string)

func handleTemplateButton(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	switch currentState {
	case "":
		currentState = "awaitingWave"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Какая у тебя волна?")
		bot.Send(msg)
	case "awaitingWave":
		userAnswers["wave"] = update.Message.Text
		currentState = "awaitingNick"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Какой у тебя ник?")
		bot.Send(msg)
	case "awaitingNick":
		userAnswers["nick"] = update.Message.Text
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Краткое описание ситуации")
		bot.Send(msg)
	case "awaitingDescription":
		userAnswers["description"] = update.Message.Text
		currentState = ""
		template := fmt.Sprintf(
			"Шаблон письма продления дедлайна:\n"+
				"Отправляется на почту\n"+
				"Кому: kazan@21-school.ru\n"+
				"Тема: %s, %s\n"+
				"Текст сообщения: Привет, %s. Спасибо большое",
			userAnswers["wave"], userAnswers["nick"], userAnswers["description"])
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, template)
		bot.Send(msg)
	}
}
