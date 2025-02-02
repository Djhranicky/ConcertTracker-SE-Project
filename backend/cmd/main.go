package main

import (
	"log"

	"github.com/djhranicky/ConcertTracker-SE-Project/cmd/api"
	"github.com/djhranicky/ConcertTracker-SE-Project/db"
	"github.com/djhranicky/ConcertTracker-SE-Project/types"
)

func main() {
	db, err := db.NewSqliteStorage()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&types.User{})
	email := "test@example.com"
	db.Create(&types.User{Name: "DJ", Email: &email})
	log.Printf("Added User %s %s to the DB\n", "DJ", email)

	server := api.NewAPIServer(":8080", db)
	server.Run()
}
