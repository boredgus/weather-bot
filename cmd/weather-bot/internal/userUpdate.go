package internal

import (
	"subscription-bot/internal/app"
	"subscription-bot/internal/commands"
	"subscription-bot/internal/domain"
	"subscription-bot/internal/tools"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type UserUpdate struct {
	chatId int64
}

func NewUserUpdate(chatId int64) UserUpdate {
	return UserUpdate{chatId: chatId}
}

func (u UserUpdate) ParseCallback(data string) commands.Replyable {
	conv := app.ConversationSvc().Get(u.chatId)
	switch conv.State {
	case domain.UpdateSub:
		if data == string(domain.UpdateLocation) || data == string(domain.UpdateTime) {
			return commands.NewUpdateOptioncommand(u.chatId, domain.ConversationState(data), conv.CurrentSubscription)
		}
		return commands.NewUpdateSubscriptionRequestCommand(u.chatId, data)
	case domain.RemoveSub:
		return commands.NewRemoveSubscriptionCommand(u.chatId, data)
	}
	return commands.NewUnknownCommand(u.chatId, "Failed to process callback. Start Your action again.")
}

func (u UserUpdate) ParseLocation(loc tgbotapi.Location) commands.Replyable {
	conv := app.ConversationSvc().Get(u.chatId)
	switch conv.State {
	case domain.SetLocation:
		return commands.NewSetLocationCommand(u.chatId, domain.NewLocation(loc.Latitude, loc.Longitude))

	case domain.UpdateLocation:
		sub := conv.CurrentSubscription
		sub.Location = domain.NewLocation(loc.Latitude, loc.Longitude)
		return commands.NewUpdateSubscriptionCommand(u.chatId, sub)
	}
	return commands.NewUnknownCommand(u.chatId, "Failed to process passed location. Start Your action again.")
}

func (u UserUpdate) ParseMessage(msg string) commands.Replyable {
	switch msg {
	case commands.StartCommandKey:
		return commands.NewStartCommand(u.chatId)
	case commands.CreateSubscriptionCommandKey:
		return commands.NewCreateSubscriptionRequestCommand(u.chatId)
	case commands.UpdateSubscriptionCommandKey:
		return commands.NewSelectSubscriptionCommand(u.chatId, domain.UpdateSub)
	case commands.RemoveSubscriptionCommandKey:
		return commands.NewSelectSubscriptionCommand(u.chatId, domain.RemoveSub)
	case commands.ShowSubscriptionsCommandKey:
		return commands.NewShowSubscriptionsCommand(u.chatId)
	}

	conv := app.ConversationSvc().Get(u.chatId)
	switch conv.State {
	case domain.SetLocation:
		loc, err := tools.MessageToLocation(msg)
		if err != nil {
			return commands.NewInvalidLocationCommand(u.chatId, err)
		}
		return commands.NewSetLocationCommand(u.chatId, loc)

	case domain.SetTime:
		time, err := tools.MessageToTime(msg)
		if err != nil {
			return commands.NewInvalidTimeCommand(u.chatId, err)
		}
		sub := conv.CurrentSubscription
		sub.Time = time.Format(tools.TimeTemplate)
		return commands.NewSetTimeCommand(u.chatId, sub)

	case domain.UpdateLocation:
		loc, err := tools.MessageToLocation(msg)
		if err != nil {
			return commands.NewInvalidLocationCommand(u.chatId, err)
		}
		sub := conv.CurrentSubscription
		sub.Location = loc
		return commands.NewUpdateSubscriptionCommand(u.chatId, sub)

	case domain.UpdateTime:
		time, err := tools.MessageToTime(msg)
		if err != nil {
			return commands.NewInvalidTimeCommand(u.chatId, err)
		}
		sub := conv.CurrentSubscription
		sub.Time = time.Format(tools.TimeTemplate)
		return commands.NewUpdateSubscriptionCommand(u.chatId, sub)
	}
	return commands.NewUnknownCommand(u.chatId, commands.StandartUnknownMessage)
}
