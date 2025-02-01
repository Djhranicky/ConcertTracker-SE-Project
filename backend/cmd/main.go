package main

import (
	"github.com/djhranicky/ConcertTracker-SE-Project/cmd/api"
	"github.com/djhranicky/ConcertTracker-SE-Project/db"
)

func main() {
	db, err := db.NewSqliteStorage()
	server := api.NewAPIServer(":8080")
	server.Run()
}
