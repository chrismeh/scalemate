.PHONY: build test audit deploy ssh

build:
	go build -o=./bin/scalemate-cli ./cmd/cli
	go build -o=./bin/scalemate-web ./cmd/web

test:
	go test -shuffle=on -race -vet=off ./...

audit: test
	go fmt ./...
	go vet ./...

deploy: audit build
	rsync ./bin/scalemate-web ${SCALEMATE_USER}@${SCALEMATE_HOST}:~
	ssh -t ${SCALEMATE_USER}@${SCALEMATE_HOST} 'sudo service scalemate restart'

ssh:
	ssh ${SCALEMATE_USER}@${SCALEMATE_HOST}