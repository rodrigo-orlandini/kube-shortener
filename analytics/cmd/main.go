package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"rodrigoorlandini/urlshortener/analytics/config"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/http"
)

func main() {
	server := http.NewServer()

	go func() {
		addr := ":8081"
		port := config.NewEnvironment().ApiPort

		if port != "" {
			addr = ":" + port
		}

		if err := server.Start(addr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := server.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
