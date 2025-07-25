package stubFactories

import (
	useCases "rodrigoorlandini/urlshortener/shortener/internal/application/use-cases"
	stubMessaging "rodrigoorlandini/urlshortener/shortener/test/messaging"
	stubRepositories "rodrigoorlandini/urlshortener/shortener/test/repositories"
)

func NewAccessURLUseCaseFactory() *useCases.AccessURLUseCase {
	stubURLsRepository := stubRepositories.NewStubURLsRepository()
	stubEventHandler := stubMessaging.NewStubEventHandler()

	return useCases.NewAccessURLUseCase(stubURLsRepository, stubEventHandler)
}
