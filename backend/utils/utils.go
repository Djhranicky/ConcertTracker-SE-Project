package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

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
