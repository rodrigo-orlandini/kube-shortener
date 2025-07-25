package repositories

import (
	"rodrigoorlandini/urlshortener/shortener/internal/application/repositories"
	"rodrigoorlandini/urlshortener/shortener/internal/domain/entities"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/database"
	"time"

	"github.com/gocql/gocql"
)

type CassandraUrlsRepository struct {
	session *gocql.Session
}

func NewCassandraUrlsRepository() repositories.URLsRepository {
	session := database.NewCassandra().GetConnection()

	return CassandraUrlsRepository{
		session: session,
	}
}

func (r CassandraUrlsRepository) Create(url entities.URL) error {
	err := r.session.Query(`
		INSERT INTO shortener.urls_by_original ("originalUrl", "shortUrl")
		VALUES (?, ?)`,
		url.OriginalURL, url.ShortURL,
	).Exec()

	if err != nil {
		return err
	}

	err = r.session.Query(`
		INSERT INTO shortener.urls_by_short ("shortUrl", "originalUrl")
		VALUES (?, ?)`,
		url.ShortURL, url.OriginalURL,
	).Exec()

	if err != nil {
		return err
	}

	return nil
}

func (r CassandraUrlsRepository) FindByOriginalURL(data string) (entities.URL, error) {
	var originalUrl string
	var shortUrl string

	err := r.session.Query(`
		SELECT "originalUrl", "shortUrl" FROM shortener.urls_by_original
		WHERE "originalUrl" = ? LIMIT 1`,
		data,
	).Scan(&originalUrl, &shortUrl)

	if err != nil {
		if err == gocql.ErrNotFound {
			return entities.URL{}, nil
		}

		return entities.URL{}, err
	}

	result, err := entities.NewURL(originalUrl, entities.WithShortURL(shortUrl))

	if err != nil {
		return entities.URL{}, err
	}

	return *result, nil
}

func (r CassandraUrlsRepository) FindByShortURL(data string) (entities.URL, error) {
	var originalUrl string
	var shortUrl string

	err := r.session.Query(`
		SELECT "originalUrl", "shortUrl" FROM shortener.urls_by_short
		WHERE "shortUrl" = ? LIMIT 1`,
		data,
	).Scan(&originalUrl, &shortUrl)

	if err != nil {
		if err == gocql.ErrNotFound {
			return entities.URL{}, nil
		}

		return entities.URL{}, err
	}

	result, err := entities.NewURL(originalUrl, entities.WithShortURL(shortUrl))

	if err != nil {
		return entities.URL{}, err
	}

	return *result, nil
}

func (r CassandraUrlsRepository) Visit(shortUrl string) error {
	id := gocql.TimeUUID()

	err := r.session.Query(`
		INSERT INTO shortener.visits (id, "shortUrl", "visitedAt")
		VALUES (?, ?, ?)`,
		id, shortUrl, time.Now(),
	).Exec()

	if err != nil {
		return err
	}

	return nil
}
