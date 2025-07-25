package useCases_test

import (
	"testing"

	useCases "rodrigoorlandini/urlshortener/shortener/internal/application/use-cases"
	stubFactories "rodrigoorlandini/urlshortener/shortener/test/factories/use-cases"

	"github.com/stretchr/testify/assert"
)

func TestShortenURLUseCase(t *testing.T) {
	t.Run("Should create new shortened URL", func(t *testing.T) {
		originalURL := "https://www.example.com/very-long-path"
		useCase := stubFactories.NewShortenURLUseCaseFactory()

		request := useCases.ShortenURLUseCaseRequest{
			OriginalURL: originalURL,
		}

		response, err := useCase.Execute(request)

		assert.NoError(t, err)
		assert.NotEmpty(t, response.ShortenedURL)
		assert.Len(t, response.ShortenedURL, 9)
	})

	t.Run("Should return existing shortened URL when URL already exists", func(t *testing.T) {
		originalURL := "https://www.google.com"
		useCase := stubFactories.NewShortenURLUseCaseFactory()

		request := useCases.ShortenURLUseCaseRequest{
			OriginalURL: originalURL,
		}

		firstResponse, err := useCase.Execute(request)
		assert.NoError(t, err)
		assert.NotEmpty(t, firstResponse.ShortenedURL)

		secondResponse, err := useCase.Execute(request)
		assert.NoError(t, err)
		assert.Equal(t, firstResponse.ShortenedURL, secondResponse.ShortenedURL)
	})
}
