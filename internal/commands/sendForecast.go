package commands

import (
	"fmt"
	"subscription-bot/internal/domain"
	"subscription-bot/pkg/weather"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type SendForecastCommand struct {
	ChatId   int64
	Location domain.Location
}

func NewSendForecastCommand(chatId int64, loc domain.Location) SendForecastCommand {
	return SendForecastCommand{ChatId: chatId, Location: loc}
}

func (c SendForecastCommand) Reply() (Message, error) {
	msg := tgbotapi.NewMessage(c.ChatId, "Failed to fetch weather forecast.")

	forecast, err := weather.NewWeatherAPI().GetWeatherForecast(c.Location)
	if err != nil {
		msg.Text = "Failed to fetch weather forecast."
		return msg, err
	}

	graphicBytes, err := forecast.RenderGraph()
	if err != nil {
		logrus.Warn("failed to render forecast graphic: %w", err)
		return msg, err
	}

	photo := tgbotapi.NewPhoto(c.ChatId, tgbotapi.FileBytes{Name: "weather_forecast", Bytes: graphicBytes})
	photo.Caption = fmt.Sprintf("Weather forecast for %v:\n", c.Location.String()) + forecast.String()
	photo.ParseMode = tgbotapi.ModeHTML
	return photo, nil
}
