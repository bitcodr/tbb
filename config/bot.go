package config

import (
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

//InitBot config
func InitBot() (*tb.Bot, error) {

	b, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		return nil, err
	}
	return b, nil
}
