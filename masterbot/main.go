package main

import (
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var b *tb.Bot

func main() {

	b, _ = tb.NewBot(tb.Settings{
		Token:  "603046148:AAEeB-sHu7Edynlxg3hb7rxuH08913bA_eY",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	initilizeRoutes()

	b.Start()
}
