package main

import (
	"log"

	"github.com/djhranicky/ConcertTracker-SE-Project/cmd/api"
	"github.com/djhranicky/ConcertTracker-SE-Project/db"
)

func main() {
	database, err := db.NewSqliteStorage("gorm.db")
	if err != nil {
		log.Fatal(err)
	}

	db.InitDatabase(database)

	server := api.NewAPIServer(":8080", database)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
