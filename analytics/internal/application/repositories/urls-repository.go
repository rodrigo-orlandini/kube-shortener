package repositories

import "rodrigoorlandini/urlshortener/analytics/internal/domain/entities"

type TopRankedURL struct {
	ShortURL   string `json:"shortUrl"`
	VisitCount int    `json:"visitCount"`
}

type URLsRepository interface {
	Create(url entities.URL) error
	GetTopRanked(limit int) ([]TopRankedURL, error)
	FindByShortURL(shortURL string) (entities.URL, error)
	IncrementVisits(shortURL string) error
}
