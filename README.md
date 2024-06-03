# Rate-Limiting-API

This project is a simple API built with Go Fiber that demonstrates how to implement rate limiting on API endpoints to prevent abuse and ensure smooth performance.

## Project Structure

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