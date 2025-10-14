package configs

import (
	"fmt"
	"os"
)

// Config holds application configuration values.
type Config struct {
	Port string // Port on which the application runs.
	Env  string // Application env (e.g., development, production).
}

// Load returns a Config struct populated from environment variables.
func Load() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("APP_ENV", "local"),
	}
}

// HTTPAddr returns the HTTP address on which the application runs.
func (c Config) HTTPAddr() string {
	return fmt.Sprintf(":%s", c.Port)
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
