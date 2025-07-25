package stubFactories

import (
	"rodrigoorlandini/urlshortener/shortener/internal/application/services"
	useCases "rodrigoorlandini/urlshortener/shortener/internal/application/use-cases"
	stubRepositories "rodrigoorlandini/urlshortener/shortener/test/repositories"
)

func NewShortenURLUseCaseFactory() *useCases.ShortenURLUseCase {
	shortenerService := services.NewShortenerService()
	urlsRepository := stubRepositories.NewStubURLsRepository()

	shortenURLUseCase := useCases.NewShortenURLUseCase(shortenerService, urlsRepository)

	return shortenURLUseCase
}
