# Include the .env/local.env file
include .env/local.env

# Export all variables
export

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root code-odessey

dropdb:
	docker exec -it postgres dropdb code-odessey

migrateup:
	migrate -path internal/db/migration -database "$(DB_SOURCE)" -verbose up

migratedown:
	migrate -path internal/db/migration -database "$(DB_SOURCE)" -verbose down

sqlc:
	sqlc generate

tidy:
	go mod tidy

test:
	go test -v --cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc tidy test