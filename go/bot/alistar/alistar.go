package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CACHE_ANW    = "/tmp/anwser"
	CACHE_LAGHAB = "/tmp/laghab"
)

var callMePrefix = []string{
	"ربات بی ناموس من چطوره ؟",
	"ربات بیناموس من چطوره ؟",
	"ربات بی ناموس",
	"ربات بیناموس",
	"ربات مادرسگ",
	"ربات پدر سگ",
	"ربات سگ",
	"ربات سگ",
	"ربات علی",
	"ربات یتیم",
	"بیناموس ذخیرش کن:",
	"رباط",
	"بیناموس",
}

func randomAnswer(user *tgbotapi.Message) string {
	fmt.Printf("[%s(%s)]: %s\n", user.From.FirstName, user.From.UserName, user.Text)

	if strings.HasPrefix(user.Text, "بیناموس ذخیرش کن:") {
		f, _ := os.OpenFile(CACHE_ANW, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		f.Write([]byte(user.Text[31:] + "\n"))
		defer f.Close()

		return "هماهنگه"
	}

	if strings.HasPrefix(user.Text, "رباط بگو") {
		return user.Text[16:]
	}

	if strings.HasPrefix(user.Text, "رباط لقب یاد بگیر: ") {
		f, _ := os.OpenFile(CACHE_LAGHAB, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		f.Write([]byte(user.Text[33:] + "\n"))
		defer f.Close()

		return "تصویب شد"
	}

	if strings.HasPrefix(user.Text, "رباط لقب بگو") {
		content, _ := os.ReadFile(CACHE_LAGHAB)
		lst := strings.Split(string(content), "\n")

		return lst[rand.Intn(len(lst))]
	}

	var answer = []string{
		"همیشه و همه جا گفتم " + user.From.FirstName + " داشمه",
		"جونم رفیق. چیکار شده ؟",
		"به به سلطان " + user.From.FirstName + " خان. چه خبر از این طرفا ؟",
	}

	content, _ := os.ReadFile(CACHE_ANW)
	answer = append(answer, strings.Split(string(content), "\n")...)

	return answer[rand.Intn(len(answer))]
}

func main() {
	bot, err := tgbotapi.NewBotAPI("6110987167:AAHyAyy3D7rGejFqsJoVlHZnZdThn5jGs20")
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

		reply := false

		for _, v := range callMePrefix {
			if strings.HasPrefix(update.Message.Text, v) {
				reply = true
				break
			}
		}

		if reply {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, randomAnswer(update.Message))
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
