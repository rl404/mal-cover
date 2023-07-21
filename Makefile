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

# Default makefile target.
.DEFAULT_GOAL := run

# Standarize go coding style for the whole project.
.PHONY: fmt
fmt:
	@$(GO_FMT) ./...

# Lint go source code.
.PHONY: lint
lint: fmt
	@golint `go list ./... | grep -v /vendor/`

# Clean project binary, test, and coverage file.
.PHONY: clean
clean:
	@$(GO_CLEAN) ./...

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
COMPOSE_CMD  := docker-compose
COMPOSE_FILE := docker-compose.yml

# Build docker images and container for the project
# then delete builder image.
.PHONY: docker-build
docker-build:
	@$(COMPOSE_CMD) -f $(COMPOSE_FILE) build
	@$(DOCKER_IMAGE) prune -f --filter label=stage=mal_cover_builder

# Start built docker containers for api.
.PHONY: docker-api
docker-api:
	@$(COMPOSE_CMD) -f $(COMPOSE_FILE) -p mal_cover-api up -d
	@$(COMPOSE_CMD) -f $(COMPOSE_FILE) -p mal_cover-api logs --follow --tail 20
