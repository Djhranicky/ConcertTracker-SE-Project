package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/gorilla/mux"
)

// TODO: fix these tests

func TestUserServiceHandlers(t *testing.T) {
	userStore := &MockUserStore{}
	handler := NewHandler(userStore)

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
}

type MockUserStore struct {
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
