package useCases

import (
	"rodrigoorlandini/urlshortener/analytics/internal/application/repositories"
)

type GetTopRankedURLsUseCase struct {
	urlsRepository repositories.URLsRepository
}

type GetTopRankedURLsUseCaseRequest struct {
	Limit int
}

type GetTopRankedURLsUseCaseResponse struct {
	TopRankedURLs []repositories.TopRankedURL
}

func NewGetTopRankedURLsUseCase(
	urlsRepository repositories.URLsRepository,
) *GetTopRankedURLsUseCase {
	return &GetTopRankedURLsUseCase{
		urlsRepository: urlsRepository,
	}
}

func (u *GetTopRankedURLsUseCase) Execute(request GetTopRankedURLsUseCaseRequest) (GetTopRankedURLsUseCaseResponse, error) {
	if request.Limit <= 0 {
		request.Limit = 5
	}

	topRankedURLs, err := u.urlsRepository.GetTopRanked(request.Limit)
	if err != nil {
		return GetTopRankedURLsUseCaseResponse{}, err
	}

	return GetTopRankedURLsUseCaseResponse{
		TopRankedURLs: topRankedURLs,
	}, nil
}
