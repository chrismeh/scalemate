.PHONY: build test

build:
	go build -o=./bin/scalemate-cli ./cmd/cli

test:
	go test -shuffle=on ./...
