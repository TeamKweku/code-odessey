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

migrateup1:
	migrate -path internal/db/migration -database "$(DB_SOURCE)" -verbose up 1

migratedown1:
	migrate -path internal/db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc:
	sqlc generate

tidy:
	go mod tidy

test:
	go test -v --cover -short ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

server:
	go run cmd/code-odessey/main.go 

build:
	docker compose -f docker/docker-compose.yaml up --build --remove-orphans

up:
	docker compose -f docker/docker-compose.yaml up

down:
	docker compose -f docker/docker-compose.yaml down

mock:
	mockgen -package mockdb -destination internal/db/mock/store.go github.com/teamkweku/code-odessey/internal/db/sqlc Store 
	mockgen -package mockwk -destination internal/worker/mock/distributor.go github.com/teamkweku/code-odessey/internal/worker TaskDistributor

db_docs:
	dbdocs build docs/db.dbml

db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

protolint:
	protolint lint --fix internal/proto/

proto:
	rm -f pb/*.go
	rm -f docs/swagger/*.swagger.json
	protoc --proto_path=internal/proto --go_out=internal/pb --go_opt=paths=source_relative \
	--grpc-gateway_out=internal/pb --grpc-gateway_opt=paths=source_relative \
    --go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
	--openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=code_odessey \
    internal/proto/*.proto
	statik -src=./docs/swagger -dest=./docs

evans:
	evans --host localhost --port 8085 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:alpine

.PHONY: postgres createdb dropdb migrateup migratedown sqlc tidy test mock migrateup1 migratedown1 down up build db_docs db_schema protolint proto evans redis
