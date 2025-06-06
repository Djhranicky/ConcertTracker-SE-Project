package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

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
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, withCredentials")
	w.Header().Set("withCredentials", "true")
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

func GetArtistDataFromAPI(url string) (map[string]interface{}, error) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create new Chrome instance
	browserCtx, _ := chromedp.NewContext(ctx)

	// Initialize variables to store HTML
	var upcoming, stats string

	// Execute Chrome tasks
	err := chromedp.Run(browserCtx,
		chromedp.Navigate(url),
		chromedp.OuterHTML(`.upcomingSetlistsList`, &upcoming, chromedp.ByQuery),
		chromedp.OuterHTML(`.artistStatsTeaser`, &stats, chromedp.ByQuery),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to navigate and extract HTML: %w", err)
	}

	// Prepare result object
	result := map[string]interface{}{
		"upcoming": []map[string]string{},
		"stats":    []map[string]string{},
	}

	// Parse upcoming events
	upcomingHTML, err := goquery.NewDocumentFromReader(strings.NewReader(upcoming))
	if err != nil {
		return nil, fmt.Errorf("failed to parse upcoming HTML: %w", err)
	}

	upcomingHTML.Find(".setlist:not(.hidden)").Each(func(i int, s *goquery.Selection) {
		day := s.Find("strong.big").Text()
		month := s.Find("strong.text-uppercase").Text()
		year := strings.TrimSpace(s.Find("span.smallDateBlock span").Text())
		venue := s.Find(".content a span strong").Text()
		city := s.Find(".content span.subline span").Text()

		dateStr := fmt.Sprintf("%s-%s-%s", day, month, year) // "11-Apr-2025"

		// Parse the date
		parsedTime, err := time.Parse("02-Jan-2006", dateStr)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}

		formattedDate := parsedTime.Format("02-01-2006")

		// Extract the link from the anchor tag
		url, exists := s.Find(".content a").Attr("href")
		formattedUrl := "https://www.setlist.fm" + url[2:]

		// Extract id from the link
		id := ""
		if exists {
			trimmed := strings.TrimSuffix(url, ".html")
			split := strings.Split(trimmed, "-")
			id = split[len(split)-1]
		}

		event := map[string]string{
			"city":  city,
			"date":  formattedDate,
			"id":    id,
			"url":   formattedUrl,
			"venue": venue,
		}

		result["upcoming"] = append(result["upcoming"].([]map[string]string), event)
	})

	// Parse stats
	statsHTML, err := goquery.NewDocumentFromReader(strings.NewReader(stats))
	if err != nil {
		return nil, fmt.Errorf("failed to parse stats HTML: %w", err)
	}

	statsHTML.Find("li").Each(func(i int, s *goquery.Selection) {
		song := s.Find("a").Text()
		count := s.Find("span").Text()

		stat := map[string]string{
			"song":  song,
			"count": count,
		}

		result["stats"] = append(result["stats"].([]map[string]string), stat)
	})

	return result, nil
}
