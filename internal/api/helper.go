package api

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teamkweku/code-odessey/config"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
	"github.com/teamkweku/code-odessey/internal/token"
	"github.com/teamkweku/code-odessey/pkg/utils"
)

// this package contains helper functions for gomock
func newTestServer(t *testing.T, store db.Store) *Server {
	config := config.Config{
		Environment:         "test",
		DBDriver:            "mock",
		DBSource:            "mock",
		HTTPServerAddress:   "0.0.0.0:8080",
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	require.NoError(t, err)

	server, err := NewServer(config, store)
	require.NoError(t, err)

	// Set the token maker for the server
	server.tokenMaker = tokenMaker

	return server
}

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration,
) {
	token, payload, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}
