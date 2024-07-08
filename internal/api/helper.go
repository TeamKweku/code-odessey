package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teamkweku/code-odessey/config"
	db "github.com/teamkweku/code-odessey/internal/db/sqlc"
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

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
