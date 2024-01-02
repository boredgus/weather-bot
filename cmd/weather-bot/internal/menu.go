package internal

import (
	"subscription-bot/internal/commands"
)

type Json interface {
	Json() string
}

var menuCmds = []Json{
	commands.ShowSubscriptionsCommand{},
	commands.UpdateSubscriptionRequestCommand{},
	commands.CreateSubscriptionRequestCommand{},
	commands.SelectSubscriptionCommand{},
}

func MenuCommandsData() (res string) {

	res = "["
	for idx, cmd := range menuCmds {
		res += cmd.Json()
		if idx < len(menuCmds)-1 {
			res += ","
		}
	}
	return res + "]"
}
