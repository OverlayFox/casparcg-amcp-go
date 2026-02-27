.DEFAULT_GOAL := all

.PHONY: all build test lint tidy verify clean

all: tidy verify lint build test

build:
	@echo "==> Building..."
	go build -v ./...

test:
	@echo "==> Running tests with race detector..."
	go test -race -v ./...

lint:
	@echo "==> Running golangci-lint..."
	# Note: You must have golangci-lint installed locally for this to work
	golangci-lint run