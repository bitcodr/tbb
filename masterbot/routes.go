package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func initilizeRoutes() {

	newBot := tb.ReplyButton{Text: "ساخت بات جدید"}
	b.Handle(&newBot, newBotHandler)

	visitor := tb.ReplyButton{Text: "بازاریابی"}
	b.Handle(&visitor, visitorHandler)

	rules := tb.ReplyButton{Text: "قوانین"}
	b.Handle(&rules, rulesHandler)

	myBots := tb.ReplyButton{Text: "بات های من"}
	b.Handle(&myBots, myBotsHandler)

	sendComment := tb.ReplyButton{Text: "ارسال نظر"}
	b.Handle(&sendComment, sendCommentHandler)

	reportAbuse := tb.ReplyButton{Text: "گزارش تخلف"}
	b.Handle(&reportAbuse, reportAbuseHandler)

	sponsers := tb.ReplyButton{Text: "حامیان ما"}
	b.Handle(&sponsers, sponsersHandler)

	addAdvertise := tb.ReplyButton{Text: "ثبت آگهی"}
	b.Handle(&addAdvertise, addAdvertiseHandler)

	helps := tb.ReplyButton{Text: "راهنما"}
	b.Handle(&helps, helpsHandler)

	replyKeys := [][]tb.ReplyButton{
		[]tb.ReplyButton{newBot},
		[]tb.ReplyButton{myBots,rules,visitor},
		[]tb.ReplyButton{reportAbuse,sendComment},
		[]tb.ReplyButton{helps,addAdvertise,sponsers},
	}

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Hello welcome to Botsaz", &tb.ReplyMarkup{
			ReplyKeyboard: replyKeys,
			ResizeReplyKeyboard : true,
			OneTimeKeyboard:true,
		})
	})

}
