package types

import "time"

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
