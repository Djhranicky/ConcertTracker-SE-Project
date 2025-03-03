package routes

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/djhranicky/ConcertTracker-SE-Project/db"
	"github.com/djhranicky/ConcertTracker-SE-Project/service/auth"
	"github.com/djhranicky/ConcertTracker-SE-Project/service/user"
	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/djhranicky/ConcertTracker-SE-Project/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func TestUserServiceHandleRegister(t *testing.T) {
	utils.Init()
	handler, database := initTestHandler()
	defer database.Migrator().DropTable(&types.User{})

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
	defer database.Migrator().DropTable(&types.User{})

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

func initTestDatabase(dbName string) *gorm.DB {
	mockDatabase, err := db.NewSqliteStorage(dbName)
	if err != nil {
		log.Fatal(err)
	}

	mockDatabase.AutoMigrate(&types.User{})

	return mockDatabase
}

func initTestHandler() (*Handler, *gorm.DB) {
	database := initTestDatabase("test.db")
	userStore := user.NewStore(database)
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

	database.Create(&user)

	return handler, database
}

type MockUserStore struct {
	db *gorm.DB
}

func (m *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *MockUserStore) GetUserByID(id uint) (*types.User, error) {
	return nil, nil
}

func (m *MockUserStore) CreateUser(types.User) error {
	return nil
}
