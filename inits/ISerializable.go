package inits

import "github.com/jmoiron/sqlx"

type ISerializable interface {
	InitDBValue(*sqlx.DB)
}
