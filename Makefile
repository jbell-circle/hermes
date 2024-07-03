.PHONY: build test

run: build
	./bin/hermes

build:
	go build -o ./bin/hermes ./cmd/hermes.go

test: build
	go test ./...
