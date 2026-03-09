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
	golangci-lint run 

install-dev:
	@echo "==> Installing development dependencies..."
	curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.11.2
	golangci-lint --version