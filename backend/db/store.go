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
		IsPublic:   *newPost.IsPublic,
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
	var follow types.Follow

	// Try to find an existing like
	result := s.db.Where("followed_user_id = ? AND user_id = ?", newFollow.FollowedUserID, newFollow.UserID).First(&follow)

	// If we found a record (no ErrRecordNotFound), delete it
	if result.Error == nil {
		return s.db.Delete(&follow).Error
	}

	// If no record was found, create a new one (only if the error was "record not found")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		newFollowRecord := types.Follow{
			UserID:         newFollow.UserID,
			FollowedUserID: newFollow.FollowedUserID,
		}
		return s.db.Create(&newFollowRecord).Error
	}

	// Return any other errors that occurred during the query
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

func (s *Store) GetActivityFeed(userID int64, pageNumber int64) ([]types.UserPostGetResponse, error) {
	var userPosts []types.UserPostGetResponse
	result := s.db.Raw(`SELECT 
		P.id AS post_id,
		U2.name AS author_name,
		P.text,
		P.type,
		P.rating,
		P.user_post_id,
		P.is_public,
		P.concert_id,
		A.name AS artist_name,
		C.date AS concert_date,
		T.name AS tour_name,
		V.name AS venue_name,
		V.city AS venue_city,
		V.country AS venue_country,
		P.created_at,
		P.updated_at
		FROM users U
		JOIN follows F ON U.id = F.user_id
		JOIN user_posts P ON F.followed_user_id = P.author_id
		JOIN users U2 ON F.followed_user_id = U2.id
		JOIN concerts C ON P.concert_id = C.id
		JOIN tours T ON C.tour_id = T.id
		JOIN venues V ON C.venue_id = V.id
		JOIN artists A ON C.artist_id = A.id
		WHERE U.id = ?
		AND P.is_public = 1
		ORDER BY P.updated_at DESC
		LIMIT 20 OFFSET ?
	;`, userID, 20*pageNumber).Scan(&userPosts)
	if result.Error != nil {
		return nil, result.Error
	}

	return userPosts, nil
}
