MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := $(abspath $(MAKEFILE_DIR)../..)

PARENT_MODULE_NAME := github.com/mnes/haas
MODULE_NAME := github.com/mnes/haas/src/services/imaging-study

####################################
# api
####################################
.PHONY: fmt
fmt: ## Format files
	@goimports -local '$(MODULE_NAME)' -w .
	@gofmt -w .
	@go mod tidy

.PHONY: build
build:
	@CGO_ENABLED=0 go build -v -o bin/api $(MODULE_NAME)/cmd

.PHONY: run
run:
	@go run cmd/*.go

.PHONY: lint
lint: ## Run linter
	@golangci-lint run