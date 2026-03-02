package queries

import (
	"fmt"
	"strings"

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

func (ff *FormFinisherQuery) SaveQuestions(questions *datatypes.Question, uuid uuid.UUID) error {
	valuesBatch := strings.Builder{}
	for ; questions != nil; questions = questions.Next {
		if questions.Text == "" {
			break
		}

		if questions.Next == nil {
			valuesBatch.WriteString(fmt.Sprintf("('%s', '%s', '%s') \n",
				questions.Text,
				"{"+strings.Join(questions.Answers, ",")+"}",
				uuid))
			continue
		}

		valuesBatch.WriteString(fmt.Sprintf("('%s', '%s', '%s'), \n",
			questions.Text,
			"{"+strings.Join(questions.Answers, ",")+"}",
			uuid))
	}
	_, err := ff.db.Exec("INSERT INTO QUESTIONS(question, answers, questionnaire_uid) VALUES " + valuesBatch.String())
	if err != nil {
		return err
	}
	return nil
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
