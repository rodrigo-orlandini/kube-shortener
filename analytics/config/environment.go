package config

import "os"

type Environment struct {
	ApiPort     string
	DatabaseURL string
	NatsHost    string
}

func NewEnvironment() *Environment {
	return &Environment{
		ApiPort:     os.Getenv("API_PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		NatsHost:    os.Getenv("NATS_HOST"),
	}
}
