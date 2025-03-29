package db

import (
	"log"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteStorage(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to DB")
	return db, nil
}

func InitDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&types.User{},
		&types.Artist{},
		&types.Tour{},
		&types.Venue{},
		&types.Concert{},
		&types.Song{},
		&types.ConcertSong{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
