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
	CreateVenue(Venue) error
	CreateVenueIfMissing(Venue)
	GetVenueByName(string) (*Venue, error)
	CreateTour(Tour) error
	CreateTourIfMissing(Tour)
	GetTourByName(string) (*Tour, error)
}

type Artist struct {
	ID        uint   `gorm:"primaryKey"`
	MBID      string `gorm:"unique"`
	Name      string `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tour struct {
	ID        uint   `gorm:"primaryKey"`
	ArtistID  uint   `gorm:"index"`
	Name      string `json:"name"`
	Artist    Artist `gorm:"foreignKey:ArtistID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Venue struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Capacity  *int   `json:"capacity,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Concert struct {
	ID        uint      `gorm:"primaryKey"`
	ArtistID  uint      `gorm:"index"`
	TourID    *uint     `gorm:"index"`
	VenueID   uint      `gorm:"index"`
	Date      time.Time `json:"date"`
	Artist    Artist    `gorm:"foreignKey:ArtistID"`
	Tour      *Tour     `gorm:"foreignKey:TourID"`
	Venue     Venue     `gorm:"foreignKey:VenueID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Song struct {
	ID        uint   `gorm:"primaryKey"`
	ArtistID  uint   `gorm:"uniqueIndex:compositeIndex;index"`
	WithID    uint   `gorm:"index"`
	CoverID   uint   `gorm:"index"`
	Name      string `json:"name" gorm:"uniqueIndex:compositeIndex"`
	Info      string `json:"info"`
	Tape      bool   `json:"tape"`
	Artist    Artist `gorm:"foreignKey:ArtistID"`
	With      Artist `gorm:"foreignKey:WithID"`
	Cover     Artist `gorm:"foreignKey:CoverID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ConcertSong struct {
	ID        uint    `gorm:"primaryKey"`
	ConcertID uint    `gorm:"index"`
	SongID    uint    `gorm:"index"`
	Concert   Concert `gorm:"foreignKey:ConcertID"`
	Song      Song    `gorm:"foreignKey:SongID"`
	Order     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
