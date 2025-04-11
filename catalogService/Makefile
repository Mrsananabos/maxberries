GOOSE=goose
MIGRATIONS_DIR=./migrations
DB_DRIVER=postgres

migrate:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable $(GOOSE) up $(MIGRATIONS_DIR)

migrate_down:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable $(GOOSE) down $(MIGRATIONS_DIR)

create_migration:
	GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable $(GOOSE) create $(name) sql

.PHONY: migrate migrate_down create_migration
