package inits

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error while loading .env file")
	}
	return &Config{
		Token: os.Getenv("TOKEN"),
	}
}
