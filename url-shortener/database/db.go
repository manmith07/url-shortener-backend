package database

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "postgresql://neondb_owner:npg_IJTOGxBA4o8R@ep-polished-wave-and6uqxt-pooler.c-6.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

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
