package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/teamkweku/code-odessey/config"
	"github.com/teamkweku/code-odessey/internal/api"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
)

func main() {
	config, err := config.LoadConfig(".env/")
	if err != nil {
		log.Fatal().Msg("cannot load config:")
	}

	// check for environment and enable pretty logging for
	// local development only
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// connect the the database and open a connection pool
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := db.NewStore(connPool)

	// Start the GinServer
	runGinServer(config, store)
}

func runGinServer(config config.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot start server")
	}
}
