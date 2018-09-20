package config

import (
	"os"

	tb "gopkg.in/telegram-bot-api.v4"
)

//InitBot config
func InitBot() (*tb.BotAPI, error) {

	bot, err := tb.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	return bot, nil
}
