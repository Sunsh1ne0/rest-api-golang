.PHONY: build

build:
	go build -v ./cmd/apiserver

.DEFAULT_GOAL := build

.PHONY: test
test: 
	go test -v ./internal/app/apiserver