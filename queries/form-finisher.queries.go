package queries

import (
	"fmt"

	"main/datatypes"

	"github.com/jmoiron/sqlx"
)

type FormFinisherQuery struct {
	DB *sqlx.DB
}

func (ff *FormFinisherQuery) SaveQuestionnaire(questions *datatypes.Question) {
	res, err := ff.DB.Exec("INSERT INTO QUESTIONS VALUES($1, $2)", questions.Text, questions.Answers)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res.RowsAffected())
}
