.PHONY: build test audit deploy ssh

build:
	env GOARCH=amd64 GOOS=linux go build -o=./bin/scalemate-cli ./cmd/cli
	env GOARCH=amd64 GOOS=linux go build -o=./bin/scalemate-web ./cmd/web

test:
	go test -shuffle=on -race -vet=off ./...

audit: test
	go fmt ./...
	go vet ./...

deploy: audit build
	rsync ./bin/scalemate-web scalemate:~
	ssh -t scalemate 'sudo service scalemate restart'

ssh:
	ssh scalemate