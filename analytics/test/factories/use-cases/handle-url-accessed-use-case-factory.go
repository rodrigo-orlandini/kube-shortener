package useCases

import (
	useCases "rodrigoorlandini/urlshortener/analytics/internal/application/use-cases"
	stubRepositories "rodrigoorlandini/urlshortener/analytics/test/repositories"
)

func NewHandleURLAccessedUseCaseFactory() *useCases.HandleURLAccessedUseCase {
	urlsRepository := stubRepositories.NewStubURLsRepository()
	return useCases.NewHandleURLAccessedUseCase(urlsRepository)
}
