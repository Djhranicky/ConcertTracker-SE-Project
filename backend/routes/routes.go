package routes

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
	"sort"
	"time"

	_ "github.com/djhranicky/ConcertTracker-SE-Project/docs"
	"github.com/djhranicky/ConcertTracker-SE-Project/service/auth"
	"github.com/djhranicky/ConcertTracker-SE-Project/service/setlist"
	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/djhranicky/ConcertTracker-SE-Project/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

const baseURL = "https://api.setlist.fm/rest/1.0"

type Handler struct {
	Store types.Store
}

func NewHandler(store types.Store) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.handleHome).Methods("GET")
	router.HandleFunc("/login", h.handleLogin).Methods("POST", "OPTIONS")
	router.HandleFunc("/register", h.handleRegister).Methods("POST", "OPTIONS")
	router.HandleFunc("/validate", h.handleValidate).Methods("GET", "OPTIONS")
	router.HandleFunc("/artist", h.handleArtist(baseURL)).Methods("GET", "OPTIONS")
	router.HandleFunc("/import", h.handleArtistImport(baseURL)).Methods("GET", "OPTIONS")
	router.HandleFunc("/concert", h.handleConcert(baseURL)).Methods("GET", "OPTIONS")
	router.HandleFunc("/userpost", h.handleUserPost()).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/like", h.handleUserLike()).Methods("GET", "POST", "OPTIONS")
	router.HandleFunc("/follow", h.handleUserFollow()).Methods("GET", "POST", "OPTIONS")

	// Serve Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

// @Summary Home Route
// @Description Returns a simple Hello World message
// @Tags Home
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router / [get]
func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"hello world"}`))
}

// @Summary Login user
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body types.UserLoginPayload true "Login Payload"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid email or password"
// @Router /login [post]
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

	utils.SetCORSHeaders(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var user types.UserLoginPayload
	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err := h.Store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	// load jwt token from .env
	err = godotenv.Load("./.env")
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("could not load .env, error %v", err))
		return
	}

	secret := []byte(os.Getenv("JWT_SECRET"))
	token, err := auth.CreateJWT(secret, u.ID, 3600*24*31)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// implement jwt
	auth.SetJWTCookie(w, token)
	utils.WriteJSON(w, http.StatusOK, nil)
}

// @Summary Register user
// @Description Registers a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body types.UserRegisterPayload true "Register Payload"
// @Success 201 {string} string "User registered successfully"
// @Failure 400 {string} string "Invalid payload or user already exists"
// @Router /register [post]
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	utils.SetCORSHeaders(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// get JSON payload
	var payload types.UserRegisterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate fields
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// check if the user already exists
	_, err := h.Store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	// if not, create new user
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	err = h.Store.CreateUser(types.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

// @Summary Validate user session
// @Description Verifies if a user's session cookie contains an authenticated token
// @Tags Auth
// @Produce json
// @Success 200 {string} string "user session validated"
// @Failure 401 {string} string "missing or invalid authorization token"
// @Router /validate [get]
func (h *Handler) handleValidate(w http.ResponseWriter, r *http.Request) {
	utils.SetCORSHeaders(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	err := auth.VerifyJWTCookie(auth.GetJWTCookie(r))
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"user session validated"}`))
}

// @Summary Serve information for a given artist
// @Description Gets information for requested artist. If information does not exist in database, it is retrieved from setlist.fm API and entered into database
// @Tags Artist
// @Param name path string true "Artist Name"
// @Produce json
// @Success 200 {object} types.ArtistResponse "Object that holds artist information"
// @Failure 400 {string} error "Error describing failure"
// @Router /artist [get]
func (h *Handler) handleArtist(inputURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SetCORSHeaders(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Get artist search from request
		searchString := r.URL.Query().Get("name")
		if searchString == "" {
			utils.WriteError(w, http.StatusBadRequest, errors.New("artist name not provided"))
			return
		}

		// Check if artist exists in db
		artist, err := h.Store.GetArtistByName(searchString)

		// If artist doesn't exist in db, search on setlist.fm
		if err != nil {
			url := fmt.Sprintf("%s/%s", inputURL, "search/artists")
			artist, err = setlist.ArtistSearch(url, searchString)

			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, err)
				return
			}

			// Use CreateArtistIfMissing to avoid duplicates
			artist = h.Store.CreateArtistIfMissing(*artist)
		}

		// Import setlist data (similar to handleArtistImport)
		mbid := artist.MBID

		// Retrieve or import artist setlists
		jsonData, err := utils.GetArtistSetlistsFromAPI(w, inputURL, mbid, 1)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		// Process artist info
		setlist.ProcessArtistInfo(h.Store, *jsonData, artist)

		// Extract tours and setlists
		tourNames := make(map[string]bool)
		setlistDates := make([]string, 0)
		recentSetlists := make([]map[string]string, 0)

		// Variables for upcoming shows and top songs
		var upcomingShows []map[string]string
		var topSongs []map[string]string

		// Process all setlists to gather information
		for i := range jsonData.Setlist {
			// Add tour to unique tours list
			if jsonData.Setlist[i].Tour.Name != "" {
				tourNames[jsonData.Setlist[i].Tour.Name] = true
			}

			// Add to setlist dates
			setlistDates = append(setlistDates, jsonData.Setlist[i].EventDate)

			// Add to recent setlists
			if len(recentSetlists) < 20 {
				setlistInfo := map[string]string{
					"id":    jsonData.Setlist[i].ID,
					"date":  jsonData.Setlist[i].EventDate,
					"venue": jsonData.Setlist[i].Venue.Name,
					"city":  jsonData.Setlist[i].Venue.City.Name,
					"url":   jsonData.Setlist[i].URL,
				}
				recentSetlists = append(recentSetlists, setlistInfo)
			}
		}

		// Get artist URL for scraping additional data
		artistURL := ""
		if len(jsonData.Setlist) > 0 {
			artistURL = jsonData.Setlist[0].Artist.URL
		}

		// Get upcoming shows and stats from artist's URL
		if artistURL != "" {
			artistData, err := utils.GetArtistDataFromAPI(artistURL)
			if err == nil {
				// Convert upcoming shows data to our format
				if upcomingData, ok := artistData["upcoming"].([]map[string]string); ok {
					upcomingShows = upcomingData
				}

				// Convert stats (top songs) data to our format
				if statsData, ok := artistData["stats"].([]map[string]string); ok {
					topSongs = statsData
				}
			} else {
				fmt.Println("Error retrieving artist data:", err)
			}
		}

		// Sort recent setlists by date (newest first)
		sort.Slice(recentSetlists, func(i, j int) bool {
			dateI, _ := time.Parse("02-01-2006", recentSetlists[i]["date"])
			dateJ, _ := time.Parse("02-01-2006", recentSetlists[j]["date"])
			return dateI.After(dateJ)
		})

		// Limit to 20 most recent
		if len(recentSetlists) > 20 {
			recentSetlists = recentSetlists[:20]
		}

		// Convert tourNames to slice
		tours := make([]string, 0, len(tourNames))
		for tourName := range tourNames {
			tours = append(tours, tourName)
		}

		// Create enhanced artist response
		enhancedResponse := map[string]interface{}{
			"artist":          artist,
			"artist_url":      artistURL,
			"number_of_tours": h.Store.GetTourTotalByArtist(artist.ID),
			"tour_names":      tours,
			"total_setlists":  h.Store.GetConcertTotalByArtist(artist.ID),
			"recent_setlists": recentSetlists,
			"upcoming_shows":  upcomingShows,
			"top_songs":       topSongs,
		}

		utils.WriteJSON(w, http.StatusOK, enhancedResponse)
	}
}

// @Summary Import information for a given artist into database
// @Description Gets setlist information from setlist.fm API for given artist, and imports it into database
// @Tags Artist
// @Param mbid path string true "Artist MBID"
// @Produce json
// @Success 201 {string} string "Message indicating success"
// @Failure 400 {string} error "Error describing failure"
// @Router /import [get]
func (h *Handler) handleArtistImport(inputURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SetCORSHeaders(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// Get artist search from request
		mbid := r.URL.Query().Get("mbid")
		if mbid == "" {
			utils.WriteError(w, http.StatusBadRequest, errors.New("artist mbid not provided"))
			return
		}
		fullImport := r.URL.Query().Get("full")
		if !(fullImport == "true" || fullImport == "") {
			utils.WriteError(w, http.StatusBadRequest, errors.New("invalid option for full parameter"))
			return
		}

		artist, err := h.Store.GetArtistByMBID(mbid)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, errors.New("artist mbid not in database"))
			return
		}
		jsonData, err := utils.GetArtistSetlistsFromAPI(w, inputURL, mbid, 1)
		if err != nil {
			return
		}
		setlist.ProcessArtistInfo(h.Store, *jsonData, artist)
		numPages := 1
		if fullImport != "" {
			numPages = int(math.Ceil(float64(jsonData.Total) / float64(jsonData.ItemsPerPage)))
		}

		i := 2
		for range time.Tick(1 * time.Second) {
			if i > numPages {
				break
			}
			jsonData, _ = utils.GetArtistSetlistsFromAPI(w, inputURL, mbid, i)
			setlist.ProcessArtistInfo(h.Store, *jsonData, artist)
			i++
		}

		utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "artist information successfully imported"})
	}
}

// @Summary Get concert setlist information
// @Description Returns details about a concert including the list of songs performed
// @Tags Concert
// @Param id path string true "Setlist ID"
// @Produce json
// @Success 200 {object} types.ConcertResponse "Concert setlist information"
// @Failure 400 {string} error "Error describing failure"
// @Router /concert [get]
func (h *Handler) handleConcert(inputURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SetCORSHeaders(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Get setlist ID from request
		setlistID := r.URL.Query().Get("id")
		if setlistID == "" {
			utils.WriteError(w, http.StatusBadRequest, errors.New("setlist ID not provided"))
			return
		}

		// Get setlist from setlist.fm API
		url := fmt.Sprintf("%s/setlist/%s", inputURL, setlistID)
		setlistData, err := setlist.GetSetlist(url)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		// Get or create artist record
		artist := h.Store.CreateArtistIfMissing(types.Artist{
			MBID: setlistData.Artist.Mbid,
			Name: setlistData.Artist.Name,
		})

		// Get or create venue record
		venue := h.Store.CreateVenueIfMissing(types.Venue{
			Name:       setlistData.Venue.Name,
			City:       setlistData.Venue.City.Name,
			Country:    setlistData.Venue.City.Country.Name,
			ExternalID: setlistData.Venue.ID,
			URL:        setlistData.Venue.URL,
		})

		// Parse event date
		eventDate, err := time.Parse("02-01-2006", setlistData.EventDate)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error parsing event date: %v", err))
			return
		}

		// Get or create tour if exists
		var tour *types.Tour
		if setlistData.Tour.Name != "" {
			tourRecord := types.Tour{
				ArtistID: artist.ID,
				Name:     setlistData.Tour.Name,
			}
			tour = h.Store.CreateTourIfMissing(tourRecord)
		}

		// Create concert record
		var tourID *uint
		if tour != nil {
			tourID = &tour.ID
		}

		concert := h.Store.CreateConcertIfMissing(types.Concert{
			ArtistID:          artist.ID,
			TourID:            tourID,
			VenueID:           venue.ID,
			Date:              eventDate,
			ExternalID:        setlistData.ID,
			ExternalVersionID: setlistData.VersionID,
		})

		// Process songs
		songsList := make([]map[string]interface{}, 0)
		songOrder := uint(1)

		for _, set := range setlistData.Sets.Set {
			for _, songData := range set.Song {
				// Process the main song
				song := types.Song{
					ArtistID: artist.ID,
					Name:     songData.Name,
					Info:     songData.Info,
					Tape:     songData.Tape,
				}

				// Process "with" artist if present
				if songData.With.Mbid != "" {
					withArtist := h.Store.CreateArtistIfMissing(types.Artist{
						MBID: songData.With.Mbid,
						Name: songData.With.Name,
					})
					withID := withArtist.ID
					song.WithID = &withID
				}

				// Process "cover" artist if present
				if songData.Cover.Mbid != "" {
					coverArtist := h.Store.CreateArtistIfMissing(types.Artist{
						MBID: songData.Cover.Mbid,
						Name: songData.Cover.Name,
					})
					coverID := coverArtist.ID
					song.CoverID = &coverID
				}

				// Save song to database
				songRecord := h.Store.CreateSongIfMissing(song)

				// Connect song to concert
				concertSong := types.ConcertSong{
					ConcertID: concert.ID,
					SongID:    songRecord.ID,
					SongOrder: songOrder,
				}
				h.Store.CreateConcertSongIfMissing(concertSong)

				// Prepare song data for response
				songInfo := map[string]interface{}{
					"name":  songData.Name,
					"info":  songData.Info,
					"tape":  songData.Tape,
					"order": songOrder,
				}

				// Add "with" information if present
				if songData.With.Mbid != "" {
					songInfo["with"] = map[string]string{
						"mbid": songData.With.Mbid,
						"name": songData.With.Name,
					}
				}

				// Add "cover" information if present
				if songData.Cover.Mbid != "" {
					songInfo["cover"] = map[string]string{
						"mbid": songData.Cover.Mbid,
						"name": songData.Cover.Name,
					}
				}

				songsList = append(songsList, songInfo)
				songOrder++
			}
		}

		// Create response
		response := map[string]interface{}{
			"id":           setlistData.ID,
			"version_id":   setlistData.VersionID,
			"event_date":   setlistData.EventDate,
			"last_updated": setlistData.LastUpdated,
			"artist": map[string]string{
				"mbid": setlistData.Artist.Mbid,
				"name": setlistData.Artist.Name,
				"url":  setlistData.Artist.URL,
			},
			"venue": map[string]interface{}{
				"id":   setlistData.Venue.ID,
				"name": setlistData.Venue.Name,
				"city": map[string]string{
					"name":    setlistData.Venue.City.Name,
					"state":   setlistData.Venue.City.State,
					"country": setlistData.Venue.City.Country.Name,
				},
				"url": setlistData.Venue.URL,
			},
			"songs": songsList,
			"url":   setlistData.URL,
		}

		// Add tour information if available
		if setlistData.Tour.Name != "" {
			response["tour"] = map[string]string{
				"name": setlistData.Tour.Name,
			}
		}

		// Add info if available
		if setlistData.Info != "" {
			response["info"] = setlistData.Info
		}

		utils.WriteJSON(w, http.StatusOK, response)
	}
}

func (h *Handler) handleUserPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SetCORSHeaders(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPost {
			h.UserPostOnPost(w, r)
		}

		if r.Method == http.MethodGet {
			h.UserPostOnGet(w, r)
		}
	}
}

func (h *Handler) handleUserLike() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SetCORSHeaders(w)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPost {
			h.UserLikeOnPost(w, r)
			return
		}

		if r.Method == http.MethodGet {
			h.UserLikeOnGet(w, r)
			return
		}
	}
}

func (h *Handler) handleUserFollow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.SetCORSHeaders(w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method == http.MethodPost {
			h.UserFollowOnPost(w, r)
			return
		}

		if r.Method == http.MethodGet {
			h.UserFollowOnGet(w, r)
			return
		}
	}
}
