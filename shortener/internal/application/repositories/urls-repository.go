package repositories

import "rodrigoorlandini/urlshortener/shortener/internal/domain/entities"

type URLsRepository interface {
	Create(url entities.URL) error
	FindByOriginalURL(originalURL string) (entities.URL, error)
	FindByShortURL(shortURL string) (entities.URL, error)
	Visit(shortURL string) error
}
