package bot

import (
	"io/ioutil"
	"log"

	"github.com/amiralii/tbb/config"
	tb "gopkg.in/telegram-bot-api.v4"
)

//NewBotHandler handler
func NewBotHandler(bot *tb.BotAPI, update *tb.Update) {

	// if cell_phone = nil && confirm_rules == false{
	// 	ruleHandler(bot, update)
	// }

	// if cell_phone = nil && confirm_rules == true{
	// 	ConfirmIDHandler(bot, update)
	// }
	
	// if cell_phone != nil && confirm_rules == true{
	// 	createBotHandler(bot, update)
	// }
}

func ruleHandler(bot *tb.BotAPI, update *tb.Update) {
	numericKeyboard := tb.NewReplyKeyboard(
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.RulesPictures),
		),
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.AcceptRules),
		),
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.PreviusBtn),
		),
	)
	rulesHTML, err := ioutil.ReadFile("./template/rules.html")
	if err != nil {
		log.Fatal(err)
	}
	msg := tb.NewMessage(update.Message.Chat.ID, string(rulesHTML))
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = numericKeyboard
	msg.ParseMode = tb.ModeHTML
	bot.Send(msg)
}

func createBotHandler(bot *tb.BotAPI, update *tb.Update) {
	numericKeyboard := tb.NewReplyKeyboard(
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.CreateBotVideo),
		),
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.PreviusBtn),
		),
	)

	newBotHTML, err := ioutil.ReadFile("./template/newBot.html")
	if err != nil {
		log.Fatal(err)
	}
	msg := tb.NewMessage(update.Message.Chat.ID, string(newBotHTML))
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = numericKeyboard
	msg.ParseMode = tb.ModeHTML
	bot.Send(msg)
}

//ConfirmIDHandler handler
func ConfirmIDHandler(bot *tb.BotAPI, update *tb.Update) {
	numericKeyboard := tb.NewReplyKeyboard(
		tb.NewKeyboardButtonRow(
			tb.NewKeyboardButton(config.ConfirmID),
		),
	)
	confirmRulesHTML, err := ioutil.ReadFile("./template/confirmRules.html")
	if err != nil {
		log.Fatal(err)
	}
	msg := tb.NewMessage(update.Message.Chat.ID, string(confirmRulesHTML))
	msg.ReplyToMessageID = update.Message.MessageID
	msg.ReplyMarkup = numericKeyboard
	msg.ParseMode = tb.ModeHTML
	bot.Send(msg)
}
