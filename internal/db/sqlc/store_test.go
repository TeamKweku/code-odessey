package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeleteBlogTx(t *testing.T) {
	store := NewStore(testDB)

	// Create a random blog, comment, and favorite for testing
	blog := createRandomBlog(t)
	comment := createRandomComment(t, blog)
	favorite := createRandomFavorite(t, blog)

	// Delete the blog transaction
	arg := DeleteBlogTxParams{
		ID: blog.ID,
	}

	result, err := store.DeleteBlogTx(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	// Check that the blog ID matches
	require.Equal(t, blog.ID, result.DeletedBlogID)

	// Check that the number of deleted comments and favorites is as expected
	require.Equal(t, int64(1), result.DeletedCommentsCount)
	require.Equal(t, int64(1), result.DeletedFavoritesCount)

	// Verify that the blog, comment, and favorite have been deleted
	_, err = store.GetBlog(context.Background(), blog.ID)
	require.Error(t, err)

	_, err = store.GetComment(context.Background(), comment.ID)
	require.Error(t, err)

	_, err = store.GetFavorite(context.Background(), favorite.ID)
	require.Error(t, err)
}
