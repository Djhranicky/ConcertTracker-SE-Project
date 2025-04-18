package db

import (
	"errors"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (s *Store) CreateUserPost(newPost types.UserPostCreatePayload) (*types.UserPost, error) {
	post := types.UserPost{
		AuthorID:   newPost.AuthorID,
		Text:       newPost.Text,
		Type:       newPost.Type,
		Rating:     newPost.Rating,
		UserPostID: newPost.UserPostID,
		IsPublic:   newPost.IsPublic,
		ConcertID:  newPost.ConcertID,
	}
	result := s.db.Clauses(clause.Returning{}).Select("AuthorID", "Text", "Type", "Rating", "PostID", "IsPublic", "ConcertID").Create(&post)

	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}

func (s *Store) ToggleUserLike(newLike types.UserLikePostPayload) error {
	var like types.Likes

	// Try to find an existing like
	result := s.db.Where("user_post_id = ? AND user_id = ?", newLike.UserPostID, newLike.UserID).First(&like)

	// If we found a record (no ErrRecordNotFound), delete it
	if result.Error == nil {
		return s.db.Delete(&like).Error
	}

	// If no record was found, create a new one (only if the error was "record not found")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newLikeRecord := types.Likes{
			UserPostID: newLike.UserPostID,
			UserID:     newLike.UserID,
		}
		return s.db.Create(&newLikeRecord).Error
	}

	// Return any other errors that occurred during the query
	return result.Error
}

func (s *Store) UserPostExists(authorID, concertID uint, postType string) (bool, error) {
	var count int64
	result := s.db.Model(&types.UserPost{}).
		Where("author_id = ? AND concert_id = ? AND type = ?", authorID, concertID, postType).
		Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func (s *Store) ToggleUserFollow(newFollow types.UserFollowPayload) error {
	var result *gorm.DB
	var follow types.Follow
	result = s.db.Where(types.Follow{UserID: newFollow.UserID, FollowedUserID: newFollow.FollowedUserID}).Attrs(types.Follow{IsFollowed: true}).FirstOrCreate(&follow)

	if result.RowsAffected == 0 {
		s.db.Model(&types.Follow{}).Select("*").Where("user_id = ? AND followed_user_id = ?", follow.UserID, follow.FollowedUserID).Update("is_followed", !follow.IsFollowed)
	}

	return result.Error
}

func (s *Store) GetNumberOfLikes(userPostID int64) (int64, error) {
	var count int64
	result := s.db.Model(&types.Likes{}).Where("user_post_id = ?", userPostID).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return count, nil
}
