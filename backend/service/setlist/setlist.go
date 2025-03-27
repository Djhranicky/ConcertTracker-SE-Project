package setlist

import (
	"encoding/json"
	"errors"
	"fmt"
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

	returnArtist := types.Artist{
		MBID: jsonData.Artist[0].Mbid,
		Name: jsonData.Artist[0].Name,
	}
	return &returnArtist, nil
}

func ProcessArtistInfo(store types.Store) {
	err := godotenv.Load("./.env")
	if err != nil {
		return
	}

	xAPIKey := []byte(os.Getenv("SETLIST_API_KEY"))

	artistMBID := "f4abc0b5-3f7a-4eff-8f78-ac078dbce533"
	artist, err := store.GetArtistByMBID(artistMBID)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.setlist.fm/rest/1.0/artist/%v/setlists", artistMBID), nil)
	if err != nil {
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", string(xAPIKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var jsonData Artist_MBID_Setlists
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return
	}

	for i := 0; i < jsonData.ItemsPerPage; i++ {
		current := jsonData.Setlist[i]
		var tour *types.Tour
		t, _ := time.Parse("02-01-2006", "17-03-2025")
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
		store.CreateConcertIfMissing(types.Concert{
			Artist:            *artist,
			Tour:              tour,
			Venue:             *venue,
			Date:              t,
			ExternalID:        current.ID,
			ExternalVersionID: current.VersionID,
		})
	}
}
