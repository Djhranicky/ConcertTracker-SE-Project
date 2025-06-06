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

func (s *Store) GetUserByUsername(username string) (*types.User, error) {
	var user types.User
	err := s.db.First(&user, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &user, nil
}

func (s *Store) GetAllUsers() ([]types.User, error) {
	var users []types.User
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
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

func (s *Store) GetTourTotalByArtist(artistID uint) int64 {
	var count int64
	result := s.db.Model(&types.Tour{}).Where("artist_id = ?", artistID).Count(&count)
	if result.Error != nil {
		return 0
	}

	return count
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

func (s *Store) GetConcertTotalByArtist(artistID uint) int64 {
	var count int64
	result := s.db.Model(&types.Concert{}).Where("artist_id = ?", artistID).Count(&count)
	if result.Error != nil {
		return 0
	}

	return count
}

func (s *Store) GetConcertByExternalID(externalConcertID string) (*types.Concert, error) {
	var concert types.Concert
	result := s.db.Model(&types.Concert{}).Where("external_id = ?", externalConcertID).Scan(&concert)

	if result.Error != nil {
		return nil, result.Error
	}

	return &concert, nil
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
	author, err := s.GetUserByUsername(newPost.AuthorUsername)
	if err != nil {
		return nil, err
	}
	concert, err := s.GetConcertByExternalID(newPost.ExternalConcertID)
	if err != nil {
		return nil, err
	}
	post := types.UserPost{
		AuthorID:   author.ID,
		Text:       newPost.Text,
		Type:       newPost.Type,
		Rating:     newPost.Rating,
		UserPostID: newPost.UserPostID,
		IsPublic:   *newPost.IsPublic,
		ConcertID:  concert.ID,
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
	result := s.db.Raw(`SELECT
L.*
FROM likes L
JOIN Users U ON L.user_id = U.id
WHERE L.user_post_id = ? AND U.username = ?
;`, newLike.UserPostID, newLike.Username).First(&like)

	// If we found a record (no ErrRecordNotFound), delete it
	if result.Error == nil {
		return s.db.Delete(&like).Error
	}

	// If no record was found, create a new one (only if the error was "record not found")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user, err := s.GetUserByUsername(newLike.Username)
		if err != nil {
			return err
		}
		newLikeRecord := types.Likes{
			UserPostID: newLike.UserPostID,
			UserID:     user.ID,
		}
		return s.db.Create(&newLikeRecord).Error
	}

	// Return any other errors that occurred during the query
	return result.Error
}

func (s *Store) UserPostExists(authorUsername string, externalConcertID string, postType string) (bool, error) {
	var count int64
	result := s.db.Raw(`SELECT
count(*)
FROM user_posts UP
JOIN concerts C ON UP.concert_id = C.id
JOIN users U ON UP.author_id = U.id
WHERE U.username = ? AND C.external_id = ? AND UP.type = ?
;`, authorUsername, externalConcertID, postType).Scan(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func (s *Store) ToggleUserFollow(newFollow types.UserFollowPayload) error {
	var follow types.Follow

	user, err := s.GetUserByUsername(newFollow.Username)
	if err != nil {
		return err
	}

	followedUser, err := s.GetUserByUsername(newFollow.FollowedUsername)
	if err != nil {
		return err
	}

	// Try to find an existing like
	result := s.db.Model(&types.Follow{}).Where("followed_user_id = ? AND user_id = ?", followedUser.ID, user.ID).First(&follow)

	// If we found a record (no ErrRecordNotFound), delete it
	if result.Error == nil {
		return s.db.Delete(&follow).Error
	}

	// If no record was found, create a new one (only if the error was "record not found")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user, err := s.GetUserByUsername(newFollow.Username)
		if err != nil {
			return err
		}
		newFollowRecord := types.Follow{
			UserID:         user.ID,
			FollowedUserID: followedUser.ID,
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

func (s *Store) GetActivityFeed(username string, pageNumber int64) ([]types.UserPostGetResponse, error) {
	var userPosts []types.UserPostGetResponse
	result := s.db.Raw(`SELECT 
		P.id AS post_id,
		U2.name AS author_username,
		P.text,
		P.type,
		P.rating,
		P.user_post_id,
		P.is_public,
		C.external_id AS external_concert_id,
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
		WHERE U.username = 'test1'
		AND P.is_public = 1
		ORDER BY P.updated_at DESC
		LIMIT 20 OFFSET 0
	;`, username, 20*pageNumber).Scan(&userPosts)

	if result.Error != nil {
		return nil, result.Error
	}

	return userPosts, nil
}

func (s *Store) GetFollowersOrFollowing(username string, followType string, pageNum int64) ([]types.UserFollowGetResponse, error) {
	var users []types.UserFollowGetResponse
	var err error

	if followType == "followers" {
		users, err = s.getFollowers(username, users, pageNum)
	} else {
		users, err = s.getFollowing(username, users, pageNum)
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Store) getFollowers(username string, users []types.UserFollowGetResponse, pageNum int64) ([]types.UserFollowGetResponse, error) {
	pageSize := int64(20)
	result := s.db.Raw(`
		SELECT U.username
		FROM (
			SELECT
			F.user_id,
			F.created_at
			FROM follows F
			JOIN users U ON F.followed_user_id = U.id
			WHERE U.username = ?
		) F
		JOIN users U ON F.user_id = U.id
		ORDER BY F.created_at DESC
		LIMIT ? OFFSET ?;
	`, username, pageSize, pageNum*pageSize).Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (s *Store) getFollowing(username string, users []types.UserFollowGetResponse, pageNum int64) ([]types.UserFollowGetResponse, error) {
	pageSize := int64(20)
	result := s.db.Raw(`
		SELECT U.username
		FROM (
			SELECT
			F.followed_user_id,
			F.created_at
			FROM follows F
			JOIN users U ON F.user_id = U.id
			WHERE U.username = ?
		) F
		JOIN users U ON F.followed_user_id = U.id
		ORDER BY F.created_at DESC
		LIMIT ? OFFSET ?;
	`, username, pageSize, pageNum*pageSize).Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
