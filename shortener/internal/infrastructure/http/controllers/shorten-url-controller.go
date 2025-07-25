package controllers

import (
	"errors"
	"rodrigoorlandini/urlshortener/shortener/internal/application/factories"
	useCases "rodrigoorlandini/urlshortener/shortener/internal/application/use-cases"
	customError "rodrigoorlandini/urlshortener/shortener/internal/domain/custom-error"

	"github.com/gofiber/fiber/v2"
)

type ShortenURLController struct{}

func NewShortenURLController() *ShortenURLController {
	return &ShortenURLController{}
}

func (c *ShortenURLController) Handle(ctx *fiber.Ctx) error {
	useCase, err := factories.NewShortenURLUseCaseFactory()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "ERROR",
			"message": "Failed to initialize database connection: " + err.Error(),
		})
	}

	request, err := c.parseBody(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "ERROR",
			"message": "Invalid request body: " + err.Error(),
		})
	}

	response, err := useCase.Execute(request)
	if err != nil {
		var invalidEntityErr *customError.InvalidEntityCreationError
		if errors.As(err, &invalidEntityErr) {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":  "ERROR",
				"message": "Invalid request body: " + err.Error(),
			})
		}

		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "ERROR",
			"message": "Internal server error: " + err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"shortenedUrl": response.ShortenedURL,
	})
}

func (c *ShortenURLController) parseBody(ctx *fiber.Ctx) (useCases.ShortenURLUseCaseRequest, error) {
	var body useCases.ShortenURLUseCaseRequest

	if err := ctx.BodyParser(&body); err != nil {
		return useCases.ShortenURLUseCaseRequest{}, err
	}

	return body, nil
}
