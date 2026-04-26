package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	connStr := fmt.Sprintf(
		"postgres://postgres:postgres@%s:5432/urlshort?sslmode=disable",
		host,
	)

	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("✅ Connected to DB")
				return db
			}
		}

		fmt.Println("⏳ Waiting for DB...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("❌ Could not connect to DB:", err)
	return nil
}
