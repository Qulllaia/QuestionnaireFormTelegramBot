package queries

import (
	"fmt"

	"main/datatypes"
	"main/inits"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type QuestionnaireQuery struct {
	db *sqlx.DB
}

func QuestionnaireQueryRepo() *QuestionnaireQuery {
	form := QuestionnaireQuery{}
	result := inits.GetRepoInstanceDatabaseConnection[*QuestionnaireQuery](&form)
	return result
}

func (qq *QuestionnaireQuery) InitDBValue(db *sqlx.DB) {
	qq.db = db
}

func (qq *QuestionnaireQuery) GetQuestions(uuid uuid.UUID) (*datatypes.Question, error) {
	var questions *datatypes.Question = &datatypes.Question{}

	rows, err := qq.db.Query("SELECT question, answers FROM questions WHERE questionnaire_uid = $1", uuid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var answers []uint8
		err := rows.Scan(&questions.Text, &answers)
		if err != nil {
			return nil, err
		}

		for _, v := range answers {
			questions.Answers = append(questions.Answers, string(v))
		}

		fmt.Println(questions.Text, questions.Answers)

		questions.Next = &datatypes.Question{
			Prev: questions,
		}
		questions = questions.Next
	}

	var head *datatypes.Question
	fmt.Println(questions.Prev)
	if questions.Prev == nil {
		head = questions
	}
	for ; questions.Prev != nil; questions = questions.Prev {

		fmt.Println(questions.Text)
		head = questions
	}

	return head, nil
}
