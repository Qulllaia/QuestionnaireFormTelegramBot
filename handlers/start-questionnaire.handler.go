package handlers

import (
	"fmt"

	"main/handlers/base"
	"main/queries"

	"github.com/google/uuid"
	tele "gopkg.in/telebot.v3"
)

type StartQuesctionnaireHandlers struct {
	*base.BaseHandler
	questionMsg *tele.Message
	sqq         queries.QuestionnaireQuery
}

func StartQuesctionnaireHandlerInit(b *tele.Bot) *StartQuesctionnaireHandlers {
	sqh := &StartQuesctionnaireHandlers{
		BaseHandler: base.BaseHandlerInit(b, "start-questionnaire"),
		sqq:         *queries.QuestionnaireQueryRepo(),
	}

	sqh.Handlers = map[string]tele.HandlerFunc{
		"backwards": sqh.GoBackwards,
		tele.OnText: sqh.OnUIDEnter,
	}

	sqh.RegisterHandlers()
	return sqh
}

func (sqh *StartQuesctionnaireHandlers) StartMessage(c tele.Context) error {
	return c.Send(sqh.Message, sqh.Menu)
}

func (sqh *StartQuesctionnaireHandlers) GoBackwards(c tele.Context) error {
	return StartHandlerInit(sqh.Bot).StartMessage(c)
}

func (sqh *StartQuesctionnaireHandlers) OnUIDEnter(c tele.Context) error {
	uid := c.Text()
	resUid, err := uuid.Parse(uid)
	if err != nil {
		return err
	}

	questions, err := sqh.sqq.GetQuestions(resUid)
	if err != nil {
		return err
	}
	fmt.Println(" ", questions.Text)

	return InterviewingHandlerInit(sqh.Bot, questions).StartMessage(c)
}
