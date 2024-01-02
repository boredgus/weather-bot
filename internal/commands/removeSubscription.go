package commands

import (
	"subscription-bot/internal/app"
	"subscription-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type RemoveSubscriptionCommand struct {
	ChatId         int64
	SubscriptionId string `json:"subscriptionId"`
}

const RemoveSubscriptionCommandKey = "/removesubscription"

func NewRemoveSubscriptionCommand(chatId int64, subId string) RemoveSubscriptionCommand {
	return RemoveSubscriptionCommand{ChatId: chatId, SubscriptionId: subId}
}

func (c RemoveSubscriptionCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, "")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	err := app.SubscriptionSvc().Remove(c.SubscriptionId)
	if err != nil {
		msg.Text = "Failed to delete subscription. Try action again."
		return msg, err
	}
	msg.Text = "Subscription was successfully deleted."
	app.ConversationSvc().Update(domain.BaseConversation(c.ChatId))
	return msg, nil
}
