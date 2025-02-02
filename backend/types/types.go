package types

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Email     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
}
