package inits

import (
	"main/handlers"
	"main/route"

	tele "gopkg.in/telebot.v3"
)

func InitHandlers(b *tele.Bot) {
	startHandler(b)
}

func startHandler(b *tele.Bot) {
	startHandler := handlers.StartHandlerInit(b)
	route.StartRoutesInit(*startHandler)
}
