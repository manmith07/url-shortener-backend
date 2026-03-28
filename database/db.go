package database

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "user=manmithsm dbname=urlshortener sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to DB")
}
