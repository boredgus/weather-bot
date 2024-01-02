package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type InvalidLocationCommand struct {
	ChatId int64
	Error  error
}

func NewInvalidLocationCommand(chatId int64, er error) InvalidLocationCommand {
	return InvalidLocationCommand{ChatId: chatId, Error: er}
}

func (c InvalidLocationCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, parseText("Invalid location format was provided. Read instructions and try again."))
	msg.ParseMode = MarkdownV2Mode
	msg.ReplyMarkup = LocationButtonMarkup
	return msg, c.Error
}
