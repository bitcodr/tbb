package main

import (
	"log"

	"github.com/amiralii/tbb/config"
	"github.com/amiralii/tbb/route"
)

func main() {

	config.Environment()

	bot, err := config.InitBot()
	if err != nil {
		log.Fatal(err)
	}

	route.Init(&bot)

	bot.Start()
}
