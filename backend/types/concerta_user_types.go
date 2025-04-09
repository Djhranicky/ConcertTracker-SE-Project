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

type Attendance struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"uniqueIndex:compositeIndex"`
	ConcertID uint `gorm:"uniqueIndex:compositeIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User    User    `gorm:"foreignKey:UserID"`
	Concert Concert `gorm:"foreignKey:ConcertID"`
}

type Favorites struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"uniqueIndex:compositeIndex"`
	ConcertID uint `gorm:"uniqueIndex:compositeIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time

	User    User    `gorm:"foreignKey:UserID"`
	Concert Concert `gorm:"foreignKey:ConcertID"`
}

type List struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}

type ListConcert struct {
	ID        uint `gorm:"primaryKey"`
	ListID    uint `gorm:"uniqueIndex:compositeIndex"`
	ConcertID uint `gorm:"uniqueIndex:compositeIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time

	List    List    `gorm:"foreignKey:ListID"`
	Concert Concert `gorm:"foreignKey:ConcertID"`
}
