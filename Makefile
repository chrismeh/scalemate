.PHONY: build test

build:
	go build -o=./bin/scalemate-cli ./cmd/cli
	go build -o=./bin/scalemate-web ./cmd/web

test:
	go test -shuffle=on ./...
