package useCases

import (
	"rodrigoorlandini/urlshortener/analytics/internal/application/repositories"
	customError "rodrigoorlandini/urlshortener/analytics/internal/domain/custom-error"
	"rodrigoorlandini/urlshortener/analytics/internal/domain/entities"
	"time"
)

type HandleURLAccessedUseCase struct {
	urlsRepository repositories.URLsRepository
}

type HandleURLAccessedUseCaseRequest struct {
	ShortURL  string    `json:"short_url"`
	VisitedAt time.Time `json:"visited_at"`
}

type HandleURLAccessedUseCaseResponse struct{}

func NewHandleURLAccessedUseCase(urlsRepository repositories.URLsRepository) *HandleURLAccessedUseCase {
	return &HandleURLAccessedUseCase{
		urlsRepository: urlsRepository,
	}
}

func (u *HandleURLAccessedUseCase) Execute(request HandleURLAccessedUseCaseRequest) (HandleURLAccessedUseCaseResponse, error) {
	url, err := u.urlsRepository.FindByShortURL(request.ShortURL)

	if err != nil {
		if _, ok := err.(*customError.NotFoundError); !ok {
			return HandleURLAccessedUseCaseResponse{}, err
		}
	}

	if url.ShortURL == "" {
		newURL, err := entities.NewURL(request.ShortURL, entities.WithVisits(1))
		if err != nil {
			return HandleURLAccessedUseCaseResponse{}, err
		}

		err = u.urlsRepository.Create(*newURL)
		if err != nil {
			return HandleURLAccessedUseCaseResponse{}, err
		}

		return HandleURLAccessedUseCaseResponse{}, nil
	}

	err = u.urlsRepository.IncrementVisits(request.ShortURL)
	if err != nil {
		return HandleURLAccessedUseCaseResponse{}, err
	}

	return HandleURLAccessedUseCaseResponse{}, nil
}
