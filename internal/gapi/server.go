package gapi

import (
	"fmt"

	"github.com/teamkweku/code-odessey/config"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
	"github.com/teamkweku/code-odessey/internal/pb"
	"github.com/teamkweku/code-odessey/internal/token"
)

// server to serve gRPC requeset for our banking service
type Server struct {
	pb.UnimplementedCodeOdesseyServer
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server setting configurations and
// RPC client calls
func NewServer(config config.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
