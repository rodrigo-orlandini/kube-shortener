package controllers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/database"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/messaging"
	"rodrigoorlandini/urlshortener/shortener/test/containers"
	"rodrigoorlandini/urlshortener/shortener/test/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessURLController(t *testing.T) {
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

	t.Run("Access URL", func(t *testing.T) {
		shortenedURL := createShortenedURL(t, apiURL, "https://www.teste.com")
		assert.NotEmpty(t, shortenedURL)

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}

		resp, err := client.Get(apiURL + shortenedURL)
		assert.NoError(t, err)
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMovedPermanently {
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Failed to read response body: %v", err)
				return
			}

			t.Logf("Status: %s", resp.Status)
			t.Logf("Response Body: %s", string(body))
		}

		assert.Equal(t, http.StatusMovedPermanently, resp.StatusCode)
		assert.Equal(t, "https://www.teste.com", resp.Header.Get("Location"))
	})
}

func createShortenedURL(t *testing.T, apiURL, originalURL string) string {
	requestBody := map[string]string{
		"originalUrl": originalURL,
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

	shortenedURL, ok := response["shortenedUrl"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, shortenedURL)

	return shortenedURL
}
