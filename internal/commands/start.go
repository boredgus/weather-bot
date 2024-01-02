package commands

import (
	"fmt"
	"subscription-bot/internal/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartCommand struct {
	ChatId int64
}

const StartCommandKey = "/start"

func NewStartCommand(chatId int64) StartCommand {
	return StartCommand{ChatId: chatId}
}

func (c StartCommand) Reply() (Message, error) {
	app.ConversationSvc().Create(c.ChatId)
	return tgbotapi.NewMessage(c.ChatId, fmt.Sprintf("Welcome!\n\nIn this bot you can create one or more subscriptions to get weather forecast for specified location in specified time.\n\nTry %v to create subscription for daily weather forecast.", CreateSubscriptionCommandKey)), nil
}
