package factories

import (
	"rodrigoorlandini/urlshortener/shortener/internal/application/services"
	useCases "rodrigoorlandini/urlshortener/shortener/internal/application/use-cases"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/database/repositories"
)

func NewShortenURLUseCaseFactory() (*useCases.ShortenURLUseCase, error) {
	shortenerService := services.NewShortenerService()
	urlsRepository := repositories.NewCassandraUrlsRepository()

	shortenURLUseCase := useCases.NewShortenURLUseCase(shortenerService, urlsRepository)

	return shortenURLUseCase, nil
}
