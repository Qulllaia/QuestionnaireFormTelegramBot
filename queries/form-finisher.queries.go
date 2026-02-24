package queries

import (
	"fmt"

	"main/datatypes"
	"main/inits"

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

func (ff *FormFinisherQuery) SaveQuestionnaire(questions *datatypes.Question) {
	res, err := ff.db.Exec("INSERT INTO QUESTIONS VALUES($1, $2)", questions.Text, questions.Answers)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res.RowsAffected())
}
