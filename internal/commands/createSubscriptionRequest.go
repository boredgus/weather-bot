package commands

import (
	"subscription-bot/internal/app"
	"subscription-bot/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CreateSubscriptionRequestCommand struct {
	ChatId int64
}

const CreateSubscriptionCommandKey = "/createsubscription"

var LocationButtonMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButtonLocation("Send my current location")))

func NewCreateSubscriptionRequestCommand(chatId int64) CreateSubscriptionRequestCommand {
	return CreateSubscriptionRequestCommand{ChatId: chatId}
}

func (c CreateSubscriptionRequestCommand) Json() string {
	return ToJson(CreateSubscriptionCommandKey, "create new weather forecast subscription")
}

func (c CreateSubscriptionRequestCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, "Failed to process command")
	err := app.ConversationSvc().Update(domain.NewConversationWithState(c.ChatId, domain.SetLocation))
	if err != nil {
		return msg, err
	}
	msg.Text = parseText("To create new weather forecast subscription we need You to send location coordinates.\n\nChoose one of ways:\n" + LocationInputDescription)
	msg.ParseMode = MarkdownV2Mode
	msg.ReplyMarkup = LocationButtonMarkup
	return msg, nil
}
