package controllers

import (
	"strconv"

	"rodrigoorlandini/urlshortener/analytics/internal/application/factories"
	useCases "rodrigoorlandini/urlshortener/analytics/internal/application/use-cases"

	"github.com/gofiber/fiber/v2"
)

type GetTopRankedURLsController struct{}

func NewGetTopRankedURLsController() *GetTopRankedURLsController {
	return &GetTopRankedURLsController{}
}

func (c *GetTopRankedURLsController) Handle(ctx *fiber.Ctx) error {
	useCase, err := factories.NewGetTopRankedURLsUseCase()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "INTERNAL_ERROR",
			"message": "Failed to initialize get top ranked URLs use case: " + err.Error(),
		})
	}

	limitStr := ctx.Query("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "INVALID_REQUEST",
			"message": "Limit must be a valid integer",
		})
	}

	request := useCases.GetTopRankedURLsUseCaseRequest{
		Limit: limit,
	}

	response, err := useCase.Execute(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "INTERNAL_ERROR",
			"message": "Failed to get top ranked URLs",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   response.TopRankedURLs,
	})
}
