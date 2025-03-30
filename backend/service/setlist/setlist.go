package setlist

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/joho/godotenv"
)

func ArtistSearch(url string, artist string) (*types.Artist, error) {
	err := godotenv.Load("./.env")
	if err != nil {
		return nil, err
	}

	xAPIKey := []byte(os.Getenv("SETLIST_API_KEY"))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", string(xAPIKey))

	q := req.URL.Query()
	q.Add("artistName", artist)
	q.Add("sort", "relevance")

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("no results found")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonData SetlistArtist
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, err
	}

	imageURL := ""
	if len(jsonData.Artist) > 0 && jsonData.Artist[0].Image != "" {
		imageURL = jsonData.Artist[0].Image
	}

	returnArtist := types.Artist{
		MBID:     jsonData.Artist[0].Mbid,
		Name:     jsonData.Artist[0].Name,
		ImageURL: imageURL,
	}
	return &returnArtist, nil
}

func ProcessArtistInfo(store types.Store, jsonData Artist_MBID_Setlists, artist *types.Artist) {
	numSetlists := len(jsonData.Setlist)
	for i := 0; i < numSetlists; i++ {
		current := jsonData.Setlist[i]
		var tour *types.Tour
		t, _ := time.Parse("02-01-2006", current.EventDate)
		venue := store.CreateVenueIfMissing(types.Venue{
			Name:       current.Venue.Name,
			City:       current.Venue.City.Name,
			Country:    current.Venue.City.Country.Name,
			ExternalID: current.Venue.ID,
			URL:        current.Venue.URL,
		})
		if current.Tour.Name != "" {
			tour = store.CreateTourIfMissing(types.Tour{
				Name:   current.Tour.Name,
				Artist: *artist,
			})
		}
		concert := store.CreateConcertIfMissing(types.Concert{
			Artist:            *artist,
			Tour:              tour,
			Venue:             *venue,
			Date:              t,
			ExternalID:        current.ID,
			ExternalVersionID: current.VersionID,
		})
		order := uint(0)
		numSets := len(current.Sets.Set)
		for j := 0; j < numSets; j++ {
			currSet := current.Sets.Set[j]
			numSongs := len(currSet.Song)
			for k := 0; k < numSongs; k++ {
				currSong := currSet.Song[k]
				var with *types.Artist
				var cover *types.Artist
				if currSong.With.Mbid != "" {
					with = store.CreateArtistIfMissing(types.Artist{
						MBID: currSong.With.Mbid,
						Name: currSong.With.Name,
					})
				}
				if currSong.Cover.Mbid != "" {
					cover = store.CreateArtistIfMissing(types.Artist{
						MBID: currSong.Cover.Mbid,
						Name: currSong.Cover.Name,
					})
				}
				song := store.CreateSongIfMissing(types.Song{
					Artist: *artist,
					With:   with,
					Cover:  cover,
					Name:   currSong.Name,
					Info:   currSong.Info,
					Tape:   currSong.Tape,
				})

				store.CreateConcertSongIfMissing(types.ConcertSong{
					Concert:   *concert,
					Song:      *song,
					SongOrder: order,
				})
				order++
			}
		}
	}
}
