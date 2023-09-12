# Base Go commands.
GO_CMD     := go
GO_FMT     := $(GO_CMD) fmt
GO_GET     := $(GO_CMD) get
GO_MOD     := $(GO_CMD) mod
GO_CLEAN   := $(GO_CMD) clean
GO_BUILD   := $(GO_CMD) build

# Project executable file, and its binary.
CMD_PATH    := ./cmd/mal-cover
BINARY_NAME := mal-cover

# Base golangci-lint commands.
GCL_CMD := golangci-lint
GCL_RUN := $(GCL_CMD) run

# Default makefile target.
.DEFAULT_GOAL := run

# Standarize go coding style for the whole project.
.PHONY: fmt
fmt:
	@$(GO_FMT) ./...

# Lint go source code.
.PHONY: lint
lint: fmt
	@$(GCL_RUN) -D errcheck --timeout 5m

# Clean project binary, test, and coverage file.
.PHONY: clean
clean:
	@$(GO_CLEAN) ./...

# Install library.
.PHONY: install
install:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.46.2
	@$(GCL_CMD) version

# Build the project executable binary.
.PHONY: build
build: clean fmt
	@cd $(CMD_PATH); \
	$(GO_BUILD) -o $(BINARY_NAME) -v .

# Build and run the binary.
.PHONY: run
run: build
	@cd $(CMD_PATH); \
	./$(BINARY_NAME) server

# Docker base command.
DOCKER_CMD   := docker
DOCKER_IMAGE := $(DOCKER_CMD) image

# Docker-compose base command and docker-compose.yml path.
COMPOSE_CMD   := docker-compose
COMPOSE_BUILD := deployment/build.yml
COMPOSE_API   := deployment/api.yml

# Build docker images and container for the project
# then delete builder image.
.PHONY: docker-build
docker-build:
	@$(COMPOSE_CMD) -f $(COMPOSE_BUILD) build
	@$(DOCKER_IMAGE) prune -f --filter label=stage=mal_cover_builder

# Start docker container.
.PHONY: docker
docker:
	@$(COMPOSE_CMD) -f $(COMPOSE_API) up
