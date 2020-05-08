.PHONY: all test fmt build

all: fmt test build rapit

fmt:
	go fmt ./...

test:
	go test ./test/...

build:
	go build ./...

rapit:
	go build -o rapit ./core/cmd
