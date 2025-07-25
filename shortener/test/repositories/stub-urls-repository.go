package stubRepositories

import (
	"errors"
	"rodrigoorlandini/urlshortener/shortener/internal/application/repositories"
	"rodrigoorlandini/urlshortener/shortener/internal/domain/entities"
	"time"
)

var INSTANCE *StubURLsRepository

type StubURLsRepository struct {
	Urls   []entities.URL
	Visits []entities.Visit
}

func NewStubURLsRepository() repositories.URLsRepository {
	if INSTANCE == nil {
		INSTANCE = &StubURLsRepository{
			Urls:   []entities.URL{},
			Visits: []entities.Visit{},
		}
	}

	return INSTANCE
}

func (r *StubURLsRepository) Create(url entities.URL) error {
	r.Urls = append(r.Urls, url)

	return nil
}

func (r *StubURLsRepository) FindByOriginalURL(originalURL string) (entities.URL, error) {
	for _, url := range r.Urls {
		if url.OriginalURL == originalURL {
			return url, nil
		}
	}

	return entities.URL{}, nil
}

func (r *StubURLsRepository) FindByShortURL(shortURL string) (entities.URL, error) {
	for _, url := range r.Urls {
		if url.ShortURL == shortURL {
			return url, nil
		}
	}

	return entities.URL{}, errors.New("URL not found")
}

func (r *StubURLsRepository) Visit(shortURL string) error {
	visit, err := entities.NewVisit(shortURL, time.Now())
	if err != nil {
		return err
	}

	r.Visits = append(r.Visits, *visit)

	return nil
}
