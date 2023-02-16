package storage

import (
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	DB *sql.DB
}

var Db *Database

func NewDatabase() (*Database, *error) {
	db := Database{}
	err := godotenv.Load(".env")
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=aller password=alleristhebest dbname=car_tool sslmode=disable")
	//conn, err := sql.Open(os.Getenv("DATABASE_DRIVER"), os.Getenv("DSN"))
	if err != nil {
		log.Panic(err.Error())
	}
	Db = &db
	db.DB = conn
	return &db, nil
}
