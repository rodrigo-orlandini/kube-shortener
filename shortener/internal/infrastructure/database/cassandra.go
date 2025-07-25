package database

import (
	"log"
	"rodrigoorlandini/urlshortener/shortener/config"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/database/migrations"
	"time"

	"github.com/gocql/gocql"
)

var SESSION *gocql.Session

type Cassandra struct {
	session *gocql.Session
}

func NewCassandra() *Cassandra {
	if SESSION == nil {
		environment := config.NewEnvironment()

		cluster := gocql.NewCluster(environment.DatabaseHost)
		cluster.Consistency = gocql.Quorum

		prepareSession, err := cluster.CreateSession()
		if err != nil {
			log.Fatal("Error creating session with Cassandra:", err)
		}

		err = migrations.RunCreateURLTablesMigration(prepareSession)
		if err != nil {
			log.Fatal("Error while applying migrations:", err)
		}

		prepareSession.Close()

		clusterWithKeyspace := gocql.NewCluster(environment.DatabaseHost)
		clusterWithKeyspace.Keyspace = "shortener"
		clusterWithKeyspace.Consistency = gocql.Quorum
		clusterWithKeyspace.Timeout = 10 * time.Second
		clusterWithKeyspace.ConnectTimeout = 10 * time.Second
		clusterWithKeyspace.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}

		session, err := clusterWithKeyspace.CreateSession()
		if err != nil {
			log.Fatal("Error creating session with Cassandra keyspace:", err)
		}

		SESSION = session
	}

	return &Cassandra{
		session: SESSION,
	}
}

func (c *Cassandra) GetConnection() *gocql.Session {
	return SESSION
}

func ResetConnection() {
	SESSION = nil
}
