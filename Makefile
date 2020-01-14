PROJECT_NAME := $(CI_PROJECT_NAME)
PROJECT_NAMESPACE := $(CI_PROJECT_NAMESPACE)
PROJECT_PATH := "$(PROJECT_NAMESPACE)/$(PROJECT_NAME)"
PACKAGE_NAME := "gitlab.com/$(PROJECT_PATH)"
PACKAGE_LIST := $(shell go list $(PACKAGE_NAME)/... | grep -v /vendor/)

.PHONY: all lint test race coverage tidy

all: tidy test race lint

lint: ## Lint the files
	@golangci-lint run ./...

test: ## Run unittests
	@go test -short $(PACKAGE_NAME)/...

race: tidy ## Run data race detector
	@go test -race -short ${PACKAGE_LIST}

coverage: ## Generate global code coverage report
	@go test -cover -v -coverpkg=$(PACKAGE_NAME)/... ${PACKAGE_LIST}

tidy: ## Get the dependencies
	@go mod tidy

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'