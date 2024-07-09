package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/teamkweku/code-odessey/config"
	"github.com/teamkweku/code-odessey/internal/api"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
	"github.com/teamkweku/code-odessey/internal/gapi"
	"github.com/teamkweku/code-odessey/internal/pb"
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

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(connPool)

	// running gRPC api gateway concurrently with gRPC server
	go runGatewayServer(config, store)

	// Start the gRPCServer
	runGrpcServer(config, store)
}

// running db migration on main file execution
func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

// Function to start gRPC server
func runGrpcServer(config config.Config, store db.Store) {
	// create new gapi server
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	// create a new gRPC server, with logger
	grpcServer := grpc.NewServer()

	pb.RegisterCodeOdesseyServer(grpcServer, server)

	// register a gRPC reflection to enable client to know
	// which grpc service are available on server and call them
	reflection.Register(grpcServer)

	// creating the server to serve gRPC
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot create listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Msg("cannot start gRPC server")
	}
}

// setting up gRPC gateway server using the inprocess translation
// method
func runGatewayServer(config config.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pb.RegisterCodeOdesseyHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Msg("cannot create listener")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())
	handler := gapi.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Msg("cannot start HTTP gateway server")
	}

}

// function to start a gin server
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
