package route

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"github.com/amiralii/tbb/handler"

)

//Init config
func Init(bot *tb.Bot) {

	handler.Bot = bot

	bot.Handle(replybtn("ساخت بات جدید"), handler.NewBotHandler)

	bot.Handle(replybtn("بازاریابی"), handler.VisitorHandler)

	bot.Handle(replybtn("قوانین"), handler.RulesHandler)

	bot.Handle(replybtn("بات های من"), handler.MyBotsHandler)

	bot.Handle(replybtn("ارسال نظر"), handler.SendCommentHandler)

	bot.Handle(replybtn("گزارش تخلف"), handler.ReportAbuseHandler)

	bot.Handle(replybtn("حامیان ما"), handler.SponsersHandler)

	bot.Handle(replybtn("ثبت آگهی"), handler.AddAdvertiseHandler)

	bot.Handle(replybtn("راهنما"), handler.HelpsHandler)

	replyKeys := [][]tb.ReplyButton{
		[]tb.ReplyButton{*replybtn("ساخت بات جدید")},
		[]tb.ReplyButton{*replybtn("بازاریابی"), *replybtn("قوانین"), *replybtn("بات های من")},
		[]tb.ReplyButton{*replybtn("ارسال نظر"), *replybtn("گزارش تخلف")},
		[]tb.ReplyButton{*replybtn("حامیان ما"), *replybtn("ثبت آگهی"), *replybtn("راهنما")},
		[]tb.ReplyButton{*replybtn("حامیان ما"), *replybtn("ثبت آگهی"), *replybtn("راهنما")},
	}

	bot.Handle("/start", func(m *tb.Message) {
		startHandler(replyKeys, m, bot)
	})

}

func startHandler(replyKeys [][]tb.ReplyButton, m *tb.Message, bot *tb.Bot) {
	if !m.Private() {
		return
	}
	bot.Send(m.Sender, "<h1>good<h1>",
		&tb.ReplyMarkup{
			ReplyKeyboard:       replyKeys,
			ResizeReplyKeyboard: true,
			OneTimeKeyboard:     true,
		})
}

func replybtn(name string) *tb.ReplyButton {
	return &tb.ReplyButton{Text: name}
}
