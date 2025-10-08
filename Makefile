.PHONY: dev build run clean help

# Default target
dev: ## Start development server with hot reloading
	@~/go/bin/air -c air.toml

build: ## Build the application
	@go build -o ./.air/main ./src/cmd/server

run: build ## Build and run the application
	@./.air/main

clean: ## Clean build artifacts
	@rm -rf .air/
	@rm -f build-errors.log
	@echo "Cleaned build artifacts"

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Make dev the default target when running 'make' with no arguments
.DEFAULT_GOAL := dev
