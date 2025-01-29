package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetDBPassword() (string, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return "", errors.New("error loading .env file")
	}

	return os.Getenv("DB_PASSWORD"), nil
}
