package messaging

import (
	"fmt"
	"log"
	"rodrigoorlandini/urlshortener/analytics/config"
	"rodrigoorlandini/urlshortener/analytics/internal/domain/events"

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

func (n *Nats) Subscribe(subject events.EventSubject, handler func(data []byte) error) error {
	_, err := n.connection.Subscribe(string(subject), func(m *nats.Msg) {
		err := handler(m.Data)

		if err != nil {
			log.Println("Error while handling event:", err)
		}
	})

	return err
}

func ResetConnection() {
	CONNECTION = nil
}
