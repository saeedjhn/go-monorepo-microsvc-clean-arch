package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/saeedjhn/go-monorepo-microsvc-clean-arch/configs/user"
)

const _Service = "user" // Define the service name as a constant

func main() {
	var envMode string

	// Parse command-line flag for environment mode with default value as development
	flag.StringVar(&envMode, "env-mode", user.Development.ToString(), "config path, e.g., -env-mode dev")
	flag.Parse()

	log.Println("Environment mode:", envMode)

	// Get the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	log.Println("Working Directory:", workingDir)

	// Load the configuration based on the environment and service name
	cfg, err := user.Load(configPath(workingDir, envMode, _Service))
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	log.Printf("Loaded Configuration: %#v", cfg)
}

// configPath constructs paths to .env files based on environment mode and service.
func configPath(workingDir, envMode, service string) []string {
	paths := make([]string, 0, 2) //nolint:mnd

	// Append paths for development environment
	if envMode == user.Development.ToString() {
		paths = append(paths,
			filepath.Join(workingDir, "deployments/development", ".env"),
			filepath.Join(workingDir, "deployments", service, "development", ".env"),
		)

		return paths
	}

	// Append paths for production environment
	if envMode == user.Production.ToString() {
		paths = append(paths,
			filepath.Join(workingDir, "deployments/production", ".env"),
			filepath.Join(workingDir, "deployments", service, "production", ".env"),
		)

		return paths
	}

	log.Printf("Warning: Unsupported environment mode '%s'", envMode)

	return nil
}
