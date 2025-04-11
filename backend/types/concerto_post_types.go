package types

import "time"

type Post struct {
	ID        uint    `gorm:"primaryKey"`
	AuthorID  uint    `gorm:"index"`
	Text      *string `json:"text"`
	Type      string  `json:"type"`
	Rating    *uint   `json:"rating"`
	PostID    *uint   `json:"postID"`
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
