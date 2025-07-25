package controllers

import (
	"rodrigoorlandini/urlshortener/shortener/internal/application/factories"
	useCases "rodrigoorlandini/urlshortener/shortener/internal/application/use-cases"

	"github.com/gofiber/fiber/v2"
)

type AccessURLController struct{}

func NewAccessURLController() *AccessURLController {
	return &AccessURLController{}
}

func (c *AccessURLController) Handle(ctx *fiber.Ctx) error {
	useCase, err := factories.NewAccessURLUseCaseFactory()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "INTERNAL_ERROR",
			"message": "Failed to initialize access URL use case: " + err.Error(),
		})
	}

	shortURL := ctx.Params("shortURL")
	if shortURL == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "INVALID_REQUEST",
			"message": "Short URL parameter is required",
		})
	}

	request := useCases.AccessURLUseCaseRequest{
		ShortURL: shortURL,
	}

	response, err := useCase.Execute(request)
	if err != nil {
		if err.Error() == "URL not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "URL_NOT_FOUND",
				"message": "The requested short URL was not found",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "INTERNAL_ERROR",
			"message": err.Error(),
		})
	}

	return ctx.Redirect(response.OriginalURL, fiber.StatusMovedPermanently)
}
