package server

import (
	"car-tool/handlers"
	"github.com/gorilla/mux"
)

func (s *Server) SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/car", handlers.GetCarList).Methods("GET")
	return r
}
