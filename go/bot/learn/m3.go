package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		update.Message.Text = "[                                        ]"
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msgM, err := bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}

		for i := 1; i <= 10; i++ {
			update.Message.Text = strings.Replace(update.Message.Text, " ", "=", 2)
			update.Message.Text = strings.Replace(update.Message.Text, " ", "", 2)

			percentage := i * 100 / 10
			output := fmt.Sprintf("%%%d%s", percentage, update.Message.Text)
			edit := tgbotapi.NewEditMessageText(
				update.Message.Chat.ID,
				msgM.MessageID,
				output)
			bot.Send(edit)
		}
	}
}
