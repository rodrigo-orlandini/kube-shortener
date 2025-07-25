package entities_test

import (
	"testing"
	"time"

	customError "rodrigoorlandini/urlshortener/shortener/internal/domain/custom-error"
	entities "rodrigoorlandini/urlshortener/shortener/internal/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestVisit(t *testing.T) {
	t.Run("Should create a visit", func(t *testing.T) {
		shortURL := "abc123"
		visitedAt := time.Now()

		visit, err := entities.NewVisit(shortURL, visitedAt)
		assert.NoError(t, err)

		assert.NotNil(t, visit)
		assert.Equal(t, shortURL, visit.ShortUrl)
		assert.Equal(t, visitedAt, visit.VisitedAt)
	})

	t.Run("Should return error for empty URL", func(t *testing.T) {
		emptyURL := ""
		visitedAt := time.Now()

		visit, err := entities.NewVisit(emptyURL, visitedAt)
		assert.Error(t, err)
		assert.Nil(t, visit)

		invalidErr, ok := err.(*customError.InvalidEntityCreationError)
		assert.True(t, ok, "Expected InvalidEntityCreationError")
		assert.Equal(t, "Visit", invalidErr.EntityName)
		assert.Equal(t, "shortURL", invalidErr.Field)
		assert.Equal(t, emptyURL, invalidErr.Value)
	})
}
