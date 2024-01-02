package commands

import (
	"subscription-bot/internal/app"
	"subscription-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdateSubscriptionCommand struct {
	ChatId     int64
	UpdatedSub domain.Subscription `json:"updatedSubscription"`
}

func NewUpdateSubscriptionCommand(chatId int64, updatedSub domain.Subscription) UpdateSubscriptionCommand {
	return UpdateSubscriptionCommand{ChatId: chatId, UpdatedSub: updatedSub}
}

func (c UpdateSubscriptionCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, "Failed to update subscription. Try again later.")
	err := app.ConversationSvc().Update(domain.NewConversationWithState(c.ChatId, domain.InitialState))
	if err != nil {
		return msg, err
	}
	err = app.SubscriptionSvc().Update(c.UpdatedSub)
	if err != nil {
		return msg, err
	}
	msg.Text = "Subscription was successfully updated."
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	return msg, nil
}
