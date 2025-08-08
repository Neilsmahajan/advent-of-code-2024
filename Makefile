.PHONY: build run test clean day1 day2 all

# Default target
build:
	go build -o bin/aoc ./cmd/aoc/main.go

# Run specific day and part
run:
	@if [ -z "$(DAY)" ]; then \
		echo "Usage: make run DAY=<day> [PART=<part>]"; \
		echo "Example: make run DAY=1 PART=1"; \
		exit 1; \
	fi
	@PART=$${PART:-1}; \
	go run ./cmd/aoc/main.go -day=$(DAY) -part=$$PART

# Quick shortcuts for common days
day1:
	go run ./cmd/aoc/main.go -day=1 -part=1
	go run ./cmd/aoc/main.go -day=1 -part=2

day2:
	go run ./cmd/aoc/main.go -day=2 -part=1
	go run ./cmd/aoc/main.go -day=2 -part=2

# Run all implemented solutions
all:
	go run ./cmd/aoc/main.go -all

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Check for updates to dependencies
check-updates:
	go list -u -m all

# Help
help:
	@echo "Available targets:"
	@echo "  build      - Build the project binary"
	@echo "  run        - Run specific day: make run DAY=1 PART=1"
	@echo "  day1       - Run both parts of day 1"
	@echo "  day2       - Run both parts of day 2"
	@echo "  all        - Run all implemented solutions"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  deps       - Install/update dependencies"
	@echo "  fmt        - Format code"
	@echo "  lint       - Run linter"
	@echo "  help       - Show this help"
