package route

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"github.com/amiralii/tbb/handler"
)

//Init config
func Init(b *tb.Bot) {

	b.Handle(replybtn("ساخت بات جدید"), handler.NewBotHandler)

	b.Handle(replybtn("بازاریابی"), handler.VisitorHandler)

	b.Handle(replybtn("قوانین"), handler.RulesHandler)

	b.Handle(replybtn("بات های من"), handler.MyBotsHandler)

	b.Handle(replybtn("ارسال نظر"), handler.SendCommentHandler)

	b.Handle(replybtn("گزارش تخلف"), handler.ReportAbuseHandler)

	b.Handle(replybtn("حامیان ما"), handler.SponsersHandler)

	b.Handle(replybtn("ثبت آگهی"), handler.AddAdvertiseHandler)

	b.Handle(replybtn("راهنما"), handler.HelpsHandler)

	replyKeys := [][]tb.ReplyButton{
		[]tb.ReplyButton{*replybtn("ساخت بات جدید")},
		[]tb.ReplyButton{*replybtn("بازاریابی"), *replybtn("قوانین"), *replybtn("بات های من")},
		[]tb.ReplyButton{*replybtn("ارسال نظر"), *replybtn("گزارش تخلف")},
		[]tb.ReplyButton{*replybtn("حامیان ما"), *replybtn("ثبت آگهی"), *replybtn("راهنما")},
	}

	b.Handle("/start", func(m *tb.Message) {
		startHandler(replyKeys, m)
	})

}

func startHandler(replyKeys [][]tb.ReplyButton, m *tb.Message) {
	if !m.Private() {
		return
	}
	b.Send(m.Sender, "<h1>با سلام؛<br />با استفاده از ربات ساز میتوانید :smiley::point_down:<br />1- <code>بدون نیاز به برنامه نویسی، ربات خود را بسازید.</code><br />2- <code>به تمامی اعضا ربات پیام ارسال کنید.</code><br />3- <code>لیست اعضا ربات خود را دانلود کنید.</code><br />4- <code>با تک تک اعضا ربات گفت و گو کنید.</code><br />5- <code>عضویت اجباری برای کانال خودتان در ربات طراحی کنید.</code><br />6- <code>برای ربات دکمه با مطالب و طراحی دلخواه ایجاد کنید.</code><br />7- <code>برای ربات پاسخ خودکار به متن کاربران طراحی کنید.</code><br />8- <code>آمار اعضای خود را به صورت نمودار برسی کنید.</code><br />9- <code>برای ربات خود عکس پروفایل و متن توضیحات اضافه کنید.</code><br />10- <code>و از تمامی امکانات ربات لذت ببرید... !</code></h1>",
		&tb.SendOptions{
			ReplyMarkup: &tb.ReplyMarkup{
				ReplyKeyboard:       replyKeys,
				ResizeReplyKeyboard: true,
				OneTimeKeyboard:     true,
				ForceReply:          true,
			},
			ParseMode: tb.ModeHTML,
		})

}

func replybtn(name string) *tb.ReplyButton {
	return &tb.ReplyButton{Text: name}
}
