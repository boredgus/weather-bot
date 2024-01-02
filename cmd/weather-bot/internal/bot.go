package internal

import (
	"fmt"
	"subscription-bot/config"
	"subscription-bot/internal/commands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func CreateBot() (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.GetConfig().TelegramToken)
	if err != nil {
		return nil, fmt.Errorf("failed to create bot api: %w", err)
	}
	setCommandMenu(bot)
	return bot, nil
}

func setCommandMenu(bot *tgbotapi.BotAPI) {
	_, err := bot.MakeRequest("setMyCommands", tgbotapi.Params{"commands": MenuCommandsData()})
	if err != nil {
		logrus.Warn("failed to set commands: " + err.Error())
	}
}

func InitBotHandler(bot *tgbotapi.BotAPI) {
	for update := range bot.GetUpdatesChan(tgbotapi.UpdateConfig{Timeout: 100}) {
		var msg commands.Message
		var err error
		if callback := update.CallbackQuery; callback != nil {
			msg, err = NewUserUpdate(callback.From.ID).ParseCallback(callback.Data).Reply()
		} else {
			upd := NewUserUpdate(update.Message.Chat.ID)
			if location := update.Message.Location; location != nil {
				msg, err = upd.ParseLocation(*location).Reply()
			} else {
				msg, err = upd.ParseMessage(update.Message.Text).Reply()
			}
		}
		if msg != nil {
			_, err = bot.Send(msg)
			if err != nil {
				logrus.Warnf("failed to send message: %v", err)
			}
		}
		if err != nil {
			logrus.Warnf("%v", err)
		}
	}
}
