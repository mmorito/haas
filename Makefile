MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := $(abspath $(MAKEFILE_DIR)..)

MODULE_NAME := github.com/mnes/haas

.PHONY: proto/setup
proto/setup: ## Do the setup needed for development
	@make docker/proto/fmt
	@make docker/proto/protoc

.PHONY: docker/proto/fmt
docker/proto/fmt: ## Build a docker image to run the proto file format
	@docker build --platform=linux/amd64 -t $(MODULE_NAME)/proto/fmt -f docker/fmt.Dockerfile .

.PHONY: docker/proto/protoc
docker/proto/protoc: ## Build a docker image to generate interfaces
	@docker build --platform=linux/x86_64 -t $(MODULE_NAME)/protoc -f docker/protoc.Dockerfile .
