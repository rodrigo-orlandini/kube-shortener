package containers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var INSTANCE *PostgresContainer

type PostgresContainer struct {
	dsn       string
	container testcontainers.Container
}

func NewPostgresContainer() *PostgresContainer {
	if INSTANCE == nil {
		INSTANCE = &PostgresContainer{
			dsn:       "",
			container: nil,
		}

		err := INSTANCE.startContainer()
		if err != nil {
			panic(err)
		}
	}

	return INSTANCE
}

func (c *PostgresContainer) CreateSchema(schemaName string) (string, error) {
	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("pgx", c.dsn)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = db.PingContext(ctx)
		cancel()

		if err == nil {
			break
		}

		db.Close()
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return "", fmt.Errorf("failed to connect to postgres after retries: %w", err)
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS \"%s\";", schemaName)
	_, err = db.Exec(query)
	if err != nil {
		return "", fmt.Errorf("failed to create schema: %w", err)
	}

	return c.dsn, nil
}

func (c *PostgresContainer) startContainer() error {
	ctx := context.Background()

	request := testcontainers.ContainerRequest{
		Image:        "postgres:16",
		ExposedPorts: []string{"5432"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "analytics",
		},
		WaitingFor: wait.ForLog("database system is ready to accept connections"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})

	if err != nil {
		return fmt.Errorf("failed to create postgres container: %w", err)
	}

	c.container = container

	host, err := container.Host(ctx)
	if err != nil {
		return fmt.Errorf("failed to get postgres container host: %w", err)
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return fmt.Errorf("failed to get postgres container mapped port: %w", err)
	}

	c.dsn = fmt.Sprintf("postgres://postgres:postgres@%s:%s/analytics?sslmode=disable", host, mappedPort.Port())

	return nil
}
