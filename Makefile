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

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

server:
	go run cmd/code-odessey/main.go 

mock:
	mockgen -package mockdb -destination internal/db/mock/store.go github.com/teamkweku/code-odessey/internal/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc tidy test mock