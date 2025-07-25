package factories

import (
	useCases "rodrigoorlandini/urlshortener/analytics/internal/application/use-cases"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/database/repositories"
)

func NewHandleURLAccessedUseCase() (*useCases.HandleURLAccessedUseCase, error) {
	urlsRepository, err := repositories.NewPostgresURLsRepository()
	if err != nil {
		return nil, err
	}

	handleURLAccessedUseCase := useCases.NewHandleURLAccessedUseCase(urlsRepository)

	return handleURLAccessedUseCase, nil
}
