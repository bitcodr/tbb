package route

import (
	"log"

	botctl "github.com/amiralii/tbb/aggregates/bot/controller"
	startctl "github.com/amiralii/tbb/aggregates/start/controller"
	"github.com/amiralii/tbb/config"
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
		case config.StartBot:
			startctl.StartHandler(bot, &update)
		case config.NewBot:
			botctl.NewBotHandler(bot, &update)
		case config.AcceptRules:
			botctl.ConfirmIDHandler(bot, &update)
		default:
			startctl.NotFoundHandler(bot, &update)
		}
	}

}
