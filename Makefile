# Makefile for Learning Go Language project

.PHONY: help build run-server run-server-gorm run-readline run-tcp-chat clean debug-server debug-server-gorm debug-readline debug-tcp-chat deps tidy

# Default target
help:
	@echo "Available commands:"
	@echo "  make build           - Build all programs"
	@echo "  make run-server      - Run the HTTP server (in-memory)"
	@echo "  make run-server-gorm - Run the HTTP server with GORM database"
	@echo "  make run-readline    - Run the read-line program"
	@echo "  make run-tcp-chat    - Run the TCP chat server"
	@echo "  make debug-server    - Start HTTP server with debugging support"
	@echo "  make debug-server-gorm - Start GORM HTTP server with debugging support"
	@echo "  make debug-readline  - Start read-line program with debugging support"
	@echo "  make debug-tcp-chat  - Start TCP chat server with debugging support"
	@echo "  make deps            - Download and install dependencies"
	@echo "  make tidy            - Clean up go.mod and go.sum"
	@echo "  make clean           - Clean built binaries"

# Build all programs
build:
	@echo "Building HTTP server..."
	@go build -o bin/http-server ./cmd/http-server
	@echo "Building HTTP server with GORM..."
	@go build -o bin/http-server-gorm ./cmd/http-server-gorm
	@echo "Building read-line program..."
	@go build -o bin/read-line ./cmd/read-line
	@echo "Building TCP chat server..."
	@go build -o bin/tcp-chat-server ./cmd/tcp-chat-server
	@echo "All programs built successfully!"

# Run HTTP server
run-server:
	@echo "Starting HTTP server on :8080..."
	@go run ./cmd/http-server

# Run HTTP server with GORM
run-server-gorm:
	@echo "Starting HTTP server with GORM database on :8080..."
	@go run ./cmd/http-server-gorm

# Run read-line program
run-readline:
	@echo "Starting read-line program..."
	@go run ./cmd/read-line

# Run TCP chat server
run-tcp-chat:
	@echo "Starting TCP chat server on :8888..."
	@echo "Connect with: netcat localhost 8888"
	@go run ./cmd/tcp-chat-server

# Start HTTP server with debugging support (dlv)
debug-server:
	@echo "Starting HTTP server with debugging on :2345..."
	@echo "Connect your debugger to localhost:2345"
	@dlv debug ./cmd/http-server --headless --listen=:2345 --api-version=2 --accept-multiclient

# Start GORM HTTP server with debugging support (dlv)
debug-server-gorm:
	@echo "Starting GORM HTTP server with debugging on :2345..."
	@echo "Connect your debugger to localhost:2345"
	@dlv debug ./cmd/http-server-gorm --headless --listen=:2345 --api-version=2 --accept-multiclient

# Start read-line program with debugging support (dlv)
debug-readline:
	@echo "Starting read-line program with debugging on :2345..."
	@echo "Connect your debugger to localhost:2345"
	@dlv debug ./cmd/read-line --headless --listen=:2345 --api-version=2 --accept-multiclient

# Start TCP chat server with debugging support (dlv)
debug-tcp-chat:
	@echo "Starting TCP chat server with debugging on :2345..."
	@echo "Connect your debugger to localhost:2345"
	@dlv debug ./cmd/tcp-chat-server --headless --listen=:2345 --api-version=2 --accept-multiclient

# Clean built binaries
clean:
	@echo "Cleaning binaries..."
	@rm -rf bin/
	@rm -f http-server read-line
	@echo "Clean complete!"

# Download and install dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@echo "Dependencies downloaded!"

# Clean up go.mod and go.sum
tidy:
	@echo "Tidying go.mod..."
	@go mod tidy
	@echo "go.mod tidied!"

# Create bin directory if it doesn't exist
bin:
	@mkdir -p bin
