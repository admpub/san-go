.PHONY: all test release build

VERSION := $(shell cat version.go| grep "\sVersion\s=" | cut -d '"' -f2)

all: test build

test:
	go vet $(go list ./...)
	go test -v -race ./...

build:
	go build

release:
	git tag v$(VERSION)
	git push origin v$(VERSION)
