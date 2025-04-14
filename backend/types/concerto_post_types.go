package types

import "time"

type PostCreatePayload struct {
	AuthorID uint    `json:"authorID" validate:"required"`
	Text     *string `json:"text,omitempty"`
	Type     string  `json:"type" validate:"required"`
	Rating   *uint   `json:"rating,omitempty"`
	PostID   *uint   `json:"postID,omitempty"`
	IsPublic bool    `json:"isPublic" validate:"required"`
}

type Post struct {
	ID        uint    `gorm:"primaryKey"`
	AuthorID  uint    `gorm:"index"`
	Text      *string `json:"text"`
	Type      string  `json:"type"`
	Rating    *uint   `json:"rating"`
	PostID    *uint   `json:"postID"`
	IsPublic  bool    `json:"isPublic"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User  `gorm:"foreignKey:AuthorID"`
	Post *Post `gorm:"foreignKey:PostID"`
}

type Likes struct {
	ID        uint `gorm:"primaryKey"`
	PostID    uint `gorm:"index"`
	UserID    uint `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Post Post `gorm:"foreignKey:PostID"`
	User User `gorm:"foreignKey:UserID"`
}
