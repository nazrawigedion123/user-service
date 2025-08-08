DB_URL=postgresql://myuser:mypassword@localhost:5432/user-db?sslmode=disable

.PHONY: build run migrate sqlc clean

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

migrate:
	migrate -path ./internal/db/migrations -database $(DB_URL) up

DB_URL=postgresql://myuser:mypassword@localhost:5432/user-db?sslmode=disable

.PHONY: build run migrate sqlc clean

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

migrate:
	migrate -path ./internal/db/migrations -database $(DB_URL) up

sqlc:
	sqlc generate -f ./sql/sqlc.yaml


clean:
	rm -rf bin/
