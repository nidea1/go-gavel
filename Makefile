DOCKER_COMPOSE = docker compose -f compose/local/docker-compose.local.yml

# Declare phony targets (these are not files)
.PHONY: proto build-base up down stop help

# Proto generation
proto:
	@echo "Generating proto files..."
	@chmod +x ./scripts/generate_protos.sh
	@./scripts/generate_protos.sh

# Build base image
build-base:
	@echo "Building base image..."
	@docker build -t go-gavel-base:latest ./build/base

# Docker commands
up:
	@echo "Starting services..."
	$(DOCKER_COMPOSE) up -d --build

down:
	@echo "Stopping and removing services, networks, and volumes..."
	$(DOCKER_COMPOSE) down -v

stop:
	@echo "Stopping services..."
	$(DOCKER_COMPOSE) stop

# Help
help:
	@echo "Available commands:"
	@echo "  make proto       - Generate proto files"
	@echo "  make build-base  - Build base Docker image"
	@echo "  make up         - Start services"
	@echo "  make down       - Stop and remove services"
	@echo "  make stop       - Stop services (preserves containers)"
	@echo "  make help       - Show this help message"

.DEFAULT_GOAL := help
