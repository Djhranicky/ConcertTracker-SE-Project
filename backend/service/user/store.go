package user

import (
	"errors"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	var user types.User
	err := s.db.First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}
