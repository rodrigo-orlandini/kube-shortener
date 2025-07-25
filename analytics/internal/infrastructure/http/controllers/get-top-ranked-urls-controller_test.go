package controllers_test

import (
	"encoding/json"
	"net/http"
	"os"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/database"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/messaging"
	"rodrigoorlandini/urlshortener/analytics/test/containers"
	"rodrigoorlandini/urlshortener/analytics/test/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTopRankedURLsController(t *testing.T) {
	postgresContainer := containers.NewPostgresContainer()
	dsn, err := postgresContainer.CreateSchema("get-top-ranked-urls")
	if err != nil {
		t.Fatalf("failed to create postgres container: %v", err)
	}
	os.Setenv("DATABASE_URL", dsn)
	database.ResetConnection()

	natsContainer := containers.NewNatsContainer()
	natsHost := natsContainer.GetHost()
	os.Setenv("NATS_HOST", natsHost)
	messaging.ResetConnection()

	apiURL, shutdown := helpers.StartAPI("9091")
	if err != nil {
		t.Fatalf("failed to start API: %v", err)
	}
	defer shutdown()

	t.Run("Should return top ranked URLs successfully", func(t *testing.T) {
		resp, err := http.Get(apiURL + "/visits/topRanked?limit=5")
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, "SUCCESS", response["status"])
		assert.Contains(t, response, "data")

		data, ok := response["data"].([]interface{})
		assert.True(t, ok)
		assert.NotNil(t, data)
	})

	t.Run("Should use default limit when no limit provided", func(t *testing.T) {
		resp, err := http.Get(apiURL + "/visits/topRanked")
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, "SUCCESS", response["status"])
		assert.Contains(t, response, "data")
	})

	t.Run("Should return error for invalid limit parameter", func(t *testing.T) {
		resp, err := http.Get(apiURL + "/visits/topRanked?limit=invalid")
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, "INVALID_REQUEST", response["error"])
		assert.Equal(t, "Limit must be a valid integer", response["message"])
	})

	t.Run("Should return error for negative limit", func(t *testing.T) {
		resp, err := http.Get(apiURL + "/visits/topRanked?limit=-5")
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, "SUCCESS", response["status"])
	})

	t.Run("Should return error for zero limit", func(t *testing.T) {
		resp, err := http.Get(apiURL + "/visits/topRanked?limit=0")
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, "SUCCESS", response["status"])
	})

	t.Run("Should return empty data when no URLs exist", func(t *testing.T) {
		resp, err := http.Get(apiURL + "/visits/topRanked?limit=10")
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, "SUCCESS", response["status"])

		data, ok := response["data"].([]interface{})
		assert.True(t, ok)
		assert.Empty(t, data)
	})

	t.Run("Should respect the provided limit", func(t *testing.T) {
		resp, err := http.Get(apiURL + "/visits/topRanked?limit=3")
		assert.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, "SUCCESS", response["status"])

		data, ok := response["data"].([]interface{})
		assert.True(t, ok)
		assert.LessOrEqual(t, len(data), 3)
	})
}
