package repository

import "database/sql"

type URLRepository struct {
	DB *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{DB: db}
}

func (r *URLRepository) Save(short, original string) error {
	query := "INSERT INTO urls (short_code, original_url) VALUES ($1, $2)"
	_, err := r.DB.Exec(query, short, original)
	return err
}

func (r *URLRepository) Get(short string) (string, error) {
	var original string
	query := "SELECT original_url FROM urls WHERE short_code=$1"
	err := r.DB.QueryRow(query, short).Scan(&original)
	return original, err
}
