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

type UserListCreatePayload struct {
	UserID uint   `json:"userID" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type UserListAddPayload struct {
	ListID    uint  `json:"listID" validate:"required"`
	ConcertID uint  `json:"concertID" validate:"required"`
	Addition  *bool `json:"addition" validate:"required"`
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

type UserFollowGetResponse struct {
	UserName string `json:"userName"`
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

type List struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `json:"userID"`
	Name      string `json:"name"`
	Type      string `json:"type"` // Limited to "ATTENDANCE", "FAVORITES", "WISHLIST", and "USERCREATED"
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}

type ListConcert struct {
	ID        uint `gorm:"primaryKey"`
	ListID    uint `gorm:"uniqueIndex:compositeIndex" json:"listID"`
	ConcertID uint `gorm:"uniqueIndex:compositeIndex" json:"concertID"`
	CreatedAt time.Time
	UpdatedAt time.Time

	List    List    `gorm:"foreignKey:ListID"`
	Concert Concert `gorm:"foreignKey:ConcertID"`
}
