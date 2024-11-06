package user

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

// Load loads environment variables from a list of file paths and binds them to a Config struct.
func Load(configPath string) (*Config, error) {
	var cfg Config

	if err := godotenv.Load(configPath); err != nil {
		// Return an error with context if unable to load a file
		return nil, fmt.Errorf("error loading file '%s': %w", configPath, err)
	}

	// Parse loaded environment variables into the Config struct
	if err := env.Parse(&cfg); err != nil {
		// Log and return error if binding to the struct fails
		return nil, fmt.Errorf("error parsing environment variables into struct: %w", err)
	}

	return &cfg, nil
}
