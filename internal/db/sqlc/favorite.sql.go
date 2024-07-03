// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: favorite.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

const createFavorite = `-- name: CreateFavorite :one
INSERT INTO favorites (
  blog_id
) VALUES (
  $1
) RETURNING id, blog_id, created_at, updated_at
`

func (q *Queries) CreateFavorite(ctx context.Context, blogID uuid.UUID) (Favorite, error) {
	row := q.db.QueryRow(ctx, createFavorite, blogID)
	var i Favorite
	err := row.Scan(
		&i.ID,
		&i.BlogID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteFavorite = `-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE id = $1
`

func (q *Queries) DeleteFavorite(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteFavorite, id)
	return err
}

const deleteFavoritesByBlog = `-- name: DeleteFavoritesByBlog :execresult
DELETE FROM favorites
WHERE blog_id = $1
`

func (q *Queries) DeleteFavoritesByBlog(ctx context.Context, blogID uuid.UUID) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, deleteFavoritesByBlog, blogID)
}

const getFavorite = `-- name: GetFavorite :one
SELECT id, blog_id, created_at, updated_at FROM favorites
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFavorite(ctx context.Context, id uuid.UUID) (Favorite, error) {
	row := q.db.QueryRow(ctx, getFavorite, id)
	var i Favorite
	err := row.Scan(
		&i.ID,
		&i.BlogID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listFavoritesByBlog = `-- name: ListFavoritesByBlog :many
SELECT id, blog_id, created_at, updated_at FROM favorites
WHERE blog_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type ListFavoritesByBlogParams struct {
	BlogID uuid.UUID `json:"blog_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListFavoritesByBlog(ctx context.Context, arg ListFavoritesByBlogParams) ([]Favorite, error) {
	rows, err := q.db.Query(ctx, listFavoritesByBlog, arg.BlogID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Favorite{}
	for rows.Next() {
		var i Favorite
		if err := rows.Scan(
			&i.ID,
			&i.BlogID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
