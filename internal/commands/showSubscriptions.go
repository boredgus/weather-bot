package commands

import (
	"fmt"
	"subscription-bot/internal/app"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ShowSubscriptionsCommand struct {
	ChatId int64
}

const ShowSubscriptionsCommandKey = "/showsubscriptions"

func NewShowSubscriptionsCommand(chatId int64) ShowSubscriptionsCommand {
	return ShowSubscriptionsCommand{ChatId: chatId}
}

func (c ShowSubscriptionsCommand) Json() string {
	return ToJson(ShowSubscriptionsCommandKey, "Show all my subscriptions")
}

func (c ShowSubscriptionsCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, parseText(fmt.Sprintf("You have no subscriptions. Try %v to create one.", CreateSubscriptionCommandKey)))
	msg.ParseMode = MarkdownV2Mode
	subs := app.SubscriptionSvc().GetAllFor(c.ChatId)
	if len(subs) > 0 {
		msg.Text = "Your subscriptions:"
		for idx, sub := range subs {
			msg.Text += fmt.Sprintf("\n\n%v. %v", idx+1, sub.String())
		}
		msg.Text = parseText(msg.Text)
	}
	return msg, nil
}
