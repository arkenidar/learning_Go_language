# Makefile for Learning Go Language project

.PHONY: help build run-server run-readline clean debug-server

# Default target
help:
	@echo "Available commands:"
	@echo "  make build        - Build all programs"
	@echo "  make run-server   - Run the HTTP server"
	@echo "  make run-readline - Run the read-line program"
	@echo "  make debug-server - Start HTTP server with debugging support"
	@echo "  make clean        - Clean built binaries"

# Build all programs
build:
	@echo "Building HTTP server..."
	@go build -o bin/http-server ./cmd/http-server
	@echo "Building read-line program..."
	@go build -o bin/read-line ./cmd/read-line
	@echo "All programs built successfully!"

# Run HTTP server
run-server:
	@echo "Starting HTTP server on :8080..."
	@go run ./cmd/http-server

# Run read-line program
run-readline:
	@echo "Starting read-line program..."
	@go run ./cmd/read-line

# Start HTTP server with debugging support (dlv)
debug-server:
	@echo "Starting HTTP server with debugging on :2345..."
	@echo "Connect your debugger to localhost:2345"
	@dlv debug ./cmd/http-server --headless --listen=:2345 --api-version=2 --accept-multiclient

# Clean built binaries
clean:
	@echo "Cleaning binaries..."
	@rm -rf bin/
	@rm -f http-server read-line
	@echo "Clean complete!"

# Create bin directory if it doesn't exist
bin:
	@mkdir -p bin
