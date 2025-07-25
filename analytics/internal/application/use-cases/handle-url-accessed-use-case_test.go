package useCases_test

import (
	"testing"

	useCases "rodrigoorlandini/urlshortener/analytics/internal/application/use-cases"
	"rodrigoorlandini/urlshortener/analytics/internal/domain/entities"
	stubFactories "rodrigoorlandini/urlshortener/analytics/test/factories/use-cases"
	stubRepositories "rodrigoorlandini/urlshortener/analytics/test/repositories"

	"github.com/stretchr/testify/assert"
)

func TestHandleURLAccessedUseCase(t *testing.T) {
	t.Run("Should increment visits for existing URL", func(t *testing.T) {
		useCase := stubFactories.NewHandleURLAccessedUseCaseFactory()
		stubRepo := stubRepositories.NewStubURLsRepository().(*stubRepositories.StubURLsRepository)
		stubRepo.Clear()

		url, _ := entities.NewURL("abc123", entities.WithVisits(10))
		stubRepo.AddTestURL(*url)

		request := useCases.HandleURLAccessedUseCaseRequest{
			ShortURL: "abc123",
		}

		_, err := useCase.Execute(request)
		assert.NoError(t, err)

		existingURLs := stubRepo.GetURLs()
		assert.Equal(t, 1, len(existingURLs))
		assert.Equal(t, 11, existingURLs[0].Visits)
	})

	t.Run("Should create new URL if it doesn't exist", func(t *testing.T) {
		useCase := stubFactories.NewHandleURLAccessedUseCaseFactory()
		stubRepo := stubRepositories.NewStubURLsRepository().(*stubRepositories.StubURLsRepository)
		stubRepo.Clear()

		request := useCases.HandleURLAccessedUseCaseRequest{
			ShortURL: "abc123",
		}

		_, err := useCase.Execute(request)
		assert.NoError(t, err)

		existingURLs := stubRepo.GetURLs()
		assert.Equal(t, 1, len(existingURLs))
		assert.Equal(t, 1, existingURLs[0].Visits)
	})
}
