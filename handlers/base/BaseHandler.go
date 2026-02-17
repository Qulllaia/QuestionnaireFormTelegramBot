package base

import (
	"main/markups"

	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

type BaseHandler struct {
	Bot           *tele.Bot
	SysmetButtons map[string]*tele.Btn
	Menu          *tele.ReplyMarkup
	Handlers      map[string]tele.HandlerFunc
	Message       string
}

func BaseHandlerInit(b *tele.Bot, pageName string) *BaseHandler {
	menu := &telebot.ReplyMarkup{}

	message, systemButtons := markups.SetMarkupData(menu, pageName)
	return &BaseHandler{
		Bot:           b,
		SysmetButtons: systemButtons,
		Menu:          menu,
		Message:       message,
	}
}

func (sh *BaseHandler) RegisterHandlers() {
	for name, handler := range sh.Handlers {
		if name == tele.OnText {
			sh.Bot.Handle(name, handler)
			continue
		}
		sh.Bot.Handle(sh.SysmetButtons[name], handler)
	}
}
