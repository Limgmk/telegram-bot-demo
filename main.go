package main

import (
	"log"
	"net/http"
	"net/url"
	"tgb/controller"
	"tgb/models"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

func main() {

	// 读取配置文件
	config := models.GetConfig()

	// 创建telegram 机器人API
	bot, _ = tgbotapi.NewBotAPI(config.TelegramBotToken)
	controller.SetBot(bot)

	log.Println("TelegramBotName: ", bot.Self.UserName)

	var updates tgbotapi.UpdatesChannel

	if config.Mode == "webhook" {

		log.Println("work mode is webhook")

		webhookListener := config.WebhookListener

		webhookURL, err := url.Parse(webhookListener.WebhookURL)
		if err != nil {
			log.Fatalln(err)
		}

		var path string

		uri := webhookURL.Path

		if uri == "" {
			path = "/"
		} else {
			path = uri
		}

		if webhookListener.TLS {
			go func() {
				err := http.ListenAndServeTLS(webhookListener.BindAddress,
					webhookListener.CertPath,
					webhookListener.KeyPath, nil)
				if err != nil {
					log.Fatalln(err)
				}
			}()
		} else {
			go func() {
				err := http.ListenAndServe(webhookListener.BindAddress, nil)
				if err != nil {
					log.Fatalln(err)
				}
			}()
		}

		var webhookConfig tgbotapi.WebhookConfig
		if webhookListener.TLS && webhookListener.SelfSigned {
			webhookConfig = tgbotapi.NewWebhookWithCert(webhookListener.WebhookURL, webhookListener.CertPath)
		} else {
			webhookConfig = tgbotapi.NewWebhook(webhookListener.WebhookURL)
		}
		if _, err := bot.SetWebhook(webhookConfig); err != nil {
			log.Fatalln(err)
		}

		updates = bot.ListenForWebhook(path)

	} else if config.Mode == "polling" {

		log.Println("work mode is polling")

		// 移除webhook
		if _, err := bot.RemoveWebhook(); err != nil {
			log.Fatalln(err)
		}

		// 轮询获取更新消息
		updateConfig := tgbotapi.NewUpdate(0)
		updateConfig.Timeout = 60

		var err error
		updates, err = bot.GetUpdatesChan(updateConfig)
		if err != nil {
			log.Println(err)
		}

	} else {
		log.Fatalln("ERR: work mode is error")
	}

	for update := range updates {
		controller.Process(&update)
	}
}
