package handlers

import (
	"fmt"
	"strings"

	. "main/datatypes"
	"main/handlers/base"
	"main/queries"

	tele "gopkg.in/telebot.v3"
)

type FormFinisherHandlers struct {
	*base.BaseHandler
	questionMsg *tele.Message
	questions   *Question
	ffq         queries.FormFinisherQuery
}

func FormFinisherHandlerInit(b *tele.Bot, q *Question) *FormFinisherHandlers {
	var questions *Question = &Question{
		nil,
		nil,
		"",
		make([]string, 0),
	}

	if q != nil {
		questions = q
	}

	ff := &FormFinisherHandlers{
		BaseHandler: base.BaseHandlerInit(b, "form-finisher"),
		questions:   questions,
	}

	ff.Handlers = map[string]tele.HandlerFunc{
		"break": ff.ReturnToCreating,
	}

	ff.RegisterHandlers()
	return ff
}

func (ff *FormFinisherHandlers) StartMessage(c tele.Context) error {
	var head *Question
	if ff.questions.Prev == nil {
		head = ff.questions
	}
	for ; ff.questions.Prev != nil; ff.questions = ff.questions.Prev {
		head = ff.questions
	}

	msgText := strings.Builder{}
	msgText.WriteString(ff.Message)
	for ; head != nil; head = head.Next {
		msgText.WriteString(fmt.Sprintf("\n%s", head.QuestionTemplate()))
	}

	msg, err := c.Bot().Send(c.Recipient(), msgText.String(), ff.Menu)
	ff.questionMsg = msg
	return err
}

func (ff *FormFinisherHandlers) ReturnToCreating(c tele.Context) error {
	return FormCreatorHandlerInit(ff.Bot, ff.questions).StartMessage(c)
}
