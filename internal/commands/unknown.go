package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UnknownCommand struct {
	Message string
	ChatId  int64
}

const StandartUnknownMessage = "Unknow command provided. Check menu to see a list of supported commands."

func NewUnknownCommand(chatId int64, msg string) UnknownCommand {
	return UnknownCommand{ChatId: chatId, Message: msg}
}

func (c UnknownCommand) Reply() (Message, error) {
	return tgbotapi.NewMessage(c.ChatId, c.Message), nil
}
