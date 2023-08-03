package main

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const CHAT_ID = -1001860018765

func main() {
	bot, _ := tgbotapi.NewBotAPI("6110987167:AAHyAyy3D7rGejFqsJoVlHZnZdThn5jGs20")
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case t := <-ticker.C:
			bot.Send(tgbotapi.NewMessage(CHAT_ID, t.Format("15:04:05")))
		}
	}
}
