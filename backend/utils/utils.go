package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/djhranicky/ConcertTracker-SE-Project/service/setlist"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing Request Body in ParseJSON")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	// w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}

func SetCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}

func GetArtistSetlistsFromAPI(w http.ResponseWriter, inputURL string, mbid string, pageNum int) (*setlist.Artist_MBID_Setlists, error) {
	URL := fmt.Sprintf("%s/artist/%s/setlists?p=%d", inputURL, mbid, pageNum)

	err := godotenv.Load("./.env")
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return nil, err
	}

	xAPIKey := []byte(os.Getenv("SETLIST_API_KEY"))

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", string(xAPIKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		WriteError(w, http.StatusBadRequest, errors.New("artist not found in external API"))
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return nil, err
	}

	var jsonData setlist.Artist_MBID_Setlists
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return nil, err
	}

	return &jsonData, nil
}

func getArtistDataFromAPI(url string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var upcoming string
	var stats string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.OuterHTML(`.upcomingSetlistsList`, &upcoming, chromedp.ByQuery),
		chromedp.OuterHTML(`.artistStatsTeaser`, &stats, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatal(err)
		return
	}

	upcomingHTML, err := goquery.NewDocumentFromReader(strings.NewReader(upcoming))

	if err != nil {
		log.Fatal(err)
		return
	}

	upcomingHTML.Find(".setlist:not(.hidden)").Each(func(i int, s *goquery.Selection) {
		day := s.Find("strong.big").Text()
		month := s.Find("strong.text-uppercase").Text()
		year := strings.TrimSpace(s.Find("span.smallDateBlock span").Text())

		venue := s.Find(".content a span strong").Text()
		location := s.Find(".content span.subline span").Text()

		log.Println(day, month, year, venue, location)
	})

	statsHTML, err := goquery.NewDocumentFromReader(strings.NewReader(stats))

	if err != nil {
		log.Fatal(err)
	}

	statsHTML.Find("li").Each(func(i int, s *goquery.Selection) {
		song := s.Find("a").Text()
		count := s.Find("span").Text()
		log.Println(song, count)
	})
}
