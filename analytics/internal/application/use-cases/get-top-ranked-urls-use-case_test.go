package useCases_test

import (
	"testing"

	useCases "rodrigoorlandini/urlshortener/analytics/internal/application/use-cases"
	"rodrigoorlandini/urlshortener/analytics/internal/domain/entities"
	stubFactories "rodrigoorlandini/urlshortener/analytics/test/factories/use-cases"
	stubRepositories "rodrigoorlandini/urlshortener/analytics/test/repositories"

	"github.com/stretchr/testify/assert"
)

func TestGetTopRankedURLsUseCase(t *testing.T) {
	t.Run("Should return top ranked URLs ordered by visit count", func(t *testing.T) {
		useCase := stubFactories.NewGetTopRankedURLsUseCaseFactory()
		stubRepo := stubRepositories.NewStubURLsRepository().(*stubRepositories.StubURLsRepository)
		stubRepo.Clear()

		url1, _ := entities.NewURL("abc123", entities.WithVisits(10))
		url2, _ := entities.NewURL("def456", entities.WithVisits(25))
		url3, _ := entities.NewURL("ghi789", entities.WithVisits(5))

		stubRepo.AddTestURL(*url1)
		stubRepo.AddTestURL(*url2)
		stubRepo.AddTestURL(*url3)

		request := useCases.GetTopRankedURLsUseCaseRequest{
			Limit: 3,
		}

		response, err := useCase.Execute(request)

		assert.NoError(t, err)
		assert.NotNil(t, response.TopRankedURLs)
		assert.Len(t, response.TopRankedURLs, 3)

		assert.Equal(t, "def456", response.TopRankedURLs[0].ShortURL)
		assert.Equal(t, 25, response.TopRankedURLs[0].VisitCount)
		assert.Equal(t, "abc123", response.TopRankedURLs[1].ShortURL)
		assert.Equal(t, 10, response.TopRankedURLs[1].VisitCount)
		assert.Equal(t, "ghi789", response.TopRankedURLs[2].ShortURL)
		assert.Equal(t, 5, response.TopRankedURLs[2].VisitCount)
	})

	t.Run("Should return empty list when no URLs exist", func(t *testing.T) {
		useCase := stubFactories.NewGetTopRankedURLsUseCaseFactory()
		stubRepo := stubRepositories.NewStubURLsRepository().(*stubRepositories.StubURLsRepository)
		stubRepo.Clear()

		request := useCases.GetTopRankedURLsUseCaseRequest{
			Limit: 10,
		}

		response, err := useCase.Execute(request)

		assert.NoError(t, err)
		assert.Empty(t, response.TopRankedURLs)
	})

	t.Run("Should respect the provided limit", func(t *testing.T) {
		useCase := stubFactories.NewGetTopRankedURLsUseCaseFactory()
		stubRepo := stubRepositories.NewStubURLsRepository().(*stubRepositories.StubURLsRepository)
		stubRepo.Clear()

		url1, _ := entities.NewURL("url1", entities.WithVisits(10))
		url2, _ := entities.NewURL("url2", entities.WithVisits(20))
		url3, _ := entities.NewURL("url3", entities.WithVisits(30))
		url4, _ := entities.NewURL("url4", entities.WithVisits(40))

		stubRepo.AddTestURL(*url1)
		stubRepo.AddTestURL(*url2)
		stubRepo.AddTestURL(*url3)
		stubRepo.AddTestURL(*url4)

		request := useCases.GetTopRankedURLsUseCaseRequest{
			Limit: 2,
		}

		response, err := useCase.Execute(request)

		assert.NoError(t, err)
		assert.Len(t, response.TopRankedURLs, 2)
		assert.Equal(t, "url4", response.TopRankedURLs[0].ShortURL)
		assert.Equal(t, 40, response.TopRankedURLs[0].VisitCount)
		assert.Equal(t, "url3", response.TopRankedURLs[1].ShortURL)
		assert.Equal(t, 30, response.TopRankedURLs[1].VisitCount)
	})
}
