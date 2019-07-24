PROJECT_NAMESPACE := $(CI_PROJECT_NAMESPACE)
PROJECT_NAME := $(CI_PROJECT_NAME)
PROJECT_PATH := "$(PROJECT_NAMESPACE)/$(PROJECT_NAME)"
PACKAGE_NAME := "gitlab.com/$(PROJECT_PATH)"
PACKAGE_PATH := "$(GOPATH)/src/$(PACKAGE_NAME)"
PACKAGE_LIST := $(shell go list $(PACKAGE_NAME)/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all lint test rase coverage dep

all: dep test race lint

lint: ## Lint the files
	@golangci-lint run ./...

test: ## Run unittests
	@go test -short $(PACKAGE_NAME)/...

race: dep ## Run data race detector
	@go test -race -short ${PACKAGE_LIST}

coverage: ## Generate global code coverage report
	@go test -cover -v -coverpkg=$(PACKAGE_NAME)/... ${PACKAGE_LIST}

dep: ## Get the dependencies
	@go get -v -d -t $(PACKAGE_NAME)/...

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
