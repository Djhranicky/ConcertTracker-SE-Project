package db

import (
	"errors"
	"time"

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

func (s *Store) CreateArtistIfMissing(artist types.Artist) *types.Artist {
	var Exists bool
	var returnArtist types.Artist
	s.db.Raw("SELECT EXISTS(SELECT 1 FROM artists WHERE mb_id = ?)", artist.MBID).Scan(&Exists)

	if !Exists {
		s.db.Create(&artist)
	}

	s.db.First(&returnArtist, "mb_id = ?", artist.MBID)
	return &returnArtist
}

func (s *Store) CreateVenue(venue types.Venue) error {
	result := s.db.Create(&venue)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) CreateVenueIfMissing(venue types.Venue) *types.Venue {
	var Exists bool
	var returnVenue types.Venue
	s.db.Raw("SELECT EXISTS(SELECT 1 FROM venues WHERE external_id = ?)", venue.ExternalID).Scan(&Exists)

	if !Exists {
		s.db.Create(&venue)
	}

	s.db.First(&returnVenue, "external_id = ?", venue.ExternalID)
	return &returnVenue
}

func (s *Store) GetVenueByName(name string) (*types.Venue, error) {
	var venue types.Venue
	err := s.db.First(&venue, "lower(name) = lower(?)", name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &venue, nil
}

func (s *Store) CreateTour(tour types.Tour) error {
	result := s.db.Create(&tour)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Store) CreateTourIfMissing(tour types.Tour) *types.Tour {
	var Exists bool
	var returnTour types.Tour
	s.db.Raw("SELECT EXISTS(SELECT 1 FROM tours WHERE name = ?)", tour.Name).Scan(&Exists)

	if !Exists {
		s.db.Create(&tour)
	}

	s.db.First(&returnTour, "name = ?", tour.Name)
	return &returnTour
}

func (s *Store) GetTourByName(name string) (*types.Tour, error) {
	var tour types.Tour
	err := s.db.First(&tour, "lower(name) = lower(?)", name).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &tour, nil
}

func (s *Store) CreateConcertIfMissing(concert types.Concert) *types.Concert {
	var Exists bool
	var returnConcert types.Concert
	s.db.Raw("SELECT EXISTS(SELECT 1 FROM concerts WHERE external_id = ?)", concert.ExternalID).Scan(&Exists)

	if !Exists {
		s.db.Create(&concert)
	}

	s.db.First(&returnConcert, "external_id = ?", concert.ExternalID)
	return &returnConcert
}

func (s *Store) CreateSongIfMissing(song types.Song) *types.Song {
	var Exists bool
	var returnSong types.Song
	s.db.Raw("SELECT EXISTS(SELECT 1 FROM songs WHERE lower(name) = lower(?) AND artist_id = ?)", song.Name, song.Artist.ID).Scan(&Exists)

	if !Exists {
		s.db.Create(&song)
	}

	s.db.First(&returnSong, "lower(name) = lower(?) AND artist_id = ?", song.Name, song.Artist.ID)
	return &returnSong
}

func (s *Store) CreateConcertSongIfMissing(concertSong types.ConcertSong) *types.ConcertSong {
	var Exists bool
	var returnConcertSong types.ConcertSong
	s.db.Raw("SELECT EXISTS(SELECT 1 FROM concert_songs WHERE concert_id = ? AND song_id = ?)", concertSong.Concert.ID, concertSong.Song.ID).Scan(&Exists)

	if !Exists {
		s.db.Create(&concertSong)
	}

	s.db.First(&returnConcertSong, "concert_id = ? AND song_id = ?", concertSong.Concert.ID, concertSong.Song.ID)
	return &returnConcertSong
}

// GetArtistTours returns all tours for a specific artist
func (s *Store) GetArtistTours(artistID uint) ([]types.Tour, error) {
	var tours []types.Tour
	err := s.db.Where("artist_id = ?", artistID).Find(&tours).Error
	return tours, err
}

// GetArtistConcertCount returns the total number of concerts for an artist
func (s *Store) GetArtistConcertCount(artistID uint) (int, error) {
	var count int64
	err := s.db.Model(&types.Concert{}).Where("artist_id = ?", artistID).Count(&count).Error
	return int(count), err
}

// GetRecentConcerts returns the most recent concerts for an artist
func (s *Store) GetRecentConcerts(artistID uint, limit int) ([]types.ConcertInfo, error) {
	var concerts []types.Concert
	err := s.db.Where("artist_id = ? AND date <= ?", artistID, time.Now()).
		Order("date DESC").
		Limit(limit).
		Preload("Venue").
		Preload("Tour").
		Find(&concerts).Error

	if err != nil {
		return nil, err
	}

	return s.concertsToConcertInfo(concerts), nil
}

// GetUpcomingConcerts returns upcoming concerts for an artist
func (s *Store) GetUpcomingConcerts(artistID uint, limit int) ([]types.ConcertInfo, error) {
	var concerts []types.Concert
	err := s.db.Where("artist_id = ? AND date > ?", artistID, time.Now()).
		Order("date ASC").
		Limit(limit).
		Preload("Venue").
		Preload("Tour").
		Find(&concerts).Error

	if err != nil {
		return nil, err
	}

	return s.concertsToConcertInfo(concerts), nil
}

// Helper function to convert Concert slice to ConcertInfo slice
func (s *Store) concertsToConcertInfo(concerts []types.Concert) []types.ConcertInfo {
	result := make([]types.ConcertInfo, len(concerts))
	for i, concert := range concerts {
		info := types.ConcertInfo{
			ID:        concert.ID,
			Date:      concert.Date,
			VenueName: concert.Venue.Name,
			City:      concert.Venue.City,
			Country:   concert.Venue.Country,
		}
		if concert.Tour != nil {
			info.TourName = concert.Tour.Name
		}
		result[i] = info
	}
	return result
}
