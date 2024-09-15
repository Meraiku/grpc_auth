include .env

.ONESHELL:

build:
	@go build -o ./.bin/auth ./cmd/sso

run:build
	@./.bin/auth --config ./config/config_local.yaml

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