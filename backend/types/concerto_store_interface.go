package types

type Store interface {
	GetUserByEmail(string) (*User, error)
	GetUserByID(uint) (*User, error)
	GetUserByUsername(string) (*User, error)
	GetAllUsers() ([]User, error)
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
	GetConcertByExternalID(string) (*Concert, error)
	CreateUserPost(UserPostCreatePayload) (*UserPost, error)
	ToggleUserLike(UserLikePostPayload) error
	ToggleUserFollow(UserFollowPayload) error
	UserPostExists(string, string, string) (bool, error)
	GetNumberOfLikes(int64) (int64, error)
	GetActivityFeed(string, int64) ([]UserPostGetResponse, error)
	GetFollowersOrFollowing(string, string, int64) ([]UserFollowGetResponse, error)
	GetConcertTotalByArtist(uint) int64
	GetTourTotalByArtist(uint) int64
}
