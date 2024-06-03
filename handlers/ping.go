package handlers

import "github.com/gofiber/fiber/v2"

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "Successful",
		"body":   "Hi! You've reached the API. How may I help you?",
	})
}
