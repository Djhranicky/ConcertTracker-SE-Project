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
	CreatedAt time.Time
	UpdatedAt time.Time

	Artist Artist `gorm:"foreignKey:ArtistID"`
}

type Venue struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name"`
	City       string `json:"city"`
	Country    string `json:"country"`
	ExternalID string `json:"externalID"`
	URL        string `json:"url"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Concert struct {
	ID                uint      `gorm:"primaryKey"`
	ArtistID          uint      `gorm:"index"`
	TourID            *uint     `gorm:"index"`
	VenueID           uint      `gorm:"index"`
	Date              time.Time `json:"date"`
	ExternalID        string    `json:"externalID"`
	ExternalVersionID string    `json:"externalVersionID"`
	CreatedAt         time.Time
	UpdatedAt         time.Time

	Artist Artist `gorm:"foreignKey:ArtistID"`
	Tour   *Tour  `gorm:"foreignKey:TourID"`
	Venue  Venue  `gorm:"foreignKey:VenueID"`
}

type Song struct {
	ID        uint   `gorm:"primaryKey"`
	ArtistID  uint   `gorm:"uniqueIndex:compositeIndex;index"`
	WithID    *uint  `gorm:"index"`
	CoverID   *uint  `gorm:"index"`
	Name      string `json:"name" gorm:"uniqueIndex:compositeIndex"`
	Info      string `json:"info"`
	Tape      bool   `json:"tape"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Artist Artist  `gorm:"foreignKey:ArtistID"`
	With   *Artist `gorm:"foreignKey:WithID"`
	Cover  *Artist `gorm:"foreignKey:CoverID"`
}

type ConcertSong struct {
	ID        uint `gorm:"primaryKey"`
	ConcertID uint `gorm:"index"`
	SongID    uint `gorm:"index"`
	SongOrder uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Concert Concert `gorm:"foreignKey:ConcertID"`
	Song    Song    `gorm:"foreignKey:SongID"`
}

type ArtistResponse struct {
	Artist         Artist               `json:"artist"`          // Basic artist info (MBID, Name, etc.)
	ArtistURL      string               `json:"artist_url"`      // Setlist.fm artist URL
	NumberOfTours  int                  `json:"number_of_tours"` // Count of distinct tours
	TourNames      []string             `json:"tour_names"`      // List of tour names
	TotalSetlists  int                  `json:"total_setlists"`  // Total number of setlists found
	RecentSetlists []RecentSetlistEntry `json:"recent_setlists"` // Most recent setlists (max 20)
	UpcomingShows  []map[string]string  `json:"upcoming_shows"`  // Scraped upcoming show data
	TopSongs       []map[string]string  `json:"top_songs"`       // Scraped top song stats
}

type RecentSetlistEntry struct {
	ID    string `json:"id"`
	Date  string `json:"date"` // Format: "02-01-2006"
	Venue string `json:"venue"`
	City  string `json:"city"`
	URL   string `json:"url"`
}
type ConcertResponse struct {
	ID          string         `json:"id"`
	VersionID   string         `json:"version_id"`
	EventDate   string         `json:"event_date"`
	LastUpdated string         `json:"last_updated"`
	Artist      ArtistMetadata `json:"artist"`
	Venue       VenueMetadata  `json:"venue"`
	Tour        *TourMetadata  `json:"tour,omitempty"` // Optional
	Info        string         `json:"info,omitempty"` // Optional additional info
	Songs       []SongMetadata `json:"songs"`
	URL         string         `json:"url"`
}

type ArtistMetadata struct {
	MBID string `json:"mbid"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VenueMetadata struct {
	ID   string       `json:"id"`
	Name string       `json:"name"`
	City CityMetadata `json:"city"`
	URL  string       `json:"url"`
}

type CityMetadata struct {
	Name    string `json:"name"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type TourMetadata struct {
	Name string `json:"name"`
}

type SongMetadata struct {
	Name  string       `json:"name"`
	Info  string       `json:"info"`
	Tape  bool         `json:"tape"`
	Order uint         `json:"order"`
	With  *ArtistBrief `json:"with,omitempty"`
	Cover *ArtistBrief `json:"cover,omitempty"`
}

type ArtistBrief struct {
	MBID string `json:"mbid"`
	Name string `json:"name"`
}
