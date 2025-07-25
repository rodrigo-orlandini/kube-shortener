package useCases

import (
	useCases "rodrigoorlandini/urlshortener/analytics/internal/application/use-cases"
	stubRepositories "rodrigoorlandini/urlshortener/analytics/test/repositories"
)

func NewGetTopRankedURLsUseCaseFactory() *useCases.GetTopRankedURLsUseCase {
	urlsRepository := stubRepositories.NewStubURLsRepository()
	return useCases.NewGetTopRankedURLsUseCase(urlsRepository)
}
