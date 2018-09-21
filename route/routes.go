package route

import (
	"log"

	"github.com/amiralii/tbb/config"
	"github.com/amiralii/tbb/handler"
	tb "gopkg.in/telegram-bot-api.v4"
)

//Init config
func Init(bot *tb.BotAPI) {

	updates, err := config.BotUpdates(bot)
	if err != nil {
		log.Fatal(err)
	}

	for update := range *updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Text {
		case "/start":
			handler.StartHandler(bot, &update)
		default:
			handler.NotFoundHandler(bot, &update)
		}
	}

}
