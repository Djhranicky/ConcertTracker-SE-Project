package setlist

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/joho/godotenv"
)

type ArtistResponse struct {
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

func ArtistSearch(artist string) *types.Artist {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Print(err)
		return nil
	}

	xAPIKey := []byte(os.Getenv("SETLIST_API_KEY"))

	url := "https://api.setlist.fm/rest/1.0/search/artists"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return nil
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
		log.Print(err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		log.Print("No results found")
		return nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil
	}

	var jsonData ArtistResponse
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		log.Print(err)
		return nil
	}

	returnArtist := types.Artist{
		MBID: jsonData.Artist[0].Mbid,
		Name: jsonData.Artist[0].Name,
	}
	return &returnArtist
}
