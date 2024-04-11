SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build


build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: ## Start application
	$(GOCMD) run cmd/main.go


wire: ## Generate wire_gen.go
	cd pkg/di && wire
swag: ## Generate swagger docs
	swag init -g cmd/main.go -o ./cmd/docs

build:
	go build -o cmd/build cmd/main.go 
