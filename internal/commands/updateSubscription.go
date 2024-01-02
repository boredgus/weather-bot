package commands

import (
	"fmt"
	"subscription-bot/internal/app"
	"subscription-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UpdateSubscriptionRequestCommand struct {
	ChatId         int64
	SubscriptionId string
}

const UpdateSubscriptionCommandKey = "/updatesubscription"

func NewUpdateSubscriptionRequestCommand(chatId int64, subId string) UpdateSubscriptionRequestCommand {
	return UpdateSubscriptionRequestCommand{ChatId: chatId, SubscriptionId: subId}
}

func (c UpdateSubscriptionRequestCommand) Json() string {
	return ToJson(UpdateSubscriptionCommandKey, "Update subscription")
}

var UpdateOptionsMarkup = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("location", string(domain.UpdateLocation))),
	tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("time", string(domain.UpdateTime))),
)

func (c UpdateSubscriptionRequestCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, "Failed to find provided subscription. Try action again from start.")
	sub, err := app.SubscriptionSvc().GetById(c.SubscriptionId)
	if err != nil {
		return msg, err
	}
	err = app.ConversationSvc().Update(domain.NewConversation(c.ChatId, domain.UpdateSub, sub))
	if err != nil {
		return msg, err
	}
	msg.Text = parseText(fmt.Sprintf("Selected subscription:\n%v\n\nSelect what You want to update:", sub.String()))
	msg.ParseMode = MarkdownV2Mode
	msg.ReplyMarkup = UpdateOptionsMarkup
	return msg, nil
}
