package datatypes

import (
	"fmt"
)

type Question struct {
	Prev    *Question
	Next    *Question
	Text    string   `db:"question"`
	Answers []string `db:"answers"`
}

func (q *Question) QuestionTemplate() string {
	questionsString := ""
	for _, v := range q.Answers {
		questionsString += fmt.Sprintf(" - %s \n", v)
	}
	return fmt.Sprintf("Вопрос: %s \n%s", q.Text, questionsString)
}

func (q *Question) GetHead() *Question {
	var head *Question
	current := q
	for ; current != nil; current = current.Prev {
		head = current
	}
	return head.Next
}
