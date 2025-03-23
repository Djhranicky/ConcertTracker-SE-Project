package setlist

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"

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

	req, err := http.NewRequest("GET", "https://api.setlist.fm/rest/1.0/artist/f4abc0b5-3f7a-4eff-8f78-ac078dbce533/setlists", nil)
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
		store.CreateVenueIfMissing(types.Venue{
			Name:    current.Venue.Name,
			City:    current.Venue.City.Name,
			Country: current.Venue.City.Country.Name,
		})
	}
}
