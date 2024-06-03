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

## Getting Started

### Prerequisites

- Go 1.16+ installed on your machine

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/aselasperera/Rate-Limiting-API
   cd Rate-Limiting-API
