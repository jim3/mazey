APP_NAME := mazey
GO := go

.PHONY: help build run clean test fmt vet tidy

help: ## Show available make targets
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z_-]+:.*##/ {printf "  %-10s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the mazey binary
	$(GO) build -o $(APP_NAME) .

run: ## Run the CLI (pass args with ARGS="...")
	$(GO) run . $(ARGS)

clean: ## Remove built binary
	rm -f $(APP_NAME)

fmt: ## Format Go files
	$(GO) fmt ./...

tidy: ## Sync and clean module dependencies
	$(GO) mod tidy
