.PHONY: build test

build:
	go build -o=./bin/scalemate ./cmd/scalemate

test:
	go test -shuffle=on ./...
