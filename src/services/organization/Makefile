MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := $(abspath $(MAKEFILE_DIR)../..)

PARENT_MODULE_NAME := github.com/mnes/haas
MODULE_NAME := github.com/mnes/haas/src/services/organization

####################################
# api
####################################
.PHONY: api/fmt
api/fmt: ## Format files
	@cd api; goimports -local '$(MODULE_NAME)' -w .
	@cd api; gofmt -w .
	@cd api; go mod tidy

.PHONY: proto/fmt
proto/fmt: ## Format proto files
	@docker run --platform=linux/amd64 --rm -v $(MAKEFILE_DIR)/proto:/app \
	-w /app \
	--entrypoint sh \
	$(PARENT_MODULE_NAME)/proto/fmt \
	-c 'find . -name "*.proto" | xargs clang-format -i'


####################################
# protoc
####################################
.PHONY: proto/gen
proto/gen: ## Generate interfaces of app domain
	@make proto/fmt
	@docker run --rm \
	-v $(MAKEFILE_DIR)proto:/proto \
	-v $(MAKEFILE_DIR)shell:/shell \
	-v $(MAKEFILE_DIR)api/pb:/api_pb \
	-w /proto \
	--entrypoint sh \
	$(PARENT_MODULE_NAME)/protoc \
	-c '/shell/protoc.sh'