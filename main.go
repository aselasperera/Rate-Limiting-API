package main

import (
	"Rate-Limiting-API/handlers"
	"Rate-Limiting-API/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Apply the rate limiting middleware globally
	app.Use(middleware.RateLimiter())

	// Define your routes
	app.Get("/", handlers.Home)
	app.Get("/ping", handlers.Ping)

	// Start the server
	app.Listen(":8080")
}
