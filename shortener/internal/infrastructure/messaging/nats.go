package messaging

import (
	"fmt"
	"log"
	"rodrigoorlandini/urlshortener/shortener/config"
	"rodrigoorlandini/urlshortener/shortener/internal/domain/events"

	"github.com/nats-io/nats.go"
)

var CONNECTION *nats.Conn

type Nats struct {
	connection *nats.Conn
}

func NewNats() *Nats {
	if CONNECTION == nil {
		natsHost := config.NewEnvironment().NatsHost

		if natsHost == "" {
			natsHost = "nats://localhost:4222"
		}

		nats, err := nats.Connect(natsHost)
		if err != nil {
			log.Fatal("Error while starting Nats:", nats)
		}

		fmt.Println("Nats connected to:", natsHost)
		CONNECTION = nats
	}

	return &Nats{
		connection: CONNECTION,
	}
}

func (n *Nats) Publish(subject events.EventSubject, data []byte) error {
	return n.connection.Publish(string(subject), data)
}

func ResetConnection() {
	CONNECTION = nil
}
