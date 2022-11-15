MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
ROOT_DIR := $(abspath $(MAKEFILE_DIR)../..)

PARENT_MODULE_NAME := github.com/mnes/haas
MODULE_NAME := github.com/mnes/haas/services/imaging-study

####################################
# api
####################################
.PHONY: api/fmt
api/fmt: ## Format files
	@cd api; goimports -local '$(MODULE_NAME)' -w .
	@cd api; gofmt -w .
	@cd api; go mod tidy

.PHONY: api/build
api/build:
	@cd api; CGO_ENABLED=0 go build -v -o bin/api $(MODULE_NAME)/cmd

.PHONY: api/run
api/run:
	@cd api; go run cmd/*.go


####################################
# openapi
####################################
.PHONY: openapi/gen
openapi/gen:
	@./shell/oapi-codegen.sh