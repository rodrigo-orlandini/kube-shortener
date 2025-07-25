package containers

import (
	"context"
	"fmt"

	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var NATS_CONTAINER *NatsContainer

type NatsContainer struct {
	host      string
	container testcontainers.Container
}

func NewNatsContainer() *NatsContainer {
	if NATS_CONTAINER == nil {
		NATS_CONTAINER = &NatsContainer{}

		err := NATS_CONTAINER.startContainer()
		if err != nil {
			panic(err)
		}
	}

	return NATS_CONTAINER
}

func (c *NatsContainer) startContainer() error {
	ctx := context.Background()

	request := testcontainers.ContainerRequest{
		Image:        "nats:2.11.6",
		ExposedPorts: []string{"4222"},
		WaitingFor:   wait.ForLog("Listening for client connections on 0.0.0.0:4222").WithStartupTimeout(30 * time.Second),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: request,
		Started:          true,
	})

	if err != nil {
		return fmt.Errorf("failed to create nats container: %w", err)
	}

	c.container = container

	host, err := container.Host(ctx)
	if err != nil {
		return fmt.Errorf("failed to get nats container host: %w", err)
	}

	mappedPort, err := container.MappedPort(ctx, "4222")
	if err != nil {
		return fmt.Errorf("failed to get nats container mapped port: %w", err)
	}

	c.host = fmt.Sprintf("nats://%s:%s", host, mappedPort.Port())

	return nil
}

func (c *NatsContainer) GetHost() string {
	return c.host
}
