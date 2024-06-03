package config

import "time"

const (
	MaxRequests     = 100
	ExpirationTime  = 60 * time.Second
	CleanupInterval = 10 * time.Minute
)
