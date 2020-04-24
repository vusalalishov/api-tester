.PHONY: all test fmt build

all: fmt test build

fmt:
	go fmt ./...

test:
	go test ./test/...

build:
	go build ./...
