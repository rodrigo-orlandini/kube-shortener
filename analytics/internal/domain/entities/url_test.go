package entities_test

import (
	"testing"

	customError "rodrigoorlandini/urlshortener/analytics/internal/domain/custom-error"
	entities "rodrigoorlandini/urlshortener/analytics/internal/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	t.Run("Should create URL with all fields", func(t *testing.T) {
		shortURL := "abc123"
		id := "test-id-123"
		visits := 42

		url, err := entities.NewURL(shortURL,
			entities.WithID(id),
			entities.WithVisits(visits),
		)
		assert.NoError(t, err)

		assert.NotNil(t, url)
		assert.Equal(t, shortURL, url.ShortURL)
		assert.NotNil(t, url.ID)
		assert.Equal(t, id, *url.ID)
		assert.Equal(t, visits, url.Visits)
	})

	t.Run("Should create URL with partial options", func(t *testing.T) {
		shortURL := "xyz789"
		visits := 15

		url, err := entities.NewURL(shortURL, entities.WithVisits(visits))
		assert.NoError(t, err)

		assert.Equal(t, shortURL, url.ShortURL)
		assert.Equal(t, visits, url.Visits)
		assert.NotNil(t, url.ID)
	})

	t.Run("Should create URL with minimal fields", func(t *testing.T) {
		shortURL := "minimal123"

		url, err := entities.NewURL(shortURL)

		assert.NoError(t, err)
		assert.Equal(t, shortURL, url.ShortURL)
		assert.Equal(t, 0, url.Visits)
		assert.NotNil(t, url.ID)
	})

	t.Run("Should handle different short URL formats", func(t *testing.T) {
		testCases := []string{
			"abc123",
			"xyz789",
			"short-url",
			"url_123",
			"test-url-456",
			"simple",
			"complex-url-with-dashes",
			"url_with_underscores",
			"mixed-case-URL",
			"numbers123",
		}

		for _, testShortURL := range testCases {
			t.Run(testShortURL, func(t *testing.T) {
				url, err := entities.NewURL(testShortURL)
				assert.NoError(t, err)
				assert.Equal(t, testShortURL, url.ShortURL)
			})
		}
	})

	t.Run("Should return error for empty short URL", func(t *testing.T) {
		emptyShortURL := ""

		url, err := entities.NewURL(emptyShortURL)
		assert.Error(t, err)
		assert.Nil(t, url)

		invalidErr, ok := err.(*customError.InvalidEntityCreationError)
		assert.True(t, ok, "Expected InvalidEntityCreationError")
		assert.Equal(t, "URL", invalidErr.EntityName)
		assert.Equal(t, "shortURL", invalidErr.Field)
		assert.Equal(t, emptyShortURL, invalidErr.Value)
	})
}
