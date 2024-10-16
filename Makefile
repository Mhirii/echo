
SERVICES = auth user

dev: docker-run watch 

watch:
	@echo "Spinning up docker compose..."
	docker compose up pg -d
	@echo "Running Microservices..."
	@tmux new-session -d -s echo-air;
	@for service in $(SERVICES); do \
		echo "Starting $$service"; \
		tmux new-window -t echo-air -n $$service; \
		tmux send-keys -t echo-air:$$service "cd $$service && air"; \
		tmux send-keys -t echo-air:$$service C-m; \
	done
	@echo "Done"


build:
	@echo "Building Microservices..."
	@for service in $(SERVICES); do \
		echo "Building $$service"; \
		cd $$service && make build; \
		cd ..;\
	done
	@echo "Done"

docker-run:
	@if docker compose up -d 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up -d; \
	fi

docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

.PHONY: dev watch build docker-run docker-down
