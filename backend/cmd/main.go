// @title Concert Tracker API
// @version 1.0
// @description API documentation for Concert Tracker.
// @host localhost:8080
// @BasePath /api
// @schemes http
package main

import (
	"log"

	"github.com/djhranicky/ConcertTracker-SE-Project/cmd/api"
	"github.com/djhranicky/ConcertTracker-SE-Project/db"
	_ "github.com/djhranicky/ConcertTracker-SE-Project/docs"
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
