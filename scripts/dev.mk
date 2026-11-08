# Go command.
GO ?= go

# Tool versions.
GCL_VERSION  := v2.15.0

# Project paths.
CMD_PATH    := ./cmd/mal-cover
BINARY_NAME := mal-cover
BINARY_PATH := $(CMD_PATH)/$(BINARY_NAME)

.PHONY: fmt clean lint install-linter build api

##@ Tools

fmt: ## Standarize go coding style for the whole project.
	@$(GO) fmt ./...

clean: ## Clean project binary, test, and coverage file.
	@$(GO) clean ./...
	@rm -f $(COVER_OUT)

lint: fmt ## Lint go source code.
	@golangci-lint run -D errcheck --timeout 5m

install-linter: ## Install golangci-lint.
	@curl -sSfL https://golangci-lint.run/install.sh | sh -s -- -b $(shell go env GOPATH)/bin $(GCL_VERSION)

##@ Build & Run

build: clean fmt ## Build the project executable binary. 
	@$(GO) build -o $(BINARY_PATH) -v $(CMD_PATH)

api: build ## Build and run API server.
	@$(BINARY_PATH) server