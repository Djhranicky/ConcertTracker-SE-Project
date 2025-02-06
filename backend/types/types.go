package types

import (
	"time"
)

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        uint
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id uint) (*User, error)
}
