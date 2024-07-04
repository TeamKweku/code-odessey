package db

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// Constants for PostgreSQL error codes
// ForeignKeyViolation represents a foreign key violation error
// UniqueViolation represents a unique constraint violation error
const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
)

// ErrRecordNotFound is returned when a query returns no rows
var ErrRecordNotFound = pgx.ErrNoRows

// ErrUniqueViolation is returned when a unique constraint is violated
var ErrUniqueViolation = &pgconn.PgError{
	Code: UniqueViolation,
}

// ErrorCode extracts the error code from a pgconn.PgError
// Function that returns the error code based on the error
// or nothing to use default error in code
func ErrorCode(err error) string {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code
	}
	return ""
}
