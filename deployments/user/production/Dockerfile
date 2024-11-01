# Stage 1: Build the application
FROM golang:1.23.2 AS builder

# Set environment variable for the project directory
ENV PROJECT_DIR=/app \
    GO111MODULE=on\
    CGO_ENABLED=0

# Set the working directory to PROJECT_DIR
WORKDIR $PROJECT_DIR

# Copy Go module files to download dependencies
COPY go.mod go.sum ./

# Download project dependencies
RUN go mod download

# Copy all project files to the working directory
COPY . .

# Build the binary and output to build/user
RUN go build -o build/user/main cmd/user/main.go

# Stage 2: Create the final production image
FROM golang:1.23.2-alpine3.20 AS production

# Re-define the environment variable for the project directory
ENV PROJECT_DIR=/app \
    GO111MODULE=on\
    CGO_ENABLED=0

# Set the working directory to PROJECT_DIR
WORKDIR $PROJECT_DIR

# (commented out) Copy all files from the builder stage to the production stage
#COPY --from=builder $PROJECT_DIR $PROJECT_DIR

# Copy the build directory from the builder stage to the production stage
COPY --from=builder $PROJECT_DIR/build $PROJECT_DIR/build

# Copy the deployments directory from the builder stage to the production stage
COPY --from=builder $PROJECT_DIR/deployments $PROJECT_DIR/deployments

# Specify the command to run when the container starts (execute the built binary)
CMD ["build/user/main", "-env-mode", "production"]

