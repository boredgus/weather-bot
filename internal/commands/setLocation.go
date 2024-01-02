package commands

import (
	"subscription-bot/internal/app"
	"subscription-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SetLocationCommand struct {
	ChatId   int64
	Location domain.Location `json:"location"`
}

var LocationInputDescription = "1. Click on button _\"Send my current location\"_ below.\n2. Send text message with location data in format *_latitude,longitude_* (e.g. _41.3874,2.1686_) where latitude and longitude are floating point numbers splited by comma. Latitude ranges from -90 to 90, longitude ranges from -180 to 180."

func NewSetLocationCommand(chatId int64, loc domain.Location) SetLocationCommand {
	return SetLocationCommand{ChatId: chatId, Location: loc}
}

func (c SetLocationCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, "Failed to update location. Try again later.")
	err := app.ConversationSvc().Update(domain.NewConversation(c.ChatId, domain.SetTime, domain.Subscription{Location: c.Location}))
	if err != nil {
		return msg, err
	}
	msg.Text = parseText("Provided location was saved.\n\n" + TimeInputDescription)
	msg.ParseMode = MarkdownV2Mode
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	return msg, nil
}
