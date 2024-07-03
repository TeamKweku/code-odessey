package db

import (
	"context"

	"github.com/google/uuid"
)

// DeleteBlogTxParams contains the input parameters of the delete blog transaction
type DeleteBlogTxParams struct {
	ID uuid.UUID `json:"id"`
}

// DeleteBlogTxResult is the result of the delete blog transaction
type DeleteBlogTxResult struct {
	DeletedBlogID         uuid.UUID `json:"deleted_blog_id"`
	DeletedCommentsCount  int64     `json:"deleted_comments_count"`
	DeletedFavoritesCount int64     `json:"deleted_favorites_count"`
}

// DeleteBlogTx performs a deletion of a blog and all its associated data.
// It deletes the blog, its comments, and its favorites within a database transaction
func (store *SQLStore) DeleteBlogTx(ctx context.Context, arg DeleteBlogTxParams) (DeleteBlogTxResult, error) {
	var result DeleteBlogTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Delete favorites
		deletedFavorites, err := q.DeleteFavoritesByBlog(ctx, arg.ID)
		if err != nil {
			return err
		}
		result.DeletedFavoritesCount = int64(deletedFavorites.RowsAffected())

		// Delete comments
		deletedComments, err := q.DeleteCommentsByBlog(ctx, arg.ID)
		if err != nil {
			return err
		}
		result.DeletedCommentsCount = int64(deletedComments.RowsAffected())

		// Delete the blog
		err = q.DeleteBlog(ctx, arg.ID)
		if err != nil {
			return err
		}
		result.DeletedBlogID = arg.ID

		return nil
	})

	return result, err
}
