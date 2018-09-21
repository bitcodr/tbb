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

//BotUpdates config
func BotUpdates(bot *tb.BotAPI) (*tb.UpdatesChannel, error) {
	u := tb.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return nil, err
	}
	return &updates, nil
}
