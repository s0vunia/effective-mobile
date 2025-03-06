LOCAL_BIN:=$(CURDIR)/bin

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/hey@latest

local-docker-compose-up:
	sudo docker compose up --build -d

local-docker-compose-down:
	sudo docker compose down

swagger-gen:
	swag init -g internal/api/songs/service.go -o api/swagger

