package main

import (
	"fmt"
	"time"

	"main/handlers"
	"main/inits"
	"main/route"

	tele "gopkg.in/telebot.v3"
)

func main() {
	config := inits.InitConfig()

	db := inits.InitDatabase(config)
	defer db.Close()

	b, err := tele.NewBot(tele.Settings{
		Token:  config.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(fmt.Errorf("error while loading telebot: %s", err.Error()))
	}

	route.StartRoutesInit(*handlers.StartHandlerInit(b))
	b.Start()
}
