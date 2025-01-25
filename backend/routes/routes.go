package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.handleHome).Methods("GET")
}

func (h *Handler) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"hello world"}`))
}
