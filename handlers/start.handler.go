package handlers

import (
	"main/markups"

	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

type StartHandlers struct {
	Bot *tele.Bot
}

func StartHandlerInit(b *tele.Bot) *StartHandlers {
	return &StartHandlers{
		Bot: b,
	}
}

func (sh *StartHandlers) StartFunction(c tele.Context) error {
	menu := &telebot.ReplyMarkup{}

	markups.SetMarkupData(menu, "start")
	return c.Send("Hello world!", menu)
}
