package routes

import (
	"fmt"
	"net/http"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
	"github.com/djhranicky/ConcertTracker-SE-Project/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	UserStore types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{UserStore: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.handleHome).Methods("GET")
	router.HandleFunc("/login", h.handleLogin).Methods("GET")
}

func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"hello world"}`))
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.UserLoginPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	_, err := h.UserStore.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s does not exist", payload.Email))
		return
	}

	// w.Write([]byte(`{"message":"Log in successful"}`))
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("User %s logged in successfully", payload.Email)})
}
