package handler

import (
	"io/ioutil"
	"log"

	tb "gopkg.in/telegram-bot-api.v4"
)

//StartHandler handler
func StartHandler(bot *tb.BotAPI, update *tb.Update) {

	numericKeyboard := tb.NewReplyKeyboard(
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

	startHTML, err := ioutil.ReadFile("./template/start.html")
	if err != nil {
		log.Fatal(err)
	}
	msg := tb.NewMessage(update.Message.Chat.ID, string(startHTML))
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = numericKeyboard
	msg.ParseMode = tb.ModeHTML
	bot.Send(msg)
}


//NotFoundHandler handler
func NotFoundHandler(bot *tb.BotAPI, update *tb.Update){
	notfoundtHTML, err := ioutil.ReadFile("./template/notfound.html")
	if err != nil {
		log.Fatal(err)
	}
	msg := tb.NewMessage(update.Message.Chat.ID, string(notfoundtHTML))
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = tb.ModeHTML
	bot.Send(msg)
}


