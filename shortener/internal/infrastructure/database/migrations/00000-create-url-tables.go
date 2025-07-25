package migrations

import "github.com/gocql/gocql"

type CreateURLTablesMigration struct{}

func RunCreateURLTablesMigration(session *gocql.Session) error {
	migrationQueries := []string{
		`CREATE KEYSPACE IF NOT EXISTS shortener WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};`,
		`CREATE TABLE IF NOT EXISTS shortener.urls_by_original (
			"originalUrl" TEXT PRIMARY KEY,
			"shortUrl" TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS shortener.urls_by_short (
			"shortUrl" TEXT PRIMARY KEY,
			"originalUrl" TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS shortener.visits (
			id UUID PRIMARY KEY,
			"shortUrl" TEXT,
			"visitedAt" TIMESTAMP
		);`,
	}

	for _, query := range migrationQueries {
		if err := session.Query(query).Exec(); err != nil {
			return err
		}
	}

	return nil
}
