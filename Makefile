.PHONY: all build run clean generate install-templ help

# Default target
all: generate build

# Generate Go code from Templ files
generate:
	@echo "Generating code from Templ files..."
	@$(shell go env GOPATH)/bin/templ generate

# Build the application
build: generate
	@echo "Building application..."
	@go build -o evbagel .

# Run the application
run: build
	@echo "Starting server..."
	@./evbagel

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f evbagel
	@rm -f components/*_templ.go

# Install Templ CLI
install-templ:
	@echo "Installing Templ CLI..."
	@go install github.com/a-h/templ/cmd/templ@latest

# Development mode - watch and rebuild
dev: generate
	@echo "Running in development mode..."
	@go run .

# Show help
help:
	@echo "EvBagel - Makefile commands:"
	@echo "  make all           - Generate Templ code and build"
	@echo "  make generate      - Generate Go code from Templ files"
	@echo "  make build         - Build the application"
	@echo "  make run           - Build and run the application"
	@echo "  make dev           - Run in development mode"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make install-templ - Install Templ CLI tool"
	@echo "  make help          - Show this help message"
