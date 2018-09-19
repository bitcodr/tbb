package handler

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

//Bot var 
var Bot *tb.Bot

//NewBotHandler handler
func NewBotHandler(m *tb.Message) {
	Bot.Send(m.Sender, "good job")
}

//VisitorHandler handler
func VisitorHandler(m *tb.Message) {
	Bot.Send(m.Sender, "good job")
}

//RulesHandler handler
func RulesHandler(m *tb.Message) {
	Bot.Send(m.Sender, "good job")
}

//MyBotsHandler handler
func MyBotsHandler(m *tb.Message) {
	Bot.Send(m.Sender, "good job")
}

//SendCommentHandler handler
func SendCommentHandler(m *tb.Message) {
	Bot.Send(m.Sender, "good job")
}

//ReportAbuseHandler handler
func ReportAbuseHandler(m *tb.Message) {
	Bot.Send(m.Sender, "good job")
}

//SponsersHandler handler
func SponsersHandler(m *tb.Message) {
	Bot.Send(m.Sender, "good job")
}

//AddAdvertiseHandler handler
func AddAdvertiseHandler(m *tb.Message) {
	Bot.Send(m.Sender, "good job")
}

//HelpsHandler handler
func HelpsHandler(m *tb.Message) {
	Bot.Send(m.Sender, "good job")
}
