package http

import (
	"log"

	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/http/controllers"
	"rodrigoorlandini/urlshortener/shortener/internal/infrastructure/http/middlewares"

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
	cacheControlMiddleware := middlewares.NewCacheControlMiddleware()

	app := fiber.New(fiber.Config{
		ErrorHandler: errorMiddleware.Handle,
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(corsMiddleware)
	app.Use(cacheControlMiddleware)

	server := &Server{
		app: app,
	}

	server.setupRoutes()

	return server
}

func (s *Server) setupRoutes() {
	s.app.Get("/health", s.healthCheck)

	accessURLController := controllers.NewAccessURLController()
	s.app.Get("/:shortURL", accessURLController.Handle)

	shortenURLController := controllers.NewShortenURLController()
	s.app.Post("/shorten", shortenURLController.Handle)
}

func (s *Server) healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "SUCCESS",
		"message": "URL Shortener service is running",
	})
}

func (s *Server) Start(addr string) error {
	log.Printf("Starting server on %s", addr)
	return s.app.Listen(addr)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
