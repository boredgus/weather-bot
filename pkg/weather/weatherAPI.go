package weather

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"subscription-bot/internal/domain"
)

type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type WeatherAPI struct {
	RequestURL *WeatherForecastRequestURL
}

func NewWeatherAPI() *WeatherAPI {
	return &WeatherAPI{RequestURL: NewWeatherForecastRequestURL()}
}

var EmptyWeatherForecast = WeatherForecast{}

func (api *WeatherAPI) proceedRequest(url string) (WeatherForecast, error) {
	res, err := http.Get(url)
	if err != nil {
		return EmptyWeatherForecast, fmt.Errorf("failed to fetch weather forecast: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return EmptyWeatherForecast, fmt.Errorf("invalid status code: %v", res.StatusCode)
	}
	var b bytes.Buffer
	_, err = b.ReadFrom(res.Body)
	if err != nil {
		return EmptyWeatherForecast, fmt.Errorf("failed to read response: %w", err)
	}

	forecast := NewWeatherForecast()
	err = json.Unmarshal(b.Bytes(), &forecast)
	if err != nil {
		return EmptyWeatherForecast, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return forecast, nil
}

func (api *WeatherAPI) GetWeatherForecast(location domain.Location) (WeatherForecast, error) {
	return api.proceedRequest(api.RequestURL.GetURL(location))
}
