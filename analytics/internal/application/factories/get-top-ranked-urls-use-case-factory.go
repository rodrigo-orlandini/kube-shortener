package factories

import (
	useCases "rodrigoorlandini/urlshortener/analytics/internal/application/use-cases"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/database/repositories"
)

func NewGetTopRankedURLsUseCase() (*useCases.GetTopRankedURLsUseCase, error) {
	urlsRepository, err := repositories.NewPostgresURLsRepository()
	if err != nil {
		return nil, err
	}

	getTopRankedURLsUseCase := useCases.NewGetTopRankedURLsUseCase(urlsRepository)

	return getTopRankedURLsUseCase, nil
}
