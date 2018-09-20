package route

import (
	"io/ioutil"
	"log"

	tb "gopkg.in/telegram-bot-api.v4"
)

var numericKeyboard = tb.NewReplyKeyboard(
	tb.NewKeyboardButtonRow(
		tb.NewKeyboardButton("ساخت بات جدید"),
	),
	tb.NewKeyboardButtonRow(
		tb.NewKeyboardButton("قوانین"),
		tb.NewKeyboardButton("بازاریابی"),
		tb.NewKeyboardButton("بات های من"),
	),
	tb.NewKeyboardButtonRow(
		tb.NewKeyboardButton("گزارش تخلف"),
		tb.NewKeyboardButton("ارسال نظر"),
	),
	tb.NewKeyboardButtonRow(
		tb.NewKeyboardButton("راهنما"),
		tb.NewKeyboardButton("ثبت آگهی"),
		tb.NewKeyboardButton("حامیان ما"),
	),
)

//Init config
func Init(bot *tb.BotAPI) {

	u := tb.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err !=nil{
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		html, err := ioutil.ReadFile("./template/start.html")
		if err != nil {
			log.Fatal(err)
		}
		msg := tb.NewMessage(update.Message.Chat.ID, string(html))
		msg.ReplyToMessageID = update.Message.MessageID
		msg.ReplyMarkup = numericKeyboard
		msg.ParseMode = tb.ModeHTML
		bot.Send(msg)
	}

}
