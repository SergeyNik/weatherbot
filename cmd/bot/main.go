package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"weatherbot/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("Token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	newBot := telegram.NewBot(bot)
	if err := newBot.Start(); err != nil {
		log.Fatal(err)
	}
}
