// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type Querier interface {
	CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateFavorite(ctx context.Context, arg CreateFavoriteParams) (Favorite, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteBlog(ctx context.Context, id uuid.UUID) error
	DeleteComment(ctx context.Context, id uuid.UUID) error
	DeleteCommentByBlogID(ctx context.Context, arg DeleteCommentByBlogIDParams) error
	DeleteCommentsByBlog(ctx context.Context, blogID uuid.UUID) (pgconn.CommandTag, error)
	DeleteFavorite(ctx context.Context, id uuid.UUID) error
	DeleteFavoritesByBlog(ctx context.Context, blogID uuid.UUID) (pgconn.CommandTag, error)
	GetBlog(ctx context.Context, id uuid.UUID) (Blog, error)
	GetBlogBySlug(ctx context.Context, slug string) (Blog, error)
	GetComment(ctx context.Context, id uuid.UUID) (Comment, error)
	GetFavorite(ctx context.Context, id uuid.UUID) (Favorite, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	ListBlogs(ctx context.Context, arg ListBlogsParams) ([]Blog, error)
	ListCommentsByBlog(ctx context.Context, arg ListCommentsByBlogParams) ([]Comment, error)
	ListFavoritesByBlog(ctx context.Context, arg ListFavoritesByBlogParams) ([]Favorite, error)
	UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error)
	UpdateCommentByBlogID(ctx context.Context, arg UpdateCommentByBlogIDParams) (Comment, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
