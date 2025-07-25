package entities

import (
	customError "rodrigoorlandini/urlshortener/shortener/internal/domain/custom-error"
	"time"
)

type Visit struct {
	ShortUrl  string    `json:"short_url"`
	VisitedAt time.Time `json:"visited_at"`
}

func NewVisit(shortUrl string, visitedAt time.Time) (*Visit, error) {
	if shortUrl == "" {
		return nil, &customError.InvalidEntityCreationError{
			EntityName: "Visit",
			Field:      "shortURL",
			Value:      shortUrl,
		}
	}

	visit := &Visit{
		ShortUrl:  shortUrl,
		VisitedAt: visitedAt,
	}

	return visit, nil
}
