package inits

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	databaseInstance *sqlx.DB
	once             sync.Once
)

func InitDatabase(config *Config) *sqlx.DB {
	once.Do(func() {
		var err error
		databaseInstance, err = sqlx.Connect("postgres", connectionStringFormat(config))
		if err != nil {
			panic(err.Error())
		}
	})

	return databaseInstance
}

func GetRepoInstanceDatabaseConnection[T ISerializable](repo T) T {
	if databaseInstance == nil {
		panic("Попытка получить репозиторий при несуществующем инстансe БД")
	}

	repo.InitDBValue(databaseInstance)
	return repo
}

func connectionStringFormat(config *Config) string {
	if config.DB_PASSWORD != "" {
		return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_SSLMODE)
	}
	return fmt.Sprintf("user=%s dbname=%s sslmode=%s", config.DB_USER, config.DB_NAME, config.DB_SSLMODE)
}

type ISerializable interface {
	InitDBValue(*sqlx.DB)
}
