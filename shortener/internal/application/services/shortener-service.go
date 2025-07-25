package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"rodrigoorlandini/urlshortener/shortener/internal/domain/entities"
)

type ShortenerService struct{}

func NewShortenerService() *ShortenerService {
	return &ShortenerService{}
}

func (s *ShortenerService) ShortenURL(url entities.URL) (entities.URL, error) {
	hash := sha256.Sum256([]byte(url.OriginalURL))
	shortCode := hex.EncodeToString(hash[:])[:8]

	shortURL := fmt.Sprintf("/%s", shortCode)
	url.ShortURL = shortURL

	return url, nil
}
