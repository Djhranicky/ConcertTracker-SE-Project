package routes

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
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

	t.Run("should fail with no name query parameter", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/artist", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/artist", handler.handleArtist(""))

		router.ServeHTTP(rr, req)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should pass if artist already in database", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/artist?name=Artist1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/artist", handler.handleArtist(""))

		router.ServeHTTP(rr, req)

		var unmarshaled types.Artist
		json.Unmarshal(rr.Body.Bytes(), &unmarshaled)

		assertEqual(t, http.StatusOK, rr.Code)
		assertEqual(t, "mbid1", unmarshaled.MBID)
		assertEqual(t, "Artist1", unmarshaled.Name)
	})

	t.Run("should fail if artist not found in external API", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{
								"code": 404,
								"status": "Not Found",
								"message": "not found",
								"timestamp": "2025-03-22T22:19:29.358+0000"
							}`))
		}))
		defer server.Close()

		req, err := http.NewRequest(http.MethodGet, "/artist?name=NotFound", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/artist", handler.handleArtist(server.URL))

		router.ServeHTTP(rr, req)

		var unmarshaled types.Artist
		json.Unmarshal(rr.Body.Bytes(), &unmarshaled)

		assertEqual(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should pass if artist found in external API", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
							"type": "artists",
							"itemsPerPage": 30,
							"page": 1,
							"total": 1,
							"artist": [
								{
									"mbid": "mbid2",
									"name": "Artist2",
									"sortName": "Artist2",
									"disambiguation": "",
									"url": "https://www.setlist.fm/setlists/Artist2.html"
								}
							]
						}`))
		}))
		defer server.Close()

		req, err := http.NewRequest(http.MethodGet, "/artist?name=Artist2", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/artist", handler.handleArtist(server.URL))

		router.ServeHTTP(rr, req)

		var unmarshaled types.Artist
		json.Unmarshal(rr.Body.Bytes(), &unmarshaled)

		assertEqual(t, http.StatusOK, rr.Code)
	})

	t.Run("should pass if previously missing artist was found in external API", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/artist?name=Artist2", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/artist", handler.handleArtist(""))

		router.ServeHTTP(rr, req)

		var unmarshaled types.Artist
		json.Unmarshal(rr.Body.Bytes(), &unmarshaled)

		assertEqual(t, http.StatusOK, rr.Code)
		assertEqual(t, uint(2), unmarshaled.ID)
		assertEqual(t, "mbid2", unmarshaled.MBID)
		assertEqual(t, "Artist2", unmarshaled.Name)
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
	)
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("expected %v (type %v), received %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	}
}
