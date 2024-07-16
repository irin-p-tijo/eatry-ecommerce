SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build


build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR)/app -v ./cmd


${BINARY_DIR}:
	mkdir -p ${BINARY_DIR}

run: ## Start application
	$(GOCMD) run cmd/main.go


wire: ## Generate wire_gen.go
	cd pkg/di && wire

swag: ## Generate swagger docs
	swag init -g cmd/main.go -o ./cmd/docs



mock:
	
	mockgen -source=pkg/usecase/interfaces/user.go -destination=pkg/usecase/mock/user_mock.go -package=mock
	mockgen -source=pkg/usecase/interfaces/admin.go -destination=pkg/usecase/mock/admin_mock.go -package=mock
	mockgen -source=pkg/repository/interfaces/admin.go -destination=pkg/repository/mock/admin_mock.go -package=mock
	mockgen -source=pkg/repository/interfaces/user.go -destination=pkg/repository/mock/user_mock.go -package=mock
	mockgen -source=pkg/repository/interfaces/cart.go -destination=pkg/repository/mock/cart_mock.go -package=mock
	
	
test: ##Run testing
	go test ./...