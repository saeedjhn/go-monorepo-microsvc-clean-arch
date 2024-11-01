package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/saeedjhn/go-monorepo-microsvc-clean-arch/internal/user/delivery/http"

	"github.com/saeedjhn/go-monorepo-microsvc-clean-arch/configs/user"
)

const _Service = "user" // Define the service name as a constant

func main() {
	var envMode string

	// Parse command-line flag for environment mode with default value as development
	flag.StringVar(&envMode, "env-mode", user.Development.String(), "config path, e.g., -env-mode dev")
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

	// Signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt) // more SIGX (SIGINT, SIGTERM, etc)

	// Start server
	server := http.New(cfg)

	go func() {
		if err = server.Run(); err != nil {
			log.Fatalln("Error running server: %w", err)
		}
	}()

	<-quit

	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, cfg.HTTPServer.GracefulShutdownTimeout)
	defer cancel()

	if err = server.Router.Shutdown(ctxWithTimeout); err != nil {
		log.Fatalf("Error shutdown server: %w", err)
	}

	<-ctxWithTimeout.Done()
}

// configPath constructs paths to .env files based on environment mode and service.
func configPath(workingDir, envMode, service string) []string {
	paths := make([]string, 0, 2) //nolint:mnd

	// Append paths for development environment
	if envMode == user.Development.String() {
		paths = append(paths,
			filepath.Join(workingDir, "deployments/development", ".env"),
			filepath.Join(workingDir, "deployments/", service, "development", ".env"),
		)

		return paths
	}

	// Append paths for production environment
	if envMode == user.Production.String() {
		paths = append(paths,
			filepath.Join(workingDir, "deployments/production", ".env"),
			filepath.Join(workingDir, "deployments", service, "production", ".env"),
		)

		return paths
	}

	log.Printf("Warning: Unsupported environment mode '%s'", envMode)

	return nil
}
