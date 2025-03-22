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

type SetlistArtist struct {
	Type         string `json:"type"`
	ItemsPerPage int    `json:"itemsPerPage"`
	Page         int    `json:"page"`
	Total        int    `json:"total"`
	Artist       []struct {
		Mbid           string `json:"mbid"`
		Name           string `json:"name"`
		SortName       string `json:"sortName"`
		Disambiguation string `json:"disambiguation"`
		URL            string `json:"url"`
	} `json:"artist"`
}

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
