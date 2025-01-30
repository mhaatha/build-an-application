package main

import (
	"log"
	"net/http"
)

func main() {
	db, err := ConnectToDB()
	if err != nil {
		log.Printf("error can't connect to db: %v", err)
	}
	server := &PlayerServer{NewPostgresPlayerStore(db)}
	log.Fatal(http.ListenAndServe(":5000", server))
}
