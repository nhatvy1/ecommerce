GOOSE_DBSTRING ?= $(STR_MYSQL)"root:123456@tcp(127.0.0.1:8811)/user-service"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql

goose_up:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

goose_down:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down

goose_reset:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

docker_up:
	docker-compose up -d

sqlgen:
	sqlc generate

.PHONY: goose_up goose_down goose_reset sqlgen docker_up