# Rate-Limiting-API

This project is a simple API built with Go Fiber that demonstrates how to implement rate limiting on API endpoints to prevent abuse and ensure smooth performance.

## Project Structure
```
Rate-Limiting-API/
├── main.go
├── handlers/
│ ├── ping.go
│ └── home.go
├── middleware/
│ └── rate_limiter.go
├── config/
│ └── config.go
├── go.mod
└── go.sum
```
## Getting Started

### Prerequisites

- Go 1.16+ installed on your machine

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/aselasperera/Rate-Limiting-API
   cd Rate-Limiting-API
2. Install the dependencies:
    ```bash
    go mod tidy

### API Endpoints
### Home Endpoint
- URL: /
- Method: GET
- Response:
    ```bash
    
    Welcome to the API!

### Ping Endpoint
- URL: /ping
- Method: GET
- Response:
     ```bash
     {
    "status": "Successful",
    "body": "Hi! You've reached the API. How may I help you?"
     }
### Rate Limiting
The API is protected by rate limiting middleware to prevent abuse. The current configuration allows a maximum of 5 requests per 10 seconds per IP address.

Rate Limit Exceeded Response:
      ```bash
      
      {
       "error": "Too many requests, please try again later."
      }
### Customizing the Rate Limiter
You can customize the rate limiter configuration in middleware/rate_limiter.go:
      ```bash
      
      {
      package middleware

      import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/limiter"
    "time"
      )

      func RateLimiter() fiber.Handler {
    return limiter.New(limiter.Config{
        Max:        5,              // Maximum number of requests
        Expiration: 10 * time.Second, // Time period for the rate limit
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


### License

This project is licensed under the MIT License. See the LICENSE file for details.


