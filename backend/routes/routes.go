package routes

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"os"
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
// @Failure 400 {string} string "missing or invalid authorization token"
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
// @Success 200 {object} types.Artist "Object that holds artist information"
// @Failure 400 {string} error "Error describing failure"
// @Router /artist [post]
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

		// If so, return info from there
		if err == nil {
			utils.WriteJSON(w, http.StatusOK, *artist)
			return
		}

		// If not, check if artist exists on setlist.fm
		url := fmt.Sprintf("%s/%s", inputURL, "search/artists")
		artist, err = setlist.ArtistSearch(url, searchString)

		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		// If so, create info in database and return info
		err = h.Store.CreateArtist(*artist)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		utils.WriteJSON(w, http.StatusOK, *artist)
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

		for i := 2; i <= numPages; i++ {
			time.Sleep(2 * time.Second)
			jsonData, _ = utils.GetArtistSetlistsFromAPI(w, inputURL, mbid, i)
			setlist.ProcessArtistInfo(h.Store, *jsonData, artist)
		}

		utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "artist information successfully imported"})
	}
}
