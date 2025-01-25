package api

import (
	"net/http"

	"github.com/djhranicky/ConcertTracker-SE-Project/routes"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	handler := routes.NewHandler()
	handler.RegisterRoutes(router)

	http.Handle("/", router)
	return http.ListenAndServe(s.addr, router)
}
