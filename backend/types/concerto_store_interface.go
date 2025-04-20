package types

type Store interface {
	GetUserByEmail(string) (*User, error)
	GetUserByID(uint) (*User, error)
	CreateUser(User) error
	GetArtistByMBID(string) (*Artist, error)
	GetArtistByName(string) (*Artist, error)
	CreateArtist(Artist) error
	CreateArtistIfMissing(Artist) *Artist
	CreateVenue(Venue) error
	CreateVenueIfMissing(Venue) *Venue
	GetVenueByName(string) (*Venue, error)
	CreateTour(Tour) error
	CreateTourIfMissing(Tour) *Tour
	GetTourByName(string) (*Tour, error)
	CreateConcertIfMissing(Concert) *Concert
	CreateSongIfMissing(Song) *Song
	CreateConcertSongIfMissing(ConcertSong) *ConcertSong
	CreateUserPost(UserPostCreatePayload) (*UserPost, error)
	ToggleUserLike(UserLikePostPayload) error
	ToggleUserFollow(UserFollowPayload) error
	UserPostExists(authorID, concertID uint, postType string) (bool, error)
	GetNumberOfLikes(int64) (int64, error)
	GetActivityFeed(int64, int64) ([]UserPostGetResponse, error)
	GetFollowersOrFollowing(int64, string, int64) ([]UserFollowGetResponse, error)
	CreateList(UserListCreatePayload) (*List, error)
	ToggleList(UserListAddPayload) error
}
