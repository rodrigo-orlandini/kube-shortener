package containers

import (
	"context"
	"fmt"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var CASSANDRA_CONTAINER *CassandraContainer

type CassandraContainer struct {
	host      string
	container testcontainers.Container
}

func NewCassandraContainer() *CassandraContainer {
	if CASSANDRA_CONTAINER == nil {
		CASSANDRA_CONTAINER = &CassandraContainer{
			host:      "",
			container: nil,
		}

		err := CASSANDRA_CONTAINER.startContainer()
		if err != nil {
			panic(err)
		}
	}

	return CASSANDRA_CONTAINER
}

func (c *CassandraContainer) startContainer() error {
	ctx := context.Background()

	request := testcontainers.ContainerRequest{
		Image:        "cassandra:5.0.4",
		ExposedPorts: []string{"9042"},
		Env: map[string]string{
			"CASSANDRA_CLUSTER_NAME": "test-cluster",
			"CASSANDRA_DC":           "datacenter1",
			"CASSANDRA_RACK":         "rack1",
		},
		WaitingFor: wait.ForListeningPort("9042").WithStartupTimeout(120 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})

	if err != nil {
		return fmt.Errorf("failed to create cassandra container: %w", err)
	}

	c.container = container

	host, err := container.Host(ctx)
	if err != nil {
		return fmt.Errorf("failed to get cassandra container host: %w", err)
	}

	mappedPort, err := container.MappedPort(ctx, "9042")
	if err != nil {
		return fmt.Errorf("failed to get cassandra container mapped port: %w", err)
	}

	c.host = fmt.Sprintf("%s:%s", host, mappedPort.Port())

	return nil
}

func (c *CassandraContainer) GetHost() string {
	return c.host
}
