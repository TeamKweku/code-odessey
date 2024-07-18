package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "github.com/teamkweku/code-odessey/docs/statik"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/teamkweku/code-odessey/config"
	"github.com/teamkweku/code-odessey/internal/api"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
	"github.com/teamkweku/code-odessey/internal/gapi"
	"github.com/teamkweku/code-odessey/internal/mail"
	"github.com/teamkweku/code-odessey/internal/pb"
	"github.com/teamkweku/code-odessey/internal/worker"
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

	// connecting to Redis
	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	// run task processor
	go runTaskProcessor(config, &redisOpt, store)

	// running gRPC api gateway concurrently with gRPC server
	go runGatewayServer(config, store, taskDistributor)

	// Start the gRPCServer
	runGrpcServer(config, store, taskDistributor)
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

// send verification email
func runTaskProcessor(config config.Config, redisOpt *asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(*redisOpt, store, mailer)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

// Function to start gRPC server
func runGrpcServer(config config.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	// create new gapi server
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	// creating a gRPC logger using interceptors
	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)

	// create a new gRPC server, with logger
	grpcServer := grpc.NewServer(grpcLogger)

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
func runGatewayServer(config config.Config, store db.Store, taskDistributor worker.TaskDistributor) {
	server, err := gapi.NewServer(config, store, taskDistributor)
	if err != nil {
		log.Fatal().Msg("cannot create server")
	}

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(jsonOption)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pb.RegisterCodeOdesseyHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	// serving static files (swagger dist files)
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik fs")
	}

	mux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msgf("Received request for: %s", r.URL.Path)

		// If the path is /swagger/ or /swagger, serve index.html
		if r.URL.Path == "/swagger/" || r.URL.Path == "/swagger" {
			content, err := statikFS.Open("/index.html")
			if err != nil {
				log.Error().Err(err).Msg("Failed to open index.html")
				http.NotFound(w, r)
				return
			}
			defer content.Close()
			http.ServeContent(w, r, "index.html", time.Now(), content)
			return
		}

		// For other files, try to serve them directly
		path := strings.TrimPrefix(r.URL.Path, "/swagger")
		content, err := statikFS.Open(path)
		if err != nil {
			log.Error().Err(err).Msgf("Failed to open file: %s", path)
			http.NotFound(w, r)
			return
		}
		defer content.Close()
		http.ServeContent(w, r, path, time.Now(), content)
	})

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
