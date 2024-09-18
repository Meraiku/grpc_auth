include .env

.ONESHELL:

build:
	@go build -o ./.bin/auth ./cmd/sso

run:build
	@./.bin/auth --config ./config/config_local.yaml

test:
	@go test ./... -cover -race -count=10

cover:
	@go test -short -count=1 -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out
	@rm coverage.out

up:
	@cd ./sql/migrations;
	@goose postgres $(DB_URL) up

down:
	@cd ./sql/migrations;
	@goose postgres $(DB_URL) down

reset:
	@cd ./sql/migrations;
	@goose postgres $(DB_URL) reset;
	@goose postgres $(DB_URL) up

docker:build
	docker build . -t $(I_PATH)/$(I_NAME)
	docker push $(I_PATH)/$(I_NAME)
	docker compose up