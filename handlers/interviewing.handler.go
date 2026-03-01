package handlers

import (
	. "main/datatypes"
	"main/handlers/base"

	tele "gopkg.in/telebot.v3"
)

type InterviewingHandlers struct {
	*base.BaseHandler
	questionMsg *tele.Message
	questions   *Question
}

func InterviewingHandlerInit(b *tele.Bot, questions *Question) *InterviewingHandlers {
	ih := &InterviewingHandlers{
		BaseHandler: base.BaseHandlerInit(b, "interviewing"),
		questions:   questions,
	}

	ih.Handlers = map[string]tele.HandlerFunc{}

	ih.RegisterHandlers()
	return ih
}

func (ih *InterviewingHandlers) StartMessage(c tele.Context) error {
	head := ih.questions
	msgText := head.Text + "\n"
	for _, v := range head.Answers {
		msgText = msgText + v
	}

	return c.Send(ih.Message+msgText, ih.Menu)
}
