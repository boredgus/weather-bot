package commands

import (
	"fmt"
	"subscription-bot/internal/app"
	"subscription-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SetTimeCommand struct {
	ChatId      int64
	SubWithTime domain.Subscription `json:"subWithTime"`
}

var TimeInputDescription = "Send text message with time when You want to receive weather forecast. Format should be like *_hh:mm_*, where _hh_ is hours and _mm_ is minutes in 24h format (e.g. _16:15_ or _08:05_).\n_Make note subscription will be sent in UTC+00:00_"

func NewSetTimeCommand(chatId int64, sub domain.Subscription) SetTimeCommand {
	return SetTimeCommand{ChatId: chatId, SubWithTime: sub}
}

func (c SetTimeCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, "Failed to create subscription. Try again later.")
	err := app.SubscriptionSvc().Insert(domain.NewSubscription(c.ChatId, c.SubWithTime.Time, c.SubWithTime.Location))
	if err == app.SubscriptionLimitError {
		msg.Text = "Sorry, you cannot create new subscription. You already have 50 subscriptions. Limit of subscriptions has reached.\n\nTry to update existed subscription or delete unneeded subscription and create new one."
		return msg, err
	}
	if err != nil {
		return msg, err
	}
	err = app.ConversationSvc().Update(domain.NewConversationWithState(c.ChatId, domain.InitialState))
	if err != nil {
		return msg, err
	}
	msg.Text = parseText(fmt.Sprintf("Subscription was successfully added.\n\nYou will receive weather forecast every day at _%v_ for location _%s_.\n\nYou can update subscription with %v or delete with %v.", c.SubWithTime.Time, c.SubWithTime.Location.String(), UpdateSubscriptionCommandKey, RemoveSubscriptionCommandKey))
	msg.ParseMode = MarkdownV2Mode
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
	return msg, nil
}
