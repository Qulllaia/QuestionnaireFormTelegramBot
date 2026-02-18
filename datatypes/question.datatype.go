package datatypes

import (
	"fmt"
	"strings"
)

type Question struct {
	Prev    *Question
	Next    *Question
	Text    string
	Answers []string
}

func (q *Question) QuestionTemplate() string {
	return fmt.Sprintf("%s \n%s", q.Text, strings.Join(q.Answers, "\n"))
}
