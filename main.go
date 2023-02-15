package main

import (
	"car-tool/server"
	"car-tool/storage"
	"log"
)

func main() {
	db, err := storage.NewDatabase()
	if err != nil {
		log.Panicln(err)
	}
	s, err := server.NewServer(db)
	if err != nil {
		log.Panicln("Can't create server")
	}
	err = s.StartServer(s.SetupRouter())
	if err != nil {
		log.Panicln("Can't start server")
	}
}
