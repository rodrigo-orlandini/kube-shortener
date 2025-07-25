package config

import "os"

type Environment struct {
	ApiPort      string
	DatabaseHost string
	NatsHost     string
}

func NewEnvironment() *Environment {
	return &Environment{
		ApiPort:      os.Getenv("API_PORT"),
		DatabaseHost: os.Getenv("DATABASE_HOST"),
		NatsHost:     os.Getenv("NATS_HOST"),
	}
}
