.PHONY: all
all: dep build test lint/golang

.PHONY: dep
dep:
	go mod download

.PHONY: build
build:
	go build ./

.PHONY: test
test:
	go test -race -coverprofile coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: lint/golang
lint/golang:
	golangci-lint run -v