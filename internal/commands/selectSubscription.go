package commands

import (
	"subscription-bot/internal/app"
	"subscription-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SelectSubscriptionCommand struct {
	ChatId     int64
	StateAfter domain.ConversationState
}

func NewSelectSubscriptionCommand(chatId int64, state domain.ConversationState) SelectSubscriptionCommand {
	return SelectSubscriptionCommand{ChatId: chatId, StateAfter: state}
}

func (c SelectSubscriptionCommand) Json() string {
	return ToJson(RemoveSubscriptionCommandKey, "Remove subscription")
}

func inlineButtonRow(txt, data string) []tgbotapi.InlineKeyboardButton {
	return tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(txt, data))
}

func subscriptionsToButtons(subs []domain.Subscription) [][]tgbotapi.InlineKeyboardButton {
	isUserCrazy := len(subs) > 50
	if isUserCrazy {
		subs = subs[:50]
	}
	btns := make([][]tgbotapi.InlineKeyboardButton, len(subs))
	for idx, sub := range subs {
		btns[idx] = inlineButtonRow(sub.String(), sub.Id)
	}
	if isUserCrazy {
		btns = append([][]tgbotapi.InlineKeyboardButton{inlineButtonRow("More 50 subscriptions? You are crazy!", subs[0].Id)}, btns...)
	}
	return btns
}

func (c SelectSubscriptionCommand) Reply() (Message, error) {
	subs := app.SubscriptionSvc().GetAllFor(c.ChatId)
	msg := tgbotapi.NewMessage(c.ChatId, "You have no subscriptions for weather forecast. If You want to create one, try "+CreateSubscriptionCommandKey)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)

	if len(subs) > 0 {
		app.ConversationSvc().Update(domain.NewConversationWithState(c.ChatId, c.StateAfter))
		msg.Text = "Select subscription to work with:"
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(subscriptionsToButtons(subs)...)
	}
	return msg, nil
}
