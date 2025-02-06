package routes

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/gorilla/mux"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &MockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Should fail if the user payload is invalid on register", func(t *testing.T) {
		payload := types.UserRegisterPayload{
			Name:     "testuser",
			Email:    "test@example.com",
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

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %v, got status code %v", http.StatusCreated, rr.Code)
		}
	})
}

type MockUserStore struct {
}

func (m *MockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, errors.New("user not found")
}

func (m *MockUserStore) GetUserByID(id uint) (*types.User, error) {
	return nil, errors.New("user not found")
}

func (m *MockUserStore) CreateUser(user types.User) error {
	return nil
}
