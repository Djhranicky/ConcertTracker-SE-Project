package routes

import (
	"net/http"

	"github.com/djhranicky/ConcertTracker-SE-Project/types"
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
}

func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"hello world"}`))
}
