package server

import (
	"car-tool/handlers"
	"github.com/gorilla/mux"
)

func (s *Server) SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Hello)
	return r
}
