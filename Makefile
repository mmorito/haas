MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := $(abspath $(MAKEFILE_DIR)..)

MODULE_NAME := github.com/mnes/haas

.PHONY: proto/setup
proto/setup: ## Do the setup needed for development
	@make docker/fmt
	@make docker/protoc

.PHONY: docker/proto/fmt
docker/proto/fmt: ## Build a docker image to run the proto file format
	@docker build --platform=linux/amd64 -t $(MODULE_NAME)/proto/fmt -f docker/fmt.Dockerfile ./proto

.PHONY: docker/proto/protoc
docker/proto/protoc: ## Build a docker image to generate interfaces
	@docker build -t $(MODULE_NAME)/protoc -f docker/protoc.Dockerfile ./proto

.PHONY: proto/fmt
proto/fmt: ## Format proto files
	@docker run --platform=linux/amd64 --rm -v $(MAKEFILE_DIR)/proto:/app \
	-w /app \
	--entrypoint sh \
	$(MODULE_NAME)/proto/fmt \
	-c 'find . -name "*.proto" | xargs clang-format -i'

.PHONY: proto/gen
proto/gen: ## Generate interfaces of app domain
	@make proto/fmt
	@docker run --rm \
	-v $(MAKEFILE_DIR)proto:/proto \
	-v $(MAKEFILE_DIR)api/pb:/api_pb \
	-v $(MAKEFILE_DIR)swagger:/swagger \
	-v $(MAKEFILE_DIR)shell:/shell \
	-v $(MAKEFILE_DIR)doc:/doc \
	-w /proto \
	--entrypoint sh \
	$(MODULE_NAME)/protoc \
	-c '../shell/protoc.sh'