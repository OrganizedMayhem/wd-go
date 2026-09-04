.DEFAULT_GOAL := help

BINARY-DEV := wd-dev
BINARY-PROD := wd
VERSION := $(shell grep '^var version = ' cmd/version.go | sed 's/.*"\(.*\)".*/\1/')
LDFLAGS_PROD := -ldflags="-s -w -X github.com/OrganizedMayhem/wd/cmd.version=$(VERSION)"
GCFLAGS_DEV  := -gcflags="all=-N -l"

.PHONY: help dev prod clean-dev clean-prod test install

help: ## List available targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z0-9_-]+:.*?## / {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

dev: ## Build the development binary
	go build $(GCFLAGS_DEV) -o $(BINARY-DEV) .

prod: ## Build the production binary
	go build $(LDFLAGS_PROD) -o $(BINARY-PROD) .

clean-dev: ## Remove the development binary
	rm -f $(BINARY-DEV)

clean-prod: ## Remove the production binary
	rm -f $(BINARY-PROD)

test: ## Run tests across all packages
	go test ./...

install: prod ## Build and install the production binary
	go install $(LDFLAGS_PROD) .
