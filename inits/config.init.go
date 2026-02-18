package inits

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token       string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSLMODE  string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error while loading .env file")
	}
	return &Config{
		Token:       os.Getenv("TOKEN"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
	}
}
