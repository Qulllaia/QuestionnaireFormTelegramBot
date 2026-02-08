package handlers

import (
	"fmt"

	"main/handlers/base"

	tele "gopkg.in/telebot.v3"
)

type FormCreatorHandlers struct {
	*base.BaseHandler
}

func FormCreatorHandlerInit(b *tele.Bot) *FormCreatorHandlers {
	fc := &FormCreatorHandlers{
		BaseHandler: base.BaseHandlerInit(b, "form-creator"),
	}

	fc.Handlers = map[string]tele.HandlerFunc{
		"btn3": fc.FirstButton,
	}

	fc.RegisterHandlers()
	return fc
}

func (fc *FormCreatorHandlers) StartMessage(c tele.Context) error {
	return c.Send(fc.Message, fc.Menu)
}

func (fc *FormCreatorHandlers) FirstButton(c tele.Context) error {
	if err := c.Delete(); err != nil {
		fmt.Println(err.Error())
	}

	if err := StartHandlerInit(fc.Bot).StartMessage(c); err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
