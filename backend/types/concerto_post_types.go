package types

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	AuthorID  uint   `gorm:"index"`
	Text      string `json:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:AuthorID"`
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
