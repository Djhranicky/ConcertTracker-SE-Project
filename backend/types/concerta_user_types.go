package types

import (
	"time"
)

type UserRegisterPayload struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type UserLoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Store interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id uint) (*User, error)
	CreateUser(User) error
	GetArtistByMBID(artist string) (*Artist, error)
	GetArtistByName(name string) (*Artist, error)
	CreateArtist(Artist) error
	CreateArtistIfMissing(Artist) *Artist
	CreateVenue(Venue) error
	CreateVenueIfMissing(Venue) *Venue
	GetVenueByName(string) (*Venue, error)
	CreateTour(Tour) error
	CreateTourIfMissing(Tour) *Tour
	GetTourByName(string) (*Tour, error)
	CreateConcertIfMissing(Concert) *Concert
	CreateSongIfMissing(Song) *Song
	CreateConcertSongIfMissing(ConcertSong) *ConcertSong
}
