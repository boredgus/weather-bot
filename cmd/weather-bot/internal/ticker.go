package internal

import (
	"subscription-bot/internal/app"
	"subscription-bot/internal/commands"
	"subscription-bot/internal/tools"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func RunTicker(botAPI *tgbotapi.BotAPI) {
	done := make(chan int)
	defer close(done)

	t(done, botAPI)
}

func t(done <-chan int, bot *tgbotapi.BotAPI) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

tick:
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			const limit = 50
			var startFromIdx int64 = 0
			for {
				subs := app.SubscriptionSvc().GetAll(limit, startFromIdx)
				if len(subs) == 0 {
					continue tick
				}
				startFromIdx += limit
				for _, sub := range subs {
					if t.UTC().Format(tools.TimeTemplate) == sub.Time {
						msg, err := commands.NewSendForecastCommand(sub.ChatId, sub.Location).Reply()
						if msg != nil {
							_, err := bot.Send(msg)
							if err != nil {
								logrus.Warnf("failed to send message: %v", err)
							}
						}
						if err != nil {
							logrus.Warnf("found error: %v", err)
						}
					}
				}
			}
		}
	}
}
