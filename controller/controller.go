package controller

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

func SetBot(b *tgbotapi.BotAPI)  {
	bot = b
}

// Process :  处理更新消息
func Process(update *tgbotapi.Update) {

	if update.Message == nil || update.Message.Text == "" {
		return
	}

	log.Println(update.Message.Text)

	//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//msg.ReplyToMessageID = update.Message.MessageID
	//
	//if _, err := bot.Send(msg); err != nil {
	//	log.Panic(err)
	//}
}