package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type InvalidTimeCommand struct {
	ChatId int64
	Error  error
}

func NewInvalidTimeCommand(chatId int64, er error) InvalidTimeCommand {
	return InvalidTimeCommand{ChatId: chatId, Error: er}
}

func (c InvalidTimeCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, parseText("Invalid time format was provided. Read instructions and try again."))
	msg.ParseMode = MarkdownV2Mode
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	return msg, c.Error
}
