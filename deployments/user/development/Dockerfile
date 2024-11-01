# Stage 1: Build the application
FROM golang:1.23.2 AS builder

# Set environment variables for the project directory and Go modules
ENV PROJECT_DIR=/app \
    GO111MODULE=on\
    CGO_ENABLED=0

# Set the working directory to PROJECT_DIR
WORKDIR $PROJECT_DIR

# Copy Go module files to the working directory
COPY go.mod go.sum ./

# Download project dependencies
RUN go mod download

# Copy all project files to the working directory
COPY . .

# Install the Air tool for live reloading during development
RUN go install github.com/air-verse/air@latest

# Stage 2: Create the development image
FROM golang:1.23.2-alpine3.20 AS development

# Set environment variables for the project directory and Go modules
ENV PROJECT_DIR=/app \
    GO111MODULE=on\
    CGO_ENABLED=0

# Set the working directory to PROJECT_DIR
WORKDIR $PROJECT_DIR

# Copy the built project files from the builder stage to the development stage
COPY --from=builder $PROJECT_DIR $PROJECT_DIR

# Copy the Air binary from the builder stage to the development stage
COPY --from=builder /go/bin/air /usr/local/bin/air

# Specify the command to run when the container starts (execute Air with the specified config)
CMD ["air", "-c", "deployments/user/development/.air.toml"]