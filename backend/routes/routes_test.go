package routes

import (
	"bytes"
	"encoding/json"
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

func TestUserServiceHandleValidate(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	t.Run("should fail when no id cookie is present", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/validate", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/validate", handler.handleValidate)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should fail when invalid jwt string is present", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/validate", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(&http.Cookie{
			Name:  "id",
			Value: "invalid jwt token",
		})

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/validate", handler.handleValidate)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusUnauthorized {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should pass when valid cookie is present", func(t *testing.T) {
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

		cookie := rr.Result().Cookies()
		req, err = http.NewRequest(http.MethodGet, "/validate", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.AddCookie(cookie[0])

		rr = httptest.NewRecorder()
		router = mux.NewRouter()

		router.HandleFunc("/validate", handler.handleValidate)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status code %v, got status code %v", http.StatusBadRequest, rr.Code)
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
			Text:       &text,
			Type:       "WISHLIST",
			Rating:     &rating,
			UserPostID: &userPostID,
			IsPublic:   &isPublic,
			ConcertID:  1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

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
			AuthorID:   1,
			Text:       &text,
			Rating:     &rating,
			UserPostID: &userPostID,
			IsPublic:   &isPublic,
			ConcertID:  1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

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
			AuthorID:   1,
			Text:       &text,
			Type:       "WISHLIST",
			Rating:     &rating,
			UserPostID: &userPostID,
			ConcertID:  1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

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
			AuthorID:   1,
			Text:       &text,
			Type:       "WISHLIST",
			Rating:     &rating,
			UserPostID: &userPostID,
			IsPublic:   &isPublic,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

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
			AuthorID:   1,
			Text:       &text,
			Type:       "WRONG_TYPE",
			Rating:     &rating,
			UserPostID: &userPostID,
			IsPublic:   &isPublic,
			ConcertID:  1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

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
			AuthorID:   1,
			Text:       &text,
			Type:       "WISHLIST",
			Rating:     &rating,
			UserPostID: &userPostID,
			IsPublic:   &isPublic,
			ConcertID:  1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/userpost", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

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
			UserID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

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

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should succeed when user first likes a post", func(t *testing.T) {
		payload := &types.UserLikePostPayload{
			UserID:     1,
			UserPostID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("should succeed when user removes like from post", func(t *testing.T) {
		payload := &types.UserLikePostPayload{
			UserID:     1,
			UserPostID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/like", handler.handleUserLike())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("should succeed when user likes a post again", func(t *testing.T) {
		payload := &types.UserLikePostPayload{
			UserID:     1,
			UserPostID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/like", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

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
			UserID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should fail if UserID not included", func(t *testing.T) {
		payload := &types.UserFollowPayload{
			FollowedUserID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should succeed when user first follows another user", func(t *testing.T) {
		payload := &types.UserFollowPayload{
			UserID:         1,
			FollowedUserID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("should succeed when user unfollows another user", func(t *testing.T) {
		payload := &types.UserFollowPayload{
			UserID:         1,
			FollowedUserID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("should succeed when user follows another user again", func(t *testing.T) {
		payload := &types.UserFollowPayload{
			UserID:         1,
			FollowedUserID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/follow", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

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

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/follow", handler.handleUserFollow())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})
}

func TestUserServiceHandleList(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer destroyDatabase(database)

	t.Run("create should fail when userID not included", func(t *testing.T) {
		payload := &types.UserListCreatePayload{
			Name: "Test Name",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/listcreate", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/listcreate", handler.handleListCreate())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("create should fail when name not included", func(t *testing.T) {
		payload := &types.UserListCreatePayload{
			UserID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/listcreate", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/listcreate", handler.handleListCreate())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("create should pass when valid fields included", func(t *testing.T) {
		payload := &types.UserListCreatePayload{
			UserID: 1,
			Name:   "Test Name",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/listcreate", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/listcreate", handler.handleListCreate())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusCreated, rr.Code)
	})

	t.Run("add should fail when listID not included", func(t *testing.T) {
		payload := &types.UserListAddPayload{
			ConcertID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/listadd", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/listadd", handler.handleListAdd())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("add should fail when concertID not included", func(t *testing.T) {
		payload := &types.UserListAddPayload{
			ListID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/listadd", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/listadd", handler.handleListAdd())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("add should pass when adding a concert", func(t *testing.T) {
		payload := &types.UserListAddPayload{
			ListID:    1,
			ConcertID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/listadd", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/listadd", handler.handleListAdd())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("add should pass when removing a concert", func(t *testing.T) {
		payload := &types.UserListAddPayload{
			ListID:    1,
			ConcertID: 1,
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/listadd", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/listadd", handler.handleListAdd())

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusOK, rr.Code)
	})
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
		&types.List{},
		&types.ListConcert{},
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
		&types.List{},
		&types.ListConcert{},
	)
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("expected %v (type %v), received %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}
