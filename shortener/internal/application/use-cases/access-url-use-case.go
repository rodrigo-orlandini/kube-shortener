package useCases

import (
	"encoding/json"
	"rodrigoorlandini/urlshortener/shortener/internal/application/repositories"
	customError "rodrigoorlandini/urlshortener/shortener/internal/domain/custom-error"
	"rodrigoorlandini/urlshortener/shortener/internal/domain/entities"
	"rodrigoorlandini/urlshortener/shortener/internal/domain/events"
	"strings"
	"time"
)

type AccessURLUseCase struct {
	urlsRepository repositories.URLsRepository
	eventHandler   events.EventHandler
}

type AccessURLUseCaseRequest struct {
	ShortURL string `json:"shortURL"`
}

type AccessURLUseCaseResponse struct {
	OriginalURL string
}

func NewAccessURLUseCase(
	urlsRepository repositories.URLsRepository,
	eventHandler events.EventHandler,
) *AccessURLUseCase {
	return &AccessURLUseCase{
		urlsRepository: urlsRepository,
		eventHandler:   eventHandler,
	}
}

func (uc *AccessURLUseCase) Execute(request AccessURLUseCaseRequest) (AccessURLUseCaseResponse, error) {
	if request.ShortURL == "" {
		return AccessURLUseCaseResponse{}, &customError.InvalidEntityCreationError{
			EntityName: "URL",
			Field:      "shortURL",
			Value:      request.ShortURL,
		}
	}

	treatedShortURL := request.ShortURL
	if !strings.HasPrefix(treatedShortURL, "/") {
		treatedShortURL = "/" + treatedShortURL
	}

	url, err := uc.urlsRepository.FindByShortURL(treatedShortURL)
	if err != nil {
		return AccessURLUseCaseResponse{}, err
	}

	visit, err := entities.NewVisit(treatedShortURL, time.Now())
	if err != nil {
		return AccessURLUseCaseResponse{}, err
	}

	err = uc.urlsRepository.Visit(visit.ShortUrl)
	if err != nil {
		return AccessURLUseCaseResponse{}, err
	}

	data, err := json.Marshal(visit)
	if err != nil {
		return AccessURLUseCaseResponse{}, err
	}

	uc.eventHandler.Publish(events.EventURLAccessed, data)

	return AccessURLUseCaseResponse{
		OriginalURL: url.OriginalURL,
	}, nil
}
