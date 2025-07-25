package entities

import (
	customError "rodrigoorlandini/urlshortener/analytics/internal/domain/custom-error"

	"github.com/google/uuid"
)

type URL struct {
	ID       *string
	ShortURL string
	Visits   int
}

type URLOption func(*URL)

func NewURL(shortURL string, options ...URLOption) (*URL, error) {
	if shortURL == "" {
		return nil, &customError.InvalidEntityCreationError{
			EntityName: "URL",
			Field:      "shortURL",
			Value:      shortURL,
		}
	}

	url := &URL{
		ShortURL: shortURL,
		Visits:   0,
	}

	for _, option := range options {
		option(url)
	}

	if url.ID == nil {
		id := uuid.New().String()
		url.ID = &id
	}

	return url, nil
}

func WithID(id string) URLOption {
	return func(url *URL) {
		url.ID = &id
	}
}

func WithVisits(visits int) URLOption {
	return func(url *URL) {
		url.Visits = visits
	}
}
