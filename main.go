package main

import (
	"os"
	"time"

	"github.com/amiralii/tbb/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

var b *tb.Bot

func main() {

	config.Environment()

	b, _ = tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	initilizeRoutes()

	b.Start()
}
