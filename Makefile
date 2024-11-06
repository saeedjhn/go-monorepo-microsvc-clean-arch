# Default Shell
export SHELL := bash

export VERSION := $(shell git describe --tags)
export COMMIT :=$(shell git rev-parse HEAD)
export BRANCH :=$(shell git rev-parse --abbrev-ref HEAD)

# Type of OS: Linux or Darwin.
export GOBASE := $(shell pwd)
export GOBUILDBASE := $(shell pwd)/build
export OSTYPE := $(shell uname -s | tr A-Z a-z)
# export ARCH := $(shell uname -m)
export ARCH = amd64
export PROJECTNAME := $(shell basename "$(PWD)")
export GOFILES := $(wildcard *.go)

export BINARY := superdo

export CMD := $(GOBASE)/cmd
export CLI := $(CMD)/cli/main.go
export HTTPSERVER := $(CMD)/httpserver/main.go
export PPROF := $(CMD)/pprof/main.go
export MIGRATION := $(CMD)/migrations/main.go
export SCHEDULER := $(CMD)/scheduler/main.go

# Setup the -ldflags option for go build here, interpolate the variable values
export LDFLAGS := -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out # -html, func, etc...

## development-up: Startup / Build services from docker-compose and air for live reloading
.PHONY: development-up
development-up:
	@echo
	@echo " > Startup / Build services from docker-compose and air for live reloading"
	@echo
	@docker-compose -f deployments/development/docker-compose.yaml up

# ==================================================================================== #
# PRODUCTION
# ==================================================================================== #
## production-up: Startup / Build services from docker-compose and air for live reloading
.PHONY: production-up
production-up:
	@echo
	@echo " > Startup / Build services from docker-compose"
	@echo
	@docker-compose -f deployments/production/docker-compose.yaml up

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
## gosec: The go security checker
.PHONY: gosec
gosec:
	@if ! command -v gosec &> /dev/null; then \
    	echo "gosec not found, installing..."; \
		go install github.com/securego/gosec/v2/cmd/gosec@latest; \
    else \
    	echo "gosec is already installed"; \
	fi
	gosec --version
	gosec ./...

## staticcheck: The advanced Go linter
.PHONY: staticcheck
staticcheck:
	@if ! command -v staticcheck &> /dev/null; then \
    	echo "staticcheck not found, installing..."; \
    	go install honnef.co/go/tools/cmd/staticcheck@latest; \
    else \
    	echo "staticcheck is already installed"; \
	fi
	staticcheck --version
	staticcheck ./...

## govulncheck: looks for vulnerabilities in Go programs using a specific build configuration. For analyzing source code
.PHONY: govulncheck
govulncheck:
	@if ! command -v govulncheck &> /dev/null; then \
    	echo "govulncheck not found, installing..."; \
		go install golang.org/x/vuln/cmd/govulncheck@latest ; \
    else \
    	echo "govulncheck is already installed"; \
	fi
	govulncheck --version
	govulncheck ./...

## golangci-lint: Smart, fast linters runner
.PHONY: golangci-lint
lint:
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found, installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		echo "golangci-lint is already installed"; \
	fi
	golangci-lint --version
	golangci-lint run --config .golangci.yml

## goimports: This tool updates your Go import lines, adding missing ones and removing unreferenced ones
.PHONY: goimports
goimports:
	goimports -w .
	@if ! command -v goimports &> /dev/null; then \
		echo "golangci-lint not found, installing..."; \
		go install golang.org/x/tools/cmd/goimports@latest; \
    else \
		echo "goimports is already installed"; \
    fi
	goimports -w .

## tidy: format code and tidy mod file
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod download
	go mod verify
	go vet ./...
	go fmt ./...
#	go install golang.org/x/tools/cmd/goimports@latest && goimports -w .
	goimports -w .
#	go install honnef.co/go/tools/cmd/staticcheck@latest && staticcheck -checks=all,-ST1000,-U1000 ./...
	staticcheck ./...
#	go install golang.org/x/vuln/cmd/govulncheck@latest && govulncheck ./...
	#govulncheck ./...
#	go install github.com/securego/gosec/v2/cmd/gosec@latest && gosec ./...
	gosec ./...
#	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && golangci-lint run --config .golangci.yml
	golangci-lint run --config .golangci.yml
	go test -v -race -buildvcs ./... # -vet=off: Nothing result

# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Usage a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
	@echo

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code
