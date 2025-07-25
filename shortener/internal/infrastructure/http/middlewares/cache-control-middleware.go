package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

type CacheControlMiddleware struct{}

func NewCacheControlMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// It is required to prevent caching of the response
		// and not compute the URL visits counting
		c.Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Set("Pragma", "no-cache")
		c.Set("Expires", "0")

		return c.Next()
	}
}
