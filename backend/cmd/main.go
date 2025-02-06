package main

import (
	"log"

	"github.com/djhranicky/ConcertTracker-SE-Project/cmd/api"
	"github.com/djhranicky/ConcertTracker-SE-Project/db"
	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"gorm.io/gorm"
)

func main() {
	db, err := db.NewSqliteStorage()
	if err != nil {
		log.Fatal(err)
	}

	initDatabase(db)

	server := api.NewAPIServer(":8080", db)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func initDatabase(db *gorm.DB) {
	db.AutoMigrate(&types.User{})
}
