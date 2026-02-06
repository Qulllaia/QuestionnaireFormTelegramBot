package main

import (
	"fmt"
	"time"

	"main/inits"

	tele "gopkg.in/telebot.v3"
)

func main() {
	config := inits.InitConfig()

	b, err := tele.NewBot(tele.Settings{
		Token:  config.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		panic(fmt.Errorf("error while loading telebot: %s", err.Error()))
	}

	inits.InitHandlers(b)

	b.Start()
}
