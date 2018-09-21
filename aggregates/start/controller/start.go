package start

import (
	"io/ioutil"
	"log"

	"github.com/amiralii/tbb/config"
	tb "gopkg.in/telegram-bot-api.v4"
)

//StartHandler handler
func StartHandler(bot *tb.BotAPI, update *tb.Update) {

	numericKeyboard := tb.NewReplyKeyboard(
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.NewBot),
		),
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.Rules),
			tb.NewKeyboardButton(config.Advertisement),
			tb.NewKeyboardButton(config.MyBots),
		),
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.ReportAbuse),
			tb.NewKeyboardButton(config.SendComment),
		),
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.Help),
			tb.NewKeyboardButton(config.AddAdvertise),
			tb.NewKeyboardButton(config.Sponsers),
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
func NotFoundHandler(bot *tb.BotAPI, update *tb.Update) {
	notfoundtHTML, err := ioutil.ReadFile("./template/notfound.html")
	if err != nil {
		log.Fatal(err)
	}
	msg := tb.NewMessage(update.Message.Chat.ID, string(notfoundtHTML))
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ParseMode = tb.ModeHTML
	bot.Send(msg)
}