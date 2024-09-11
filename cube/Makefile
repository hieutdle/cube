# Makefile

.PHONY: test

test: 
	@echo "Running tests for all folders..."
	@go test ./...
	@echo "All tests passed successfully!"

.PHONY: format

format:
	@echo "Formatting all files..."
	@go fmt ./...
	@echo "All files formatted successfully!"

.PHONY: usage

usage:
	@echo "Usage: make test"
	@echo "       make format"