package weather

import (
	"bytes"
	"fmt"
	"subscription-bot/config"
	"subscription-bot/internal/domain"
	"sync"
	"text/template"

	"github.com/sirupsen/logrus"
)

const ForecastRequestURL = "https://api.openweathermap.org/data/2.5/forecast/?&appid={{.AppId}}&units={{.Units}}&cnt={{.ForecastAmount}}"

type urlParams struct {
	AppId          string `json:"appId"`
	Units          string `json:"units"`
	ForecastAmount int    `json:"forecastAmount"`
}

func precompileURL(params urlParams) (string, error) {
	tmpl, err := template.New("weather_forecast_request_url").Parse(ForecastRequestURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse weather forecast request url template: %w", err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, params)
	if err != nil {
		return "", fmt.Errorf("failed to execute weather forecast request url template: %w", err)
	}
	return buffer.String(), nil
}

type WeatherForecastRequestURL struct {
	BaseURL string `json:"baseURL"`
}

const ForecastAmount = 8

var requestURL *WeatherForecastRequestURL
var once sync.Once

func NewWeatherForecastRequestURL() *WeatherForecastRequestURL {
	if requestURL == nil {
		once.Do(func() {
			baseURL, err := precompileURL(urlParams{
				AppId:          config.GetConfig().OpenWeatherAPIKey,
				Units:          "metric",
				ForecastAmount: ForecastAmount,
			})
			if err != nil {
				logrus.Warnf("failed to precompile weather forecast url: %v", err.Error())
				requestURL = &WeatherForecastRequestURL{}
			} else {
				requestURL = &WeatherForecastRequestURL{BaseURL: baseURL}
			}
		})
	}
	return requestURL
}

func (r WeatherForecastRequestURL) GetURL(loc domain.Location) string {
	return fmt.Sprintf("%v&lat=%v&lon=%v", r.BaseURL, loc.Latitude, loc.Longitude)
}
