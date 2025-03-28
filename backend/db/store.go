package db

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

func (s *Store) GetUserByID(id uint) (*types.User, error) {
	var user types.User
	err := s.db.First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}

func (s *Store) CreateUser(user types.User) error {
	result := s.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) GetArtistByMBID(mbid string) (*types.Artist, error) {
	var artist types.Artist
	err := s.db.First(&artist, "mb_id = ?", mbid).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &artist, nil
}

func (s *Store) GetArtistByName(name string) (*types.Artist, error) {
	var artist types.Artist
	err := s.db.First(&artist, "lower(name) = lower(?)", name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &artist, nil
}

func (s *Store) CreateArtist(artist types.Artist) error {
	result := s.db.Create(&artist)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
