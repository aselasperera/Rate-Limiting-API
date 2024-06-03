package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,              // Maximum number of requests
		Expiration: 60 * time.Second, // Time period for the rate limit
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Use IP address as the key for rate limiting
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests, please try again later.",
			})
		},
	})
}
