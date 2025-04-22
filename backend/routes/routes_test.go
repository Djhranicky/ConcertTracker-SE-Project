package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/djhranicky/ConcertTracker-SE-Project/db"
	"github.com/djhranicky/ConcertTracker-SE-Project/service/auth"
	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/djhranicky/ConcertTracker-SE-Project/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func TestUserServiceHandleRegister(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	t.Run("Should fail if request body is empty", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/register", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.UserRegisterPayload{
			Name:     "testuser",
			Email:    "test", // Bad email field
			Password: "test123",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should fail if user exists", func(t *testing.T) {
		payload := types.UserRegisterPayload{
			Name:     "John Doe",
			Email:    "test@example.com",
			Password: "testpw",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should succeed when new user is created", func(t *testing.T) {
		payload := types.UserRegisterPayload{
			Name:     "Created User",
			Email:    "test2@example.com",
			Username: "createduser",
			Password: "testpw",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		log.Print(rr.Body)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %v, got status code %v", http.StatusCreated, rr.Code)
		}
	})
}

func TestUserServiceHandleLogin(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	t.Run("Should fail if request body is empty", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, "/login", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.handleLogin)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should fail if payload is invalid", func(t *testing.T) {
		payload := &types.UserRegisterPayload{
			Name:     "John Doe",
			Email:    "test", //Bad Email
			Password: "test123",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.handleLogin)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should fail if user does not exist", func(t *testing.T) {
		payload := &types.UserRegisterPayload{
			Name:     "John Doe",
			Email:    "doesnotexist@example.com",
			Password: "test123",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.handleLogin)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should fail if user enters wrong password", func(t *testing.T) {
		payload := &types.UserRegisterPayload{
			Name:     "John Doe",
			Email:    "test@example.com",
			Password: "wrongpassword",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.handleLogin)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("Should pass if user enters correct user name and password", func(t *testing.T) {
		payload := &types.UserRegisterPayload{
			Name:     "John Doe",
			Email:    "test@example.com",
			Password: "test",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/login", handler.handleLogin)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %v, got status code %v. JSON Body: %v", http.StatusOK, rr.Code, rr.Body)
		}
	})
}

func TestArtistServiceHandleArtist(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	// Create test data in database
	artist1 := types.Artist{
		MBID: "mbid1",
		Name: "Artist1",
	}
	database.Create(&artist1)

	// Mock setlist.fm API responses
	mockSetlistServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "search/artists") {
			// Artist search response
			w.Write([]byte(`{
				"type": "artists",
				"artist": [{
					"MBID": "mbid3", 
					"name": "Artist3"
				}]
			}`))
		} else if strings.Contains(r.URL.Path, "setlists") {
			// Setlist response
			w.Write([]byte(`{
				"setlist": [{
					"artist": {"MBID": "mbid3", "name": "Artist3"},
					"venue": {"name": "Venue1"},
					"eventDate": "01-01-2025"
				}]
			}`))
		}
	}))
	defer mockSetlistServer.Close()

	t.Run("should pass if artist already in database", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/artist?name=Artist1", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/artist", handler.handleArtist(mockSetlistServer.URL))
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rr.Code)
			t.Logf("response: %s", rr.Body.String())
			return
		}

		var response map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("failed to parse response: %v", err)
		}

		artist := response["artist"].(map[string]interface{})
		if artist["MBID"] != "mbid1" {
			t.Errorf("expected mbid1, got %v", artist["mbid"])
		}
	})

	t.Run("should pass if artist found in external API", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/artist?name=Artist3", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/artist", handler.handleArtist(mockSetlistServer.URL))
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status 200, got %d", rr.Code)
			t.Logf("response: %s", rr.Body.String())
			return
		}

		var response map[string]interface{}
		if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
			t.Fatalf("failed to parse response: %v", err)
		}

		artist := response["artist"].(map[string]interface{})
		if artist["MBID"] != "mbid3" {
			t.Errorf("expected mbid3, got %v", artist["mbid"])
		}
	})
}

func TestArtistServiceHandleImport(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	t.Run("should fail with no name query parameter", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/import", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/import", handler.handleArtistImport(""))

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail with invalid full query parameter", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/import?name=test&full=Failure", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/import", handler.handleArtistImport(""))

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail if artist mbid not in database", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/import?mbid=mbidNotFound", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/import", handler.handleArtistImport(""))

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail if artist mbid not in external API", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{
						"code": 404,
						"status": "Not Found",
						"message": "not found",
						"timestamp": "2025-03-29T00:29:24.574+0000"
					}`))
		}))
		defer server.Close()
		req, err := http.NewRequest(http.MethodGet, "/import?mbid=Artist1", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/import", handler.handleArtistImport(""))

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should pass if artist mbid in database", func(t *testing.T) {
		data, err := os.ReadFile("./routes/testdata/mock_ArtistMBIDSetlist.json")
		if err != nil {
			t.Fatal(err)
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}))
		defer server.Close()

		req, err := http.NewRequest(http.MethodGet, "/import?mbid=mbid1", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/import", handler.handleArtistImport(server.URL))

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusCreated, rr.Code)
	})

	t.Run("should pass if artist mbid in database for full import", func(t *testing.T) {
		data, err := os.ReadFile("./routes/testdata/mock_ArtistMBIDSetlist.json")
		if err != nil {
			t.Fatal(err)
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		}))
		defer server.Close()

		req, err := http.NewRequest(http.MethodGet, "/import?mbid=mbid1&full=true", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/import", handler.handleArtistImport(server.URL))

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusCreated, rr.Code)
	})
}

func TestConcertServiceHandleConcert(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	t.Run("should fail with no id query parameter", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/concert", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/concert", handler.handleConcert(""))

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail if setlist not found in external API", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{
				"code": 404,
				"status": "Not Found",
				"message": "not found",
				"timestamp": "2025-03-29T00:29:24.574+0000"
			}`))
		}))
		defer server.Close()

		req, err := http.NewRequest(http.MethodGet, "/concert?id=notfound", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/concert", handler.handleConcert(server.URL))

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})
}

func TestUserServiceHandlePost(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	t.Run("should fail with no 'authorID' field included in incoming json", func(t *testing.T) {
		text := "Test"
		rating := uint(1)
		userPostID := uint(1)
		isPublic := true
		payload := &types.UserPostCreatePayload{
			Text:              &text,
			Type:              "WISHLIST",
			Rating:            &rating,
			UserPostID:        &userPostID,
			IsPublic:          &isPublic,
			ExternalConcertID: "test",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userpost", handler.handleUserPost())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail with no 'type' field included in incoming json", func(t *testing.T) {
		text := "Test"
		rating := uint(1)
		userPostID := uint(1)
		isPublic := true
		payload := &types.UserPostCreatePayload{
			AuthorUsername:    "johndoe",
			Text:              &text,
			Rating:            &rating,
			UserPostID:        &userPostID,
			IsPublic:          &isPublic,
			ExternalConcertID: "test",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userpost", handler.handleUserPost())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail with no 'isPublic' field included in incoming json", func(t *testing.T) {
		text := "Test"
		rating := uint(1)
		userPostID := uint(1)
		payload := &types.UserPostCreatePayload{
			AuthorUsername:    "johndoe",
			Text:              &text,
			Type:              "WISHLIST",
			Rating:            &rating,
			UserPostID:        &userPostID,
			ExternalConcertID: "test",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userpost", handler.handleUserPost())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail with no 'concertID' field included in incoming json", func(t *testing.T) {
		text := "Test"
		rating := uint(1)
		userPostID := uint(1)
		isPublic := true
		payload := &types.UserPostCreatePayload{
			AuthorUsername: "johndoe",
			Text:           &text,
			Type:           "WISHLIST",
			Rating:         &rating,
			UserPostID:     &userPostID,
			IsPublic:       &isPublic,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userpost", handler.handleUserPost())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail if invalid 'type' supplied", func(t *testing.T) {
		text := "Test"
		rating := uint(1)
		userPostID := uint(1)
		isPublic := true
		payload := &types.UserPostCreatePayload{
			AuthorUsername:    "johndoe",
			Text:              &text,
			Type:              "WRONG_TYPE",
			Rating:            &rating,
			UserPostID:        &userPostID,
			IsPublic:          &isPublic,
			ExternalConcertID: "test",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userpost", handler.handleUserPost())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should pass and create post in database", func(t *testing.T) {
		text := "Test"
		rating := uint(1)
		userPostID := uint(1)
		isPublic := true
		payload := &types.UserPostCreatePayload{
			AuthorUsername:    "johndoe",
			Text:              &text,
			Type:              "WISHLIST",
			Rating:            &rating,
			UserPostID:        &userPostID,
			IsPublic:          &isPublic,
			ExternalConcertID: "test",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userpost", handler.handleUserPost())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusCreated, rr.Code)
	})

	t.Run("GET should fail with no userID query parameter", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/userpost", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userpost", handler.handleUserPost())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("GET should fail with bad userID query parameter", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/userpost?userID=test", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userpost", handler.handleUserPost())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("GET should pass with valid userID query parameter", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/userpost?userID=1", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userpost", handler.handleUserPost())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})
}

func TestUserServiceHandleLike(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	t.Run("should fail if UserPostID not included", func(t *testing.T) {
		payload := &types.UserLikePostPayload{
			Username: "johndoe",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail if UserID not included", func(t *testing.T) {
		payload := &types.UserLikePostPayload{
			UserPostID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should succeed when user first likes a post", func(t *testing.T) {
		payload := &types.UserLikePostPayload{
			Username:   "johndoe",
			UserPostID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("should succeed when user removes like from post", func(t *testing.T) {
		payload := &types.UserLikePostPayload{
			Username:   "johndoe",
			UserPostID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("should succeed when user likes a post again", func(t *testing.T) {
		payload := &types.UserLikePostPayload{
			Username:   "johndoe",
			UserPostID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("GET should fail when query param not provided", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/like", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("GET should fail when bad query param provided", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/like?userPostID=test", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("GET should return like count for valid input", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/like?userPostID=1", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})
}

func TestUserServiceHandleFollow(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	t.Run("should fail if FollowedUserID not included", func(t *testing.T) {
		payload := &types.UserFollowPayload{
			Username: "johndoe",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail if UserID not included", func(t *testing.T) {
		payload := &types.UserFollowPayload{
			FollowedUsername: "janedoe",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should succeed when user first follows another user", func(t *testing.T) {
		payload := &types.UserFollowPayload{
			Username:         "johndoe",
			FollowedUsername: "janedoe",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("should succeed when user unfollows another user", func(t *testing.T) {
		payload := &types.UserFollowPayload{
			Username:         "johndoe",
			FollowedUsername: "janedoe",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("should succeed when user follows another user again", func(t *testing.T) {
		payload := &types.UserFollowPayload{
			Username:         "johndoe",
			FollowedUsername: "janedoe",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("GET should fail if userID param not included", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/follow?type=followers", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("GET should fail if type param not included", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/follow?userID=1", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("GET should fail if type param invalid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/follow?userID=1&type=blah", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("GET should fail bad userID provided", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/follow?userID=test&type=followers", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("GET should pass with valid userID and type = followers", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/follow?userID=1&type=followers", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("GET should pass with valid userID and type = following", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/follow?userID=1&type=following", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})
}

func TestSessionMethods(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("cannot load env file")
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	t.Run("should fail if request has no cookie", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		_, err = auth.GetJWTCookie(req)
		if err != http.ErrNoCookie {
			t.Errorf("expected error code %v, got nothing", http.ErrNoCookie)
		}
	})

	t.Run("should pass if request has cookie", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:     "id",
			Value:    "",
			HttpOnly: true,
		})
		cookie, err := auth.GetJWTCookie(req)
		if err != nil {
			t.Errorf("expected cookie in request, got %v", cookie)
		}
	})

	t.Run("verification should fail if no cookie present", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		err = auth.ValidateUser(req, handler.Store)
		if err != http.ErrNoCookie {
			t.Errorf("expected error code %v, got %v", http.ErrNoCookie, err)
		}
	})

	t.Run("verification should fail if no JWT token present", func(t *testing.T) {
		expectedErr := fmt.Errorf("missing authorization token")
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:     "id",
			Value:    "",
			HttpOnly: true,
		})
		req.Header.Add("username", "johndoe")
		err = auth.ValidateUser(req, handler.Store)
		if err.Error() != expectedErr.Error() {
			t.Errorf("expected error code %v, got %v", expectedErr, err)
		}
	})

	t.Run("verification should fail if JWT token is expired", func(t *testing.T) {
		expectedErr := fmt.Errorf("token has invalid claims: token is expired")
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		token, _ := auth.CreateJWT(secret, 1, -1)
		req.AddCookie(&http.Cookie{
			Name:     "id",
			Value:    token,
			HttpOnly: true,
		})
		req.Header.Add("username", "johndoe")
		err = auth.ValidateUser(req, handler.Store)
		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("expected error code %v, got %v", expectedErr, err)
		}
	})

	t.Run("verification should succeed if JWT token is valid", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		token, _ := auth.CreateJWT(secret, 1, 10)
		req.AddCookie(&http.Cookie{
			Name:     "id",
			Value:    token,
			HttpOnly: true,
		})
		req.Header.Add("username", "johndoe")
		err = auth.ValidateUser(req, handler.Store)
		if err != nil {
			t.Errorf("expected no error code, got %v", err)
		}
	})
}

func TestUserInfoRoute(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	// Create a test user with unique credentials
	uniqueEmail := "unique_test_user@example.com"
	uniqueUsername := "unique_testuser"
	testUser := types.User{
		Name:     "Unique Test User",
		Username: uniqueUsername,
		Email:    uniqueEmail,
		Password: "hashedpassword",
	}
	_ = handler.Store.CreateUser(testUser)

	t.Run("should fail without authorization", func(t *testing.T) {
		payload := map[string]string{
			"username": uniqueUsername,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodGet, "/userinfo", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		// No auth token added

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userinfo", handler.handleUserInfo)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should fail with no username in payload", func(t *testing.T) {
		payload := map[string]string{}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodGet, "/userinfo", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userinfo", handler.handleUserInfo)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail with empty username", func(t *testing.T) {
		payload := map[string]string{
			"username": "",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodGet, "/userinfo", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userinfo", handler.handleUserInfo)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail with non-existent username", func(t *testing.T) {
		payload := map[string]string{
			"username": "nonexistentuser_123456",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodGet, "/userinfo", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userinfo", handler.handleUserInfo)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusNotFound, rr.Code)
	})

	t.Run("should succeed with valid username", func(t *testing.T) {
		payload := map[string]string{
			"username": uniqueUsername,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodGet, "/userinfo", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userinfo", handler.handleUserInfo)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)

		// Check response contains expected fields
		var response map[string]string
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response JSON: %v", err)
		}

		if response["name"] != "Unique Test User" {
			t.Errorf("Expected name 'Unique Test User', got '%s'", response["name"])
		}
		if response["email"] != uniqueEmail {
			t.Errorf("Expected email '%s', got '%s'", uniqueEmail, response["email"])
		}
	})

	t.Run("should handle OPTIONS request correctly", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodOptions, "/userinfo", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/userinfo", handler.handleUserInfo)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})
}

func TestUsersRoute(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	// First, get the current count of users to use as a baseline
	var initialUserCount int

	// Get all users first to know what we're starting with
	existingUsers, _ := handler.Store.GetAllUsers()
	initialUserCount = len(existingUsers)

	t.Run("should fail without authorization", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}
		// No auth token added

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users", handler.handleUserList)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("should return current users list", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users", handler.handleUserList)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response JSON: %v", err)
		}

		usernames, ok := response["usernames"].([]interface{})
		if !ok {
			t.Fatalf("Expected usernames field to be array, got %T", response["usernames"])
		}

		count, ok := response["count"].(float64)
		if !ok {
			t.Fatalf("Expected count field to be number, got %T", response["count"])
		}

		if int(count) != len(usernames) {
			t.Errorf("Count %v doesn't match actual number of usernames %v", count, len(usernames))
		}

		if int(count) != initialUserCount {
			t.Errorf("Expected %d users, got %v", initialUserCount, count)
		}
	})

	t.Run("should return updated list after adding new users", func(t *testing.T) {
		// Create test users with unique credentials
		testUsers := []types.User{
			{
				Name:     "Test User One",
				Username: "test_user1",
				Email:    "test_user1@example.com",
				Password: "hashedpassword1",
			},
			{
				Name:     "Test User Two",
				Username: "test_user2",
				Email:    "test_user2@example.com",
				Password: "hashedpassword2",
			},
		}

		for _, user := range testUsers {
			err := handler.Store.CreateUser(user)
			if err != nil {
				t.Fatalf("Failed to create test user: %v", err)
			}
		}

		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}
		addDefaultValidation(req)

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users", handler.handleUserList)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		if err != nil {
			t.Fatalf("Failed to parse response JSON: %v", err)
		}

		usernames, ok := response["usernames"].([]interface{})
		if !ok {
			t.Fatalf("Expected usernames field to be array, got %T", response["usernames"])
		}

		count, ok := response["count"].(float64)
		if !ok {
			t.Fatalf("Expected count field to be number, got %T", response["count"])
		}

		expectedCount := initialUserCount + len(testUsers)
		if int(count) != expectedCount {
			t.Errorf("Expected %d users, got %v", expectedCount, count)
		}

		// Check if all our new test users are in the response
		foundUsers := make(map[string]bool)
		for _, username := range usernames {
			foundUsers[username.(string)] = true
		}

		for _, user := range testUsers {
			if !foundUsers[user.Name] {
				t.Errorf("Expected user %s in response, but not found", user.Name)
			}
		}
	})

	t.Run("should handle OPTIONS request correctly", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodOptions, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/users", handler.handleUserList)

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})
}

func addDefaultValidation(req *http.Request) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("cannot load env file")
	}
	secret := []byte(os.Getenv("JWT_SECRET"))
	token, _ := auth.CreateJWT(secret, 1, 10)
	req.AddCookie(&http.Cookie{
		Name:     "id",
		Value:    token,
		HttpOnly: true,
	})
	req.Header.Add("username", "johndoe")
}

func initTestDatabase(dbName string) *gorm.DB {
	mockDatabase, err := db.NewSqliteStorage(dbName)
	if err != nil {
		log.Fatal(err)
	}

	mockDatabase.AutoMigrate(
		&types.User{},
		&types.Artist{},
		&types.Tour{},
		&types.Venue{},
		&types.Concert{},
		&types.Song{},
		&types.ConcertSong{},
		&types.UserPost{},
		&types.Likes{},
		&types.Follow{},
	)

	return mockDatabase
}

func initTestHandler() (*Handler, *gorm.DB) {
	database := initTestDatabase("test.db")
	userStore := db.NewStore(database)
	handler := NewHandler(userStore)

	hashedPassword, err := auth.HashPassword("test")
	if err != nil {
		log.Fatal(err)
	}
	user := types.User{
		Name:     "John Doe",
		Email:    "test@example.com",
		Username: "johndoe",
		Password: hashedPassword,
	}
	artist := types.Artist{
		MBID: "mbid1",
		Name: "Artist1",
	}

	database.Create(&user)
	database.Create(&artist)

	return handler, database
}

func destroyDatabase(database *gorm.DB) {
	database.Migrator().DropTable(
		&types.User{},
		&types.Artist{},
		&types.Tour{},
		&types.Venue{},
		&types.Concert{},
		&types.Song{},
		&types.ConcertSong{},
		&types.UserPost{},
		&types.Likes{},
		&types.Follow{},
	)
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("expected %v (type %v), received %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}
