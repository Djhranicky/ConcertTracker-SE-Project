package types

import "time"

type UserPostCreatePayload struct {
	AuthorID   uint    `json:"authorID" validate:"required"`
	Text       *string `json:"text,omitempty"`
	Type       string  `json:"type" validate:"required"`
	Rating     *uint   `json:"rating,omitempty"`
	UserPostID *uint   `json:"userPostID,omitempty"`
	IsPublic   bool    `json:"isPublic" validate:"required"`
	ConcertID  uint    `json:"concertID" validate:"required"`
}

type LikeCreatePayload struct {
	UserID     uint `json:"userID" validate:"required"`
	UserPostID uint `json:"userPostID" validate:"required"`
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
	IsLiked    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time

	UserPost UserPost `gorm:"foreignKey:UserPostID"`
	User     User     `gorm:"foreignKey:UserID"`
}
