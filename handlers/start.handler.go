package handlers

import (
	"fmt"

	"main/handlers/base"

	tele "gopkg.in/telebot.v3"
)

type StartHandlers struct {
	*base.BaseHandler
}

func StartHandlerInit(b *tele.Bot) *StartHandlers {
	sh := &StartHandlers{
		BaseHandler: base.BaseHandlerInit(b, "start"),
	}

	sh.Handlers = map[string]tele.HandlerFunc{
		"btn1": sh.FirstButton,
		"btn2": sh.SecondButton,
	}

	sh.RegisterHandlers()
	return sh
}

func (sh *StartHandlers) StartMessage(c tele.Context) error {
	return c.Send(sh.Message, sh.Menu)
}

func (sh *StartHandlers) FirstButton(c tele.Context) error {
	if err := c.Delete(); err != nil {
		fmt.Println(err.Error())
	}

	if err := FormCreatorHandlerInit(sh.Bot).StartMessage(c); err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func (sh *StartHandlers) SecondButton(c tele.Context) error {
	if err := c.Delete(); err != nil {
		fmt.Println(err.Error())
	}

	if err := StartHandlerInit(sh.Bot).StartMessage(c); err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
