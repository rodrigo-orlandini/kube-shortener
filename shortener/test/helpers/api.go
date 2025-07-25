package helpers

import (
	"fmt"
	"log"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/http"
)

func StartAPI(port string) (string, func()) {
	server := http.NewServer()

	go func() {
		addr := ":" + port
		fmt.Println("Starting server on port", addr)

		if err := server.Start(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	return "http://localhost:" + port, func() {
		fmt.Println("Shutting down server...")
		if err := server.Shutdown(); err != nil {
			log.Printf("Error shutting down server: %v", err)
		}
	}
}
