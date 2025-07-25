package entities_test

import (
	"testing"

	customError "rodrigoorlandini/urlshortener/shortener/internal/domain/custom-error"
	entities "rodrigoorlandini/urlshortener/shortener/internal/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	t.Run("Should create URL with all fields", func(t *testing.T) {
		originalURL := "https://www.example.com/path?param=value#fragment"
		shortURL := "abc123"

		url, err := entities.NewURL(originalURL,
			entities.WithShortURL(shortURL),
		)
		assert.NoError(t, err)

		assert.NotNil(t, url)
		assert.Equal(t, originalURL, url.OriginalURL)
		assert.Equal(t, shortURL, url.ShortURL)
	})

	t.Run("Should create URL with partial options", func(t *testing.T) {
		originalURL := "https://example.org"

		url, err := entities.NewURL(originalURL)
		assert.NoError(t, err)

		assert.Equal(t, originalURL, url.OriginalURL)
	})

	t.Run("Should handle different URL formats", func(t *testing.T) {
		testCases := []string{
			"https://example.com",
			"http://example.org",
			"https://subdomain.example.com",
			"https://example.com:8080",
			"https://example.com/path",
			"https://example.com/path?param=value",
			"https://example.com/path#fragment",
			"https://example.com/path?param=value#fragment",
			"example.com",
			"subdomain.example.org",
		}

		for _, testURL := range testCases {
			t.Run(testURL, func(t *testing.T) {
				url, err := entities.NewURL(testURL)
				assert.NoError(t, err)
				assert.Equal(t, testURL, url.OriginalURL)
			})
		}
	})

	t.Run("Should return error for empty URL", func(t *testing.T) {
		emptyURL := ""

		url, err := entities.NewURL(emptyURL)
		assert.Error(t, err)
		assert.Nil(t, url)

		invalidErr, ok := err.(*customError.InvalidEntityCreationError)
		assert.True(t, ok, "Expected InvalidEntityCreationError")
		assert.Equal(t, "URL", invalidErr.EntityName)
		assert.Equal(t, "originalURL", invalidErr.Field)
		assert.Equal(t, emptyURL, invalidErr.Value)
	})

	t.Run("Should return error for invalid URL format", func(t *testing.T) {
		invalidURLs := []string{
			"not-a-url",
			"ftp://example.com",
			"https://",
			"http://invalid",
			"://example.com",
			"https://.com",
			"https://example.",
			"",
			"   ",
			"invalid domain",
			"https://",
			"http://",
		}

		for _, invalidURL := range invalidURLs {
			t.Run(invalidURL, func(t *testing.T) {
				url, err := entities.NewURL(invalidURL)
				assert.Error(t, err)
				assert.Nil(t, url)

				invalidErr, ok := err.(*customError.InvalidEntityCreationError)
				assert.True(t, ok, "Expected InvalidEntityCreationError for URL '%s'", invalidURL)
				assert.Equal(t, "URL", invalidErr.EntityName)
				assert.Equal(t, "originalURL", invalidErr.Field)
				assert.Equal(t, invalidURL, invalidErr.Value)
			})
		}
	})
}
