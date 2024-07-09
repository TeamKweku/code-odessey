# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY ../ .

RUN go build -o main /app/cmd/code-odessey
# RUN apk add curl
# RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/main .
# COPY --from=builder /app/migrate ./migrate
COPY ../.env/local.env /app/.env/
COPY start.sh .
COPY wait-for.sh .

# Make scripts executable
RUN chmod +x start.sh wait-for.sh

COPY ../internal/db/migration ./internal/db/migration

# List the contents of /app after copying the migration files
RUN ls -la /app

EXPOSE 8080
CMD ["./main"]
ENTRYPOINT [ "/app/start.sh" ]