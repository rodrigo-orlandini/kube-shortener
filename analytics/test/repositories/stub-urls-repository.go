package repositories

import (
	"rodrigoorlandini/urlshortener/analytics/internal/application/repositories"
	"rodrigoorlandini/urlshortener/analytics/internal/domain/entities"
)

var INSTANCE *StubURLsRepository

type StubURLsRepository struct {
	urls []entities.URL
}

func NewStubURLsRepository() repositories.URLsRepository {
	if INSTANCE == nil {
		INSTANCE = &StubURLsRepository{
			urls: []entities.URL{},
		}
	}

	return INSTANCE
}

func (r *StubURLsRepository) Create(url entities.URL) error {
	r.urls = append(r.urls, url)
	return nil
}

func (r *StubURLsRepository) FindByShortURL(shortURL string) (entities.URL, error) {
	for _, url := range r.urls {
		if url.ShortURL == shortURL {
			return url, nil
		}
	}
	return entities.URL{}, nil
}

func (r *StubURLsRepository) IncrementVisits(shortURL string) error {
	for i, url := range r.urls {
		if url.ShortURL == shortURL {
			r.urls[i].Visits++
			return nil
		}
	}
	return nil
}

func (r *StubURLsRepository) GetTopRanked(limit int) ([]repositories.TopRankedURL, error) {
	var topRanked []repositories.TopRankedURL

	for _, url := range r.urls {
		topRanked = append(topRanked, repositories.TopRankedURL{
			ShortURL:   url.ShortURL,
			VisitCount: url.Visits,
		})
	}

	for i := 0; i < len(topRanked); i++ {
		for j := i + 1; j < len(topRanked); j++ {
			if topRanked[i].VisitCount < topRanked[j].VisitCount {
				topRanked[i], topRanked[j] = topRanked[j], topRanked[i]
			}
		}
	}

	if len(topRanked) > limit {
		topRanked = topRanked[:limit]
	}

	return topRanked, nil
}

func (r *StubURLsRepository) AddTestURL(url entities.URL) {
	r.urls = append(r.urls, url)
}

func (r *StubURLsRepository) Clear() {
	r.urls = []entities.URL{}
}

func (r *StubURLsRepository) GetURLs() []entities.URL {
	return r.urls
}
