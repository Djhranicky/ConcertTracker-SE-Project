package api

import (
	"fmt"
	"net/http"

	"github.com/djhranicky/ConcertTracker-SE-Project/db"
	"github.com/djhranicky/ConcertTracker-SE-Project/routes"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()

	userStore := db.NewStore(s.db)
	handler := routes.NewHandler(userStore)
	handler.RegisterRoutes(subrouter)

	fmt.Println("Listening on port", s.addr)

	http.Handle("/", subrouter)
	return http.ListenAndServe(s.addr, router)
}
