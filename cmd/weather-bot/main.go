package main

import (
	bot "subscription-bot/cmd/weather-bot/internal"
	"subscription-bot/config"
)

func init() {
	config.InitConfig()
}

func main() {
	b, err := bot.CreateBot()
	if err != nil {
		return
	}
	go bot.InitBotHandler(b)
	bot.RunTicker(b)
}
