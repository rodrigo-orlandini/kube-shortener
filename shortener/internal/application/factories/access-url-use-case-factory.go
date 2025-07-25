package factories

import (
	useCases "rodrigoorlandini/urlshortener/shortener/internal/application/use-cases"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/database/repositories"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/messaging"
)

func NewAccessURLUseCaseFactory() (*useCases.AccessURLUseCase, error) {
	urlsRepository := repositories.NewCassandraUrlsRepository()
	eventHandler := messaging.NewNats()

	accessURLUseCase := useCases.NewAccessURLUseCase(urlsRepository, eventHandler)

	return accessURLUseCase, nil
}
