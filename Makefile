MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := $(abspath $(MAKEFILE_DIR)..)

MODULE_NAME := github.com/mnes/haas
# TOOLS := $(shell cat api/tools/tools.go | egrep '^\s_ ' | awk '{ print $$2 }')

####################################
# API
####################################
# .PHONY: api/setup
# api/setup:
# 	-cd api; go mod download golang.org/x/tools
# 	-cd api; go install -tags tools -mod=readonly $(TOOLS)

# .PHONY: api/fmt
# api/fmt: ## Format files
# 	@cd api; goimports -local '$(MODULE_NAME)' -w .
# 	@cd api; gofmt -w .
# 	@cd api; go mod tidy

####################################
# Protocol Buffers
####################################
.PHONY: proto/setup
proto/setup: ## Do the setup needed for development
	@make docker/proto/fmt
	@make docker/proto/protoc

.PHONY: docker/proto/fmt
docker/proto/fmt: ## Build a docker image to run the proto file format
	@docker build --platform=linux/amd64 -t $(MODULE_NAME)/proto/fmt -f docker/fmt.Dockerfile .

.PHONY: docker/proto/protoc
docker/proto/protoc: ## Build a docker image to generate interfaces
	@docker build --platform=linux/amd64 -t $(MODULE_NAME)/protoc -f docker/protoc.Dockerfile . --no-cache

.PHONY: proto/fmt
proto/fmt: ## Format proto files
	@docker run --platform=linux/amd64 --rm -v $(MAKEFILE_DIR)/services:/app \
	-w /app \
	--entrypoint sh \
	$(MODULE_NAME)/proto/fmt \
	-c 'find . -name "*.proto" | xargs clang-format -i'

.PHONY: proto/gen
proto/gen: ## Generate interfaces of app domain
	@make proto/fmt
	@docker run --rm \
	-v $(MAKEFILE_DIR)services:/services \
	-v $(MAKEFILE_DIR)shell:/shell \
	-v $(MAKEFILE_DIR)doc:/doc \
	-w /services \
	--entrypoint sh \
	$(MODULE_NAME)/protoc \
	-c '../shell/protoc.sh'
	# @make api/fmt

.PHONY: openapi/gen
openapi/gen:
	@./shell/oapi-codegen.sh