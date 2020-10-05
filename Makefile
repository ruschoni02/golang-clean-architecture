.PHONY: all build run up deps

GOPATH ?= $(HOME)/go

all: build

build:
	go build -o golang-clean-architecture main.go

up:
	docker-compose up -d --force-recreate

mod:
	go mod download

deps: mod
	( cd /tmp; \
		go get github.com/cespare/reflex; \
		go get honnef.co/go/tools/cmd/staticcheck )

run:
	go run -race main.go

dev:
	$(GOPATH)/bin/reflex -s -r '\.go$$' make run

staticcheck:
	$(GOPATH)/bin/staticcheck ./...

test:
	go test ./core/... ./pkg/...
