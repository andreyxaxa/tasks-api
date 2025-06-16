BASE_STACK = docker compose -f docker-compose.yml

compose-up: ### Run docker compose
		$(BASE_STACK) up --build -d
.PHONY: compose-up

compose-down: ### Down docker compose
		$(BASE_STACK) down
.PHONY: compose-down

deps: ### deps tidy + verify
		go mod tidy && go mod verify
.PHONY: deps

run: deps ### Run
		go mod download && \
		CGO_ENABLED=0 go run ./cmd/app
.PHONY: run