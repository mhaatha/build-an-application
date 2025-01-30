package main

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	t.Run("it returns the password stored in the mock environment", func(t *testing.T) {
		mockPassword := "MockPassword_12345$#"
		os.Setenv("DB_PASSWORD", mockPassword)
		defer os.Unsetenv("DB_PASSWORD")

		got, _ := GetDBPassword("./.env")
		want := mockPassword

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
