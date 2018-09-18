package main

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var b *tb.Bot

func main() {

	b, _ = tb.NewBot(tb.Settings{
		Token:  "",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	initilizeRoutes()

	b.Start()
}
