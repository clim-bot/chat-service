# Variables
DOCKER_COMPOSE = docker-compose
APP_CONTAINER = app

# Targets
.PHONY: all build run stop logs

all: build run

build:
	@echo "Building Docker containers..."
	$(DOCKER_COMPOSE) build

run:
	@echo "Running Docker containers..."
	$(DOCKER_COMPOSE) up -d

stop:
	@echo "Stopping Docker containers..."
	$(DOCKER_COMPOSE) down

logs:
	@echo "Fetching logs..."
	$(DOCKER_COMPOSE) logs -f $(APP_CONTAINER)
