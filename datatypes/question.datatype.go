package datatypes

import (
	"fmt"
	"strings"
)

type Question struct {
	Prev    *Question
	Next    *Question
	Text    string   `db:"question"`
	Answers []string `db:"answers"`
}

func (q *Question) QuestionTemplate() string {
	return fmt.Sprintf("%s \n%s", q.Text, strings.Join(q.Answers, "\n"))
}
