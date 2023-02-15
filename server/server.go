package server

import (
	"car-tool/storage"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Port     string
	Database *storage.Database
}

func NewServer(db *storage.Database) (*Server, *error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, &err
	}
	s := Server{
		Port:     os.Getenv("SERVER_PORT"),
		Database: db,
	}
	return &s, nil
}

func (s *Server) StartServer(router *mux.Router) *error {
	err := http.ListenAndServe(s.Port, router)
	if err != nil {
		log.Panic(err.Error())
	}
	return &err
}
