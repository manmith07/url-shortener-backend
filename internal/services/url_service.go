package services

import (
	"math/rand"
	"time"

	"github.com/manmithsm/url-shortener/internal/repository"
)

type URLService struct {
	Repo *repository.URLRepository
}

func NewURLService(r *repository.URLRepository) *URLService {
	return &URLService{Repo: r}
}

func generateCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	rand.Seed(time.Now().UnixNano())

	code := make([]byte, 6)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

func (s *URLService) CreateShortURL(original string) (string, error) {
	code := generateCode()
	err := s.Repo.Save(code, original)
	return code, err
}

func (s *URLService) GetOriginalURL(code string) (string, error) {
	return s.Repo.Get(code)
}
