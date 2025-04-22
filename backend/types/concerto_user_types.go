package types

import (
	"time"
)

type UserRegisterPayload struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type UserLoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUsernamePayload struct {
	Username string `json:"username"`
}

type UserFollowPayload struct {
	Username       string `json:"username" validate:"required"`
	FollowedUserID uint   `json:"followedUserID" validate:"required"`
}

type UserInfoResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique"`
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
	CreatedAt      time.Time
	UpdatedAt      time.Time

	User         User `gorm:"foreignKey:UserID"`
	FollowedUser User `gorm:"foreignKey:FollowedUserID"`
}
