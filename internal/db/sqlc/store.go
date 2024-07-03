package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// store defines all the functions related to execute db queries
// and also contains code for transactions in DB
type Store interface {
	Querier
	DeleteBlogTx(ctx context.Context, arg DeleteBlogTxParams) (DeleteBlogTxResult, error)
}

// provides functionality for executing all SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// creating a NewStore
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
