package queries

import (
	"fmt"

	"main/datatypes"
	"main/inits"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FormFinisherQuery struct {
	db *sqlx.DB
}

func FormFinisherQueryRepo() *FormFinisherQuery {
	form := FormFinisherQuery{}
	result := inits.GetRepoInstanceDatabaseConnection[*FormFinisherQuery](&form)
	return result
}

func (ff *FormFinisherQuery) InitDBValue(db *sqlx.DB) {
	ff.db = db
}

func (ff *FormFinisherQuery) SaveQuestions(questions *datatypes.Question, uuid uuid.UUID) {
	res, err := ff.db.Exec("INSERT INTO QUESTIONS(question, answers, questionnaire_uid) VALUES($1, $2, $3)", questions.Text, questions.Answers, uuid)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res.RowsAffected())
}

func (ff *FormFinisherQuery) CreateQuestionnaire(uuid uuid.UUID, userid int64) error {
	res, err := ff.db.Exec("INSERT INTO QUESTIONNAIRE VALUES($1, $2)", userid, uuid)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(res.RowsAffected())
	return nil
}
