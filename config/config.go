package config

import (
	log "github.com/sirupsen/logrus"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type config struct {
	TelegramToken     string `env:"TELEGRAM_TOKEN"`
	OpenWeatherAPIKey string `env:"OPEN_WEATHER_API_KEY"`
	MongoDBUsername   string `env:"MONGO_DB_USERNAME"`
	MongoDBPassword   string `env:"MONGO_DB_PASSWORD"`
	GoogleMapsAPIKey  string `env:"GOOGLE_MAPS_API_KEY"`
	DbServer          string `env:"DB_CONTAINER"`
}

func LoadEnvFile(envFilePath string) {
	if err := godotenv.Load(envFilePath); err != nil {
		log.Info("failed to load env file", err)
	}
}

func GetConfig() config {
	cfg := config{}

	if err := env.Parse(&cfg); err != nil {
		log.Error("failed to load env file", err)
	}
	return cfg
}

func InitConfig() {
	LoadEnvFile(".env")
	log.Infof("env config: %+v\n", GetConfig())
}
