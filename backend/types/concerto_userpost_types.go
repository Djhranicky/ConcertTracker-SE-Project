package types

import "time"

type UserPostCreatePayload struct {
	AuthorID   uint    `json:"authorID" validate:"required"`
	Text       *string `json:"text,omitempty"`
	Type       string  `json:"type" validate:"required,oneof=ATTENDED WISHLIST REVIEW LISTCREATED"`
	Rating     *uint   `json:"rating,omitempty"`
	UserPostID *uint   `json:"userPostID,omitempty"`
	IsPublic   *bool   `json:"isPublic" validate:"required"`
	ConcertID  uint    `json:"concertID" validate:"required"`
}

type UserLikePostPayload struct {
	UserID     uint `json:"userID" validate:"required"`
	UserPostID uint `json:"userPostID" validate:"required"`
}

type UserLikeGetResponse struct {
	Count int64 `json:"count"`
}

type UserPostGetResponse struct {
	PostID       uint      `json:"postID"`
	AuthorName   string    `json:"authorName"`
	Text         *string   `json:"text"`
	Type         string    `json:"type"`
	Rating       *uint     `json:"rating"`
	UserPostID   *uint     `json:"userPostID"`
	IsPublic     bool      `json:"isPublic"`
	ConcertID    uint      `json:"concertID"`
	ArtistName   string    `json:"artistName"`
	ConcertDate  time.Time `json:"concertDate"`
	TourName     string    `json:"tourName"`
	VenueName    string    `json:"venueName"`
	VenueCity    string    `json:"venueCity"`
	VenueCountry string    `json:"venueCountry"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserPost struct {
	ID         uint    `gorm:"primaryKey"`
	AuthorID   uint    `gorm:"index"`
	Text       *string `json:"text"`
	Type       string  `json:"type"`
	Rating     *uint   `json:"rating"`
	UserPostID *uint   `json:"userPostID"`
	IsPublic   bool    `json:"isPublic"`
	ConcertID  uint    `json:"concertID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User     User      `gorm:"foreignKey:AuthorID"`
	UserPost *UserPost `gorm:"foreignKey:UserPostID"`
	Concert  Concert   `gorm:"foreignKey:ConcertID"`
}

type Likes struct {
	ID         uint `gorm:"primaryKey"`
	UserPostID uint `gorm:"index"`
	UserID     uint `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	UserPost UserPost `gorm:"foreignKey:UserPostID"`
	User     User     `gorm:"foreignKey:UserID"`
}
