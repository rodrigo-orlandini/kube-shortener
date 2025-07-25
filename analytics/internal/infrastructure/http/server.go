package http

import (
	"log"

	"rodrigoorlandini/urlshortener/analytics/internal/domain/events"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/http/controllers"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/http/middlewares"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/messaging"
	"rodrigoorlandini/urlshortener/analytics/internal/infrastructure/messaging/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	errorMiddleware := middlewares.NewErrorMiddleware()
	corsMiddleware := middlewares.NewCorsMiddleware()

	app := fiber.New(fiber.Config{
		ErrorHandler: errorMiddleware.Handle,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(corsMiddleware)

	server := &Server{
		app: app,
	}

	server.setupRoutes()
	server.setupEventHandlers()

	return server
}

func (s *Server) setupRoutes() {
	s.app.Get("/health", s.healthCheck)

	getTopRankedURLsController := controllers.NewGetTopRankedURLsController()
	s.app.Get("/visits/topRanked", getTopRankedURLsController.Handle)
}

func (s *Server) setupEventHandlers() {
	eventHandler := messaging.NewNats()

	handleURLAccessedHandler := handlers.NewURLAccessedHandler()
	eventHandler.Subscribe(events.EventURLAccessed, handleURLAccessedHandler.Handle)
}

func (s *Server) healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "SUCCESS",
		"message": "Analytics service is running",
	})
}

func (s *Server) Start(addr string) error {
	log.Printf("Starting server on %s", addr)
	return s.app.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
