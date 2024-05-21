GO_VERSION := 1.22.2

APP_IMAGE := avito_app
DB_IMAGE := postgres:16

APP_CONTAINER := app
DB_CONTAINER := database

DB_CREDENTIALS := "postgresql://boss:pass@127.0.0.1:5431/avito_db?sslmode=disable"
GOOSE := goose -dir db/migrations postgres $(DB_CREDENTIALS)


build:
	docker compose down
	docker build -t $(APP_IMAGE):1 ./
	docker compose up


fillDB:
	$(GOOSE) up

dropDB:
	$(GOOSE) reset