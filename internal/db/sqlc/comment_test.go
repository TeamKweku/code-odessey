package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
	"github.com/teamkweku/code-odessey/pkg/utils"
)

func createRandomComment(t *testing.T, blog Blog) Comment {
	arg := CreateCommentParams{
		BlogID: blog.ID,
		Body:   utils.RandomDescription(),
	}

	comment, err := testQueries.CreateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, arg.BlogID, comment.BlogID)
	require.Equal(t, arg.Body, comment.Body)

	require.NotZero(t, comment.ID)
	require.NotZero(t, comment.CreatedAt)

	return comment
}

func TestCreateEntry(t *testing.T) {
	blog := createRandomBlog(t)
	createRandomComment(t, blog)
}

func TestCommentEntry(t *testing.T) {
	blog := createRandomBlog(t)
	comment1 := createRandomComment(t, blog)

	comment2, err := testQueries.GetComment(context.Background(), comment1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, comment1.ID, comment2.ID)
	require.Equal(t, comment1.BlogID, comment2.BlogID)
	require.Equal(t, comment1.Body, comment2.Body)
	require.WithinDuration(t, comment1.CreatedAt, comment2.CreatedAt, time.Second)
	require.True(t, comment1.UpdatedAt.IsZero())
}

// Testing GetComment function
func TestGetCommentNotFound(t *testing.T) {
	nonExistentID := uuid.New()

	comment, err := testQueries.GetComment(context.Background(), nonExistentID)
	require.Error(t, err)
	require.Empty(t, comment)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
}

// Testing DeleteBlog function
func TestDeleteComment(t *testing.T) {
	blog := createRandomBlog(t)

	comment1 := createRandomComment(t, blog)

	err := testQueries.DeleteComment(context.Background(), comment1.ID)
	require.NoError(t, err)

	comment2, err := testQueries.GetComment(context.Background(), comment1.ID)
	require.EqualError(t, err, pgx.ErrNoRows.Error())

	require.Empty(t, comment2)
}

func TestListCommentsByBlog(t *testing.T) {
	blog := createRandomBlog(t)

	for i := 0; i < 10; i++ {
		createRandomComment(t, blog)
	}

	arg := ListCommentsByBlogParams{
		BlogID: blog.ID,
		Limit:  5,
		Offset: 5,
	}

	comments, err := testQueries.ListCommentsByBlog(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, comments, 5)

	for _, comment := range comments {
		require.NotEmpty(t, comment)
		require.Equal(t, arg.BlogID, comment.BlogID)
	}
}

func TestUpdateComment(t *testing.T) {
	blog := createRandomBlog(t)
	comment1 := createRandomComment(t, blog)

	arg := UpdateCommentParams{
		ID:   comment1.ID,
		Body: utils.RandomContent(),
	}

	comment2, err := testQueries.UpdateComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	// Verify that the blog was updated
	require.Equal(t, comment1.ID, comment2.ID)
	require.NotEqual(t, comment1.Body, comment2.Body)
	require.Equal(t, comment1.BlogID, comment2.BlogID)

}
