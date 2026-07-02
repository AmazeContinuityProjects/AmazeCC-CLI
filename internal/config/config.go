// Package config holds runtime settings (API URL, timeouts, etc.).
// Values come from environment variables so you can point at dev/staging/prod
// without recompiling.
package config

import (
	"os"
	"time"
)

const (
	defaultAPIBaseURL = "https://api.amazecc.com"
	defaultAPITimeout = 10 * time.Second
)

// Config is the full app configuration loaded at startup.
type Config struct {
	APIBaseURL string
	APITimeout time.Duration
}

// Load reads config from the environment with sensible production defaults.
func Load() Config {
	return Config{
		APIBaseURL: envOr("AMAZE_API_URL", defaultAPIBaseURL),
		APITimeout: defaultAPITimeout,
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
