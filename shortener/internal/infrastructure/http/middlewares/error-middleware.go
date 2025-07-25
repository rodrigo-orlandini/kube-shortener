package middlewares

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

type ErrorMiddleware struct{}

func NewErrorMiddleware() *ErrorMiddleware {
	return &ErrorMiddleware{}
}

func (h *ErrorMiddleware) Handle(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(ErrorResponse{
		Error:   "INTERNAL_ERROR",
		Message: err.Error(),
	})
}
