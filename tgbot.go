package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"strings"
)

func getHoroscope(sign string) (string, error) {
	url := "https://horo.mail.ru/prediction/" + sign + "/today/"

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch horoscope for %s", sign)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	//horoscope := doc.Find(".article__item.article__item_html").Text()
	//horoscope := doc.Find("#horoscope-aries p").Text()
	//horoscope := doc.Find(".article__item.article__item_html p").First().Text()
	horoscope := doc.Find("div.article__item.article__item_html > p:nth-of-type(1)").Text()
	if horoscope == "" {
		return "", fmt.Errorf("failed to find horoscope content")
	}

	return strings.TrimSpace(horoscope), nil
}

func main() {
	bot, err := tgbotapi.NewBotAPI("7257975100:AAGqXX_TsAHAig-Pd6cHioZOJ4RBJNXqFII")
	if err != nil {
		log.Panic(err) //!!
	}
	bot.Debug = true
	log.Printf("Account authorization %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			var msg tgbotapi.MessageConfig

			switch update.Message.Command() {
			case "start":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome c: \n I`m your horoscope bot")

				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("virgo"),
						tgbotapi.NewKeyboardButton("capricorn")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("taurus"),
						tgbotapi.NewKeyboardButton("sagittarius")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("scorpio"),
						tgbotapi.NewKeyboardButton("cancer")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("pisces"),
						tgbotapi.NewKeyboardButton("leo")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("gemini"),
						tgbotapi.NewKeyboardButton("libra")),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("aquarius"),
						tgbotapi.NewKeyboardButton("aries"),
					),
				)
				msg.ReplyMarkup = keyboard

			case "help":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "I can help you with the following commands:\n/start - Start the bot\n/help - Display this help message")

			default:
				sign := strings.ToLower(update.Message.Text)
				signMap := map[string]string{
					"virgo":       "deva",
					"capricorn":   "kozerog",
					"taurus":      "telec",
					"sagittarius": "strelec",
					"scorpio":     "scorpion",
					"cancer":      "rak",
					"pisces":      "ryby",
					"leo":         "lev",
					"gemini":      "bliznecy",
					"libra":       "vesy",
					"aquarius":    "vodolej",
					"aries":       "oven",
				}
				if zodiacSign, exists := signMap[sign]; exists {
					horoscope, err := getHoroscope(zodiacSign)
					if err != nil {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, I couldn't fetch the horoscope.")
					} else {
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, horoscope)
					}
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command")
				}
			}
			bot.Send(msg)
		}
	}
}

//t.me/assistant215252_bot
