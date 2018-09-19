package handler

import (
	"github.com/amiralii/tbb/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

var bot, _ = config.InitBot()

//NewBotHandler handler
func NewBotHandler(m *tb.Message) {
	bot.Send(m.Sender, "good job")
}

//VisitorHandler handler
func VisitorHandler(m *tb.Message) {
	bot.Send(m.Sender, "good job")
}

//RulesHandler handler
func RulesHandler(m *tb.Message) {
	bot.Send(m.Sender, "good job")
}

//MyBotsHandler handler
func MyBotsHandler(m *tb.Message) {
	bot.Send(m.Sender, "good job")
}

//SendCommentHandler handler
func SendCommentHandler(m *tb.Message) {
	bot.Send(m.Sender, "good job")
}

//ReportAbuseHandler handler
func ReportAbuseHandler(m *tb.Message) {
	bot.Send(m.Sender, "good job")
}

//SponsersHandler handler
func SponsersHandler(m *tb.Message) {
	bot.Send(m.Sender, "good job")
}

//AddAdvertiseHandler handler
func AddAdvertiseHandler(m *tb.Message) {
	bot.Send(m.Sender, "good job")
}

//HelpsHandler handler
func HelpsHandler(m *tb.Message) {
	bot.Send(m.Sender, "good job")
}
