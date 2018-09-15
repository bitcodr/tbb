package main

import (
	"time"
	"log"
	
	tb "gopkg.in/tucnak/telebot.v2"
)

func main(){
	b, err := tb.NewBot(tb.Settings{
		Token:  "603046148:AAEeB-sHu7Edynlxg3hb7rxuH08913bA_eY",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "good works")
	})

	b.Start()
}