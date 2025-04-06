package types

type Post struct {
	ID       uint   `gorm:"primaryKey"`
	AuthorID uint   `gorm:"index"`
	Text     string `json:"text"`

	User User `gorm:"foreignKey:AuthorID"`
}

type Likes struct {
	ID     uint `gorm:"primaryKey"`
	PostID uint `gorm:"index"`
	UserID uint `gorm:"index"`

	Post Post `gorm:"foreignKey:PostID"`
	User User `gorm:"foreignKey:UserID"`
}
