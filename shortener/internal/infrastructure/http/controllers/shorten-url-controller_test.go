package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/database"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/messaging"
	"rodrigoorlandini/urlshortener/shortener/test/containers"
	"rodrigoorlandini/urlshortener/shortener/test/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenURLController(t *testing.T) {
	cassandraContainer := containers.NewCassandraContainer()
	cassandraHost := cassandraContainer.GetHost()
	os.Setenv("DATABASE_HOST", cassandraHost)
	database.ResetConnection()

	natsContainer := containers.NewNatsContainer()
	natsHost := natsContainer.GetHost()
	os.Setenv("NATS_HOST", natsHost)
	messaging.ResetConnection()

	apiURL, shutdown := helpers.StartAPI("9090")
	defer shutdown()

	t.Run("Shorten URL", func(t *testing.T) {
		requestBody := map[string]string{
			"originalUrl": "https://www.google.com",
		}

		jsonBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		resp, err := http.Post(apiURL+"/shorten", "application/json", bytes.NewBuffer(jsonBody))
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Contains(t, response, "shortenedUrl")

		shortenedURL, ok := response["shortenedUrl"].(string)
		assert.True(t, ok)
		assert.NotEmpty(t, shortenedURL)
	})
}
