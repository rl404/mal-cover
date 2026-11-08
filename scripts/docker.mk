# Compose command.
COMPOSE ?= docker compose

# Compose files helper.
compose = $(COMPOSE) -f deployment/compose.$(1).yml -p mal-cover-$(2)

.PHONY: docker-lint docker-build docker-api docker-stop

##@ Docker Tools

docker-lint: ## Run linter in docker.
	@$(call compose,lint,lint) --profile lint up
	@$(call compose,lint,lint) --profile lint rm -f

##@ Docker Build & Run

docker-build: ## Build docker image then prune the builder image.
	@$(call compose,build,build) --profile build build 
	@docker image prune -f --filter label=stage=mal_cover_builder

docker-api: ## Start API container.
	@$(call compose,api,api) up -d
	@$(call compose,api,api) logs --follow --tail 20

docker-stop: ## Stop running container.
	@$(call compose,api,api) stop
	@$(call compose,api,api) rm -f
