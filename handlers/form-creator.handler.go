package handlers

import (
	"fmt"
	"strings"

	"main/handlers/base"

	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

type Question struct {
	Prev    *Question
	Next    *Question
	Text    string
	Answers []string
}

type FormCreatorHandlers struct {
	*base.BaseHandler
	questionMsg *tele.Message
	questions   *Question
}

func FormCreatorHandlerInit(b *tele.Bot) *FormCreatorHandlers {
	fc := &FormCreatorHandlers{
		BaseHandler: base.BaseHandlerInit(b, "form-creator"),
		questions: &Question{
			nil,
			nil,
			"",
			make([]string, 0),
		},
	}

	fc.Handlers = map[string]tele.HandlerFunc{
		"btn3":                    fc.FirstButtons,
		"next_question":           fc.StartMessage,
		"prev_question":           fc.PrevQuestion,
		"stop_creating_questions": fc.StopMessage,
		tele.OnText:               fc.OnQuestionEnter,
	}

	fc.RegisterHandlers()
	return fc
}

func (fc *FormCreatorHandlers) StartMessage(c tele.Context) error {
	if fc.questions.Prev != nil && (fc.questions.Text == "" || len(fc.questions.Answers) == 0) {
		err := fc.Bot.Respond(c.Callback(),
			&telebot.CallbackResponse{
				Text:      "Сначала введите данные для этого вопроса",
				ShowAlert: true,
			})
		if err != nil {
			return err
		}

		return nil
	}

	var err error
	var msg *tele.Message
	var msgText string

	if fc.questions.Next == nil {
		newQuestion := &Question{
			fc.questions,
			nil,
			"",
			make([]string, 0),
		}
		fc.questions.Next = newQuestion
		fc.questions = newQuestion
		msgText = fc.Message

	} else {
		fc.questions = fc.questions.Next
		if fc.questions.Text == "" {
			msgText = fc.Message
		} else {
			msgText = fc.QuestionTemplate(fc.questions)
		}

	}

	msg, err = c.Bot().Send(c.Recipient(), msgText, fc.Menu)
	fc.questionMsg = msg
	return err
}

func (fc *FormCreatorHandlers) PrevQuestion(c tele.Context) error {
	fc.questions = fc.questions.Prev
	msg, err := c.Bot().Send(c.Recipient(), fc.QuestionTemplate(fc.questions), fc.Menu)
	fc.questionMsg = msg
	return err
}

func (fc *FormCreatorHandlers) StopMessage(c tele.Context) error {
	var head *Question
	if fc.questions.Prev == nil {
		head = fc.questions
	}
	for ; fc.questions.Prev != nil; fc.questions = fc.questions.Prev {
		head = fc.questions
	}

	msgText := strings.Builder{}
	msgText.WriteString("Ваши итоговые вопросы: ")
	for ; head != nil; head = head.Next {
		msgText.WriteString(fmt.Sprintf("\n%s", fc.QuestionTemplate(head)))
	}

	msg, err := c.Bot().Send(c.Recipient(), msgText.String(), fc.Menu)
	fc.questionMsg = msg
	return err
}

func (fc *FormCreatorHandlers) FirstButtons(c tele.Context) error {
	if err := c.Delete(); err != nil {
		fmt.Println(err.Error())
	}

	if err := StartHandlerInit(fc.Bot).StartMessage(c); err != nil {
		fmt.Println(err.Error())
	}

	fc.Bot.Handle(tele.OnText, func(e tele.Context) error { return nil })

	return nil
}

func (fc *FormCreatorHandlers) OnQuestionEnter(c tele.Context) error {
	text := c.Text()
	if fc.questions.Text == "" {
		fc.questions.Text = fmt.Sprintf("Вопрос: %s", text)
	} else {
		fc.questions.Answers = append(fc.questions.Answers, fmt.Sprintf(" - %s", text))
	}

	_, err := fc.Bot.Edit(
		fc.questionMsg,
		fc.QuestionTemplate(fc.questions),
		fc.questionMsg.ReplyMarkup,
	)

	return err
}

func (fc *FormCreatorHandlers) QuestionTemplate(q *Question) string {
	return fmt.Sprintf("%s \n%s", q.Text, strings.Join(q.Answers, "\n"))
}
