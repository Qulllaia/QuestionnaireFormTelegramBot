package queries

import (
	"main/datatypes"
	"main/inits"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	questionsHead := &datatypes.Question{}

	rows, err := qq.db.Query("SELECT question, answers FROM questions WHERE questionnaire_uid = $1", uuid)
	if err != nil {
		return nil, err
	}

	questions := questionsHead

	for rows.Next() {
		var answers []string
		err := rows.Scan(&questions.Text, pq.Array(&answers))
		if err != nil {
			return nil, err
		}

		questions.Answers = answers

		questions.Next = &datatypes.Question{
			Prev: questions,
		}
		questions = questions.Next
	}

	return questionsHead, nil
}
