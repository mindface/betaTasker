# Makefile for betaTasker

.PHONY: help dev prod build test clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development commands
dev: ## Start development environment
	docker-compose up -d

dev-logs: ## Show development logs
	docker-compose logs -f

dev-stop: ## Stop development environment
	docker-compose down

dev-clean: ## Clean development environment
	docker-compose down -v

# Production commands
prod: ## Start production environment
	docker-compose -f docker-compose.prod.yml up -d

prod-build: ## Build production images
	docker-compose -f docker-compose.prod.yml build

prod-logs: ## Show production logs
	docker-compose -f docker-compose.prod.yml logs -f

prod-stop: ## Stop production environment
	docker-compose -f docker-compose.prod.yml down

prod-clean: ## Clean production environment
	docker-compose -f docker-compose.prod.yml down -v

# Testing commands
test: ## Run all tests
	@echo "Running backend tests..."
	cd backer/godotask && go test ./... -v
	@echo "Running frontend tests..."
	cd fronter && yarn test

test-backend: ## Run backend tests
	cd backer/godotask && go test ./... -v

test-frontend: ## Run frontend tests
	cd fronter && yarn test

test-api: ## Run API tests
	./test_error_codes.sh

# Build commands
build: ## Build all Docker images
	docker-compose build

build-backend: ## Build backend image
	docker build -f backer/godotask/Dockerfile.prod -t betatasker-backend:latest backer/godotask

build-frontend: ## Build frontend image
	docker build -f fronter/Dockerfile.prod -t betatasker-frontend:latest fronter

# Database commands
db-migrate: ## Run database migrations
	docker exec backender go run /usr/local/go/godotask/seed/seed.go

db-shell: ## Access database shell
	docker exec -it db-prod psql -U dbgodotask -d dbgodotask

# Utility commands
clean: ## Clean all containers and volumes
	docker-compose down -v
	docker-compose -f docker-compose.prod.yml down -v
	docker system prune -f

logs: ## Show all logs
	docker-compose logs -f

status: ## Show container status
	docker-compose ps
	docker-compose -f docker-compose.prod.yml ps

shell-backend: ## Access backend shell
	docker exec -it backender /bin/bash

shell-frontend: ## Access frontend shell
	docker exec -it fronter /bin/sh

# Deployment commands
deploy: ## Deploy to production (requires configured server)
	@echo "Building production images..."
	$(MAKE) prod-build
	@echo "Pushing to registry..."
	docker-compose -f docker-compose.prod.yml push
	@echo "Deployment complete. Images are ready to be pulled on production server."