package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func GetDBPassword(path string) (string, error) {
	if path != "" {
		err := godotenv.Load(path)
		if err != nil {
			return "", errors.New("error loading .env file")
		}
	}

	return os.Getenv("DB_PASSWORD"), nil
}
