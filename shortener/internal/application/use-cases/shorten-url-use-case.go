package useCases

import (
	"rodrigoorlandini/urlshortener/shortener/internal/application/repositories"
	"rodrigoorlandini/urlshortener/shortener/internal/application/services"
	"rodrigoorlandini/urlshortener/shortener/internal/domain/entities"
)

type ShortenURLUseCase struct {
	shortenerService *services.ShortenerService
	urlsRepository   repositories.URLsRepository
}

type ShortenURLUseCaseRequest struct {
	OriginalURL string `json:"originalUrl"`
}

type ShortenURLUseCaseResponse struct {
	ShortenedURL string
}

func NewShortenURLUseCase(
	shortenerService *services.ShortenerService,
	urlsRepository repositories.URLsRepository,
) *ShortenURLUseCase {
	return &ShortenURLUseCase{
		shortenerService: shortenerService,
		urlsRepository:   urlsRepository,
	}
}

func (u *ShortenURLUseCase) Execute(request ShortenURLUseCaseRequest) (ShortenURLUseCaseResponse, error) {
	existingURL, err := u.urlsRepository.FindByOriginalURL(request.OriginalURL)
	if err != nil {
		return ShortenURLUseCaseResponse{}, err
	}

	if existingURL.OriginalURL != "" {
		return ShortenURLUseCaseResponse{
			ShortenedURL: existingURL.ShortURL,
		}, nil
	}

	url, err := entities.NewURL(request.OriginalURL)
	if err != nil {
		return ShortenURLUseCaseResponse{}, err
	}

	shortURL, err := u.shortenerService.ShortenURL(*url)
	if err != nil {
		return ShortenURLUseCaseResponse{}, err
	}

	err = u.urlsRepository.Create(shortURL)
	if err != nil {
		return ShortenURLUseCaseResponse{}, err
	}

	return ShortenURLUseCaseResponse{
		ShortenedURL: shortURL.ShortURL,
	}, nil
}
