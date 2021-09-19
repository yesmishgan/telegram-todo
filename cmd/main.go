package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func main() {
	logger := zap.NewExample() //Change to zap.NewProduction()
	if err := InitConfig(); err != nil {
		logger.Fatal(fmt.Sprintf("error initializing configs: %s", err.Error()))
	}

	if err := godotenv.Load(); err != nil {
		logger.Fatal(fmt.Sprintf("error loading env variables: %s", err.Error()))
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		logger.Fatal(fmt.Sprintf("error initializing tg token: %s", err.Error()))
	}

	logger.Info(fmt.Sprintf("authorized on account %s", bot.Self.UserName))

	bot.Debug = true
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
