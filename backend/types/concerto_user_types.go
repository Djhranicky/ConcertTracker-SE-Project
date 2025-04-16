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

type UserFollowPayload struct {
	UserID         uint `json:"userID" validate:"required"`
	FollowedUserID uint `json:"followedUserID" validate:"required"`
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Follow struct {
	ID             uint `gorm:"primaryKey"`
	UserID         uint `json:"userID"`
	FollowedUserID uint `json:"followedUserID"`
	IsFollowed     bool `json:"isFollowed"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	User         User `gorm:"foreignKey:UserID"`
	FollowedUser User `gorm:"foreignKey:FollowedUserID"`
}
