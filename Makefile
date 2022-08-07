.PHONY: all
all: dep build test

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
