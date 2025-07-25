package useCases_test

import (
	"testing"

	useCases "rodrigoorlandini/urlshortener/shortener/internal/application/use-cases"
	stubFactories "rodrigoorlandini/urlshortener/shortener/test/factories/use-cases"

	"github.com/stretchr/testify/assert"
)

func TestAccessURLUseCase(t *testing.T) {
	t.Run("Should access existing shortened URL and increment visits", func(t *testing.T) {
		originalURL := "https://www.example.com/very-long-path"
		shortenUseCase := stubFactories.NewShortenURLUseCaseFactory()
		shortenRequest := useCases.ShortenURLUseCaseRequest{
			OriginalURL: originalURL,
		}
		shortenResponse, err := shortenUseCase.Execute(shortenRequest)
		assert.NoError(t, err)
		assert.NotEmpty(t, shortenResponse.ShortenedURL)

		accessUseCase := stubFactories.NewAccessURLUseCaseFactory()
		accessRequest := useCases.AccessURLUseCaseRequest{
			ShortURL: shortenResponse.ShortenedURL,
		}

		response, err := accessUseCase.Execute(accessRequest)

		assert.NoError(t, err)
		assert.Equal(t, originalURL, response.OriginalURL)
	})

	t.Run("Should return error for non-existent short URL", func(t *testing.T) {
		accessUseCase := stubFactories.NewAccessURLUseCaseFactory()

		request := useCases.AccessURLUseCaseRequest{
			ShortURL: "nonexistent",
		}

		response, err := accessUseCase.Execute(request)

		assert.Error(t, err)
		assert.Equal(t, "URL not found", err.Error())
		assert.Empty(t, response.OriginalURL)
	})
}
