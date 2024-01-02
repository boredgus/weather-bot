package commands

import (
	"subscription-bot/internal/app"
	"subscription-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdateOptionCommand struct {
	ChatId               int64
	Option               domain.ConversationState
	SubscriptionToUpdate domain.Subscription
}

func NewUpdateOptioncommand(chatId int64, opt domain.ConversationState, sub domain.Subscription) UpdateOptionCommand {
	return UpdateOptionCommand{ChatId: chatId, Option: opt, SubscriptionToUpdate: sub}
}

func (c UpdateOptionCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, "Failed to select option to update. Try again later.")
	err := app.ConversationSvc().Update(domain.NewConversation(c.ChatId, c.Option, c.SubscriptionToUpdate))
	if err != nil {
		return msg, err
	}
	msg.ParseMode = MarkdownV2Mode

	if c.Option == domain.UpdateLocation {
		msg.Text = parseText(LocationInputDescription)
		msg.ReplyMarkup = LocationButtonMarkup
	}
	if c.Option == domain.UpdateTime {
		msg.Text = parseText(TimeInputDescription)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	}
	return msg, nil
}
