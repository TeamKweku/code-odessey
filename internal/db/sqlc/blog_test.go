package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/teamkweku/code-odessey/pkg/utils"
)

func createRandomBlog(t *testing.T) Blog {
	// Capture the current time
	now := time.Now().UTC()

	arg := CreateBlogParams{
		Title:       utils.RandomTitle(),
		Slug:        utils.RandomSlug() + "-" + uuid.New().String(),
		Description: utils.RandomDescription(),
		BannerImage: utils.RandomImageURL(),
		Body:        utils.RandomParagraph(),
	}

	blog, err := testQueries.CreateBlog(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, blog)

	require.Equal(t, arg.Title, blog.Title)
	require.Equal(t, arg.Slug, blog.Slug)
	require.Equal(t, arg.Description, blog.Description)
	require.Equal(t, arg.BannerImage, blog.BannerImage)

	require.NotZero(t, blog.ID)
	require.NotZero(t, blog.CreatedAt)
	require.True(t, blog.UpdatedAt.IsZero())

	// Check that CreatedAt is set to a recent timestamp
	require.WithinDuration(t, now, blog.CreatedAt.UTC(), 2*time.Second)

	require.Equal(t, blog.UpdatedAt.UTC(), time.Time{}.UTC())

	return blog
}

func TestCreateBlog(t *testing.T) {
	createRandomBlog(t)
}

func TestGetBlogbyID(t *testing.T) {
	blog1 := createRandomBlog(t)

	// check for the created account in database
	blog2, err := testQueries.GetBlog(context.Background(), blog1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, blog2)

	require.Equal(t, blog1.ID, blog2.ID)
	require.Equal(t, blog1.Title, blog2.Title)
	require.Equal(t, blog1.Slug, blog2.Slug)
	require.Equal(t, blog1.Description, blog2.Description)
	require.Equal(t, blog1.BannerImage, blog2.BannerImage)
	require.Equal(t, blog1.Body, blog2.Body)
	require.True(t, blog2.UpdatedAt.IsZero())

	require.WithinDuration(t, blog1.CreatedAt, blog2.CreatedAt, time.Second)
}

func TestGetBlogbySlug(t *testing.T) {
	blog1 := createRandomBlog(t)

	// check for the created account in database
	blog2, err := testQueries.GetBlogBySlug(context.Background(), blog1.Slug)

	require.NoError(t, err)
	require.NotEmpty(t, blog2)

	require.Equal(t, blog1.Slug, blog2.Slug)
	require.Equal(t, blog1.Title, blog2.Title)
	require.Equal(t, blog1.ID, blog2.ID)
	require.Equal(t, blog1.Description, blog2.Description)
	require.Equal(t, blog1.BannerImage, blog2.BannerImage)
	require.Equal(t, blog1.Body, blog2.Body)
	require.True(t, blog2.UpdatedAt.IsZero())

	require.WithinDuration(t, blog1.CreatedAt, blog2.CreatedAt, time.Second)
}

func TestGetBlogBySlugNotFound(t *testing.T) {
	nonExistentSlug := "non-existent-slug"

	blog, err := testQueries.GetBlogBySlug(context.Background(), nonExistentSlug)

	require.Error(t, err)
	require.Empty(t, blog)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
}

func TestGetBlogNotFound(t *testing.T) {
	nonExistentID := uuid.New()

	blog, err := testQueries.GetBlog(context.Background(), nonExistentID)

	require.Error(t, err)
	require.Empty(t, blog)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
}

// Test for timeout
func TestGetBlogTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := testQueries.GetBlog(ctx, uuid.New())
	require.Error(t, err)
}

func TestDeleteBlog(t *testing.T) {
	blog1 := createRandomBlog(t)

	err := testQueries.DeleteBlog(context.Background(), blog1.ID)
	require.NoError(t, err)

	blog2, err := testQueries.GetBlog(context.Background(), blog1.ID)
	require.EqualError(t, err, pgx.ErrNoRows.Error())

	require.Empty(t, blog2)
}

func TestDeleteBlogNotFound(t *testing.T) {
	nonExistentID := uuid.New()

	err := testQueries.DeleteBlog(context.Background(), nonExistentID)
	require.NoError(t, err)

	// Attempting to delete the same non-existent blog again should still succeed without error
	err = testQueries.DeleteBlog(context.Background(), nonExistentID)
	require.NoError(t, err)
}

func TestListBlogs(t *testing.T) {
	// var lastBlog Blog
	for i := 0; i < 10; i++ {
		createRandomBlog(t)
	}

	arg := ListBlogsParams{
		Limit:  5,
		Offset: 0,
	}

	blogs, err := testQueries.ListBlogs(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, blogs)
	require.Len(t, blogs, 5)

	arg.Offset = 5
	blogs, err = testQueries.ListBlogs(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, blogs)
	require.Len(t, blogs, 5)

	for _, blog := range blogs {
		require.NotEmpty(t, blog)
	}
}

// Testing the UpdateBlog with all fields provided
func TestUpdateBlogAllFields(t *testing.T) {
	oldBlog := createRandomBlog(t)

	// check if updated_at is its zero value
	require.Equal(t, oldBlog.UpdatedAt.UTC(), time.Time{}.UTC())

	newTitle := utils.RandomTitle()
	newSlug := utils.RandomSlug() + "-" + uuid.New().String()
	newDescription := utils.RandomDescription()
	newBanner := utils.RandomImageURL()
	newBody := utils.RandomParagraph()

	updatedBlog, err := testQueries.UpdateBlog(context.Background(), UpdateBlogParams{
		ID: oldBlog.ID,
		Title: pgtype.Text{
			String: newTitle,
			Valid:  true,
		},
		Description: pgtype.Text{
			String: newDescription,
			Valid:  true,
		},
		Slug: pgtype.Text{
			String: newSlug,
			Valid:  true,
		},
		BannerImage: pgtype.Text{
			String: newBanner,
			Valid:  true,
		},
		Body: pgtype.Text{
			String: newBody,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldBlog.Title, updatedBlog.Title)
	require.NotEqual(t, oldBlog.Slug, updatedBlog.Slug)
	require.NotEqual(t, oldBlog.BannerImage, updatedBlog.BannerImage)
	require.NotEqual(t, oldBlog.Description, updatedBlog.Description)
	require.NotEqual(t, oldBlog.Body, updatedBlog.Body)

	require.WithinDuration(t, time.Now().UTC(), updatedBlog.UpdatedAt.UTC(), 2*time.Second)
}

func TestUpdateBlogWithOnlyTitle(t *testing.T) {
	oldBlog := createRandomBlog(t)

	newTitle := utils.RandomTitle()
	updatedBlog, err := testQueries.UpdateBlog(context.Background(), UpdateBlogParams{
		ID: oldBlog.ID,
		Title: pgtype.Text{
			String: newTitle,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldBlog.Title, updatedBlog.Title)
	require.Equal(t, oldBlog.Slug, updatedBlog.Slug)
	require.Equal(t, oldBlog.BannerImage, updatedBlog.BannerImage)
	require.Equal(t, oldBlog.Description, updatedBlog.Description)
	require.Equal(t, oldBlog.Body, updatedBlog.Body)

}

func TestUpdateBlogWithOnlyBody(t *testing.T) {
	oldBlog := createRandomBlog(t)

	newBody := utils.RandomParagraph()
	updatedBlog, err := testQueries.UpdateBlog(context.Background(), UpdateBlogParams{
		ID: oldBlog.ID,
		Body: pgtype.Text{
			String: newBody,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldBlog.Body, updatedBlog.Body)
	require.Equal(t, oldBlog.Slug, updatedBlog.Slug)
	require.Equal(t, oldBlog.BannerImage, updatedBlog.BannerImage)
	require.Equal(t, oldBlog.Description, updatedBlog.Description)
	require.Equal(t, oldBlog.Title, updatedBlog.Title)
}

func TestUpdateBlogWithOnlyBanner(t *testing.T) {
	oldBlog := createRandomBlog(t)

	newBanner := utils.RandomImageURL()
	updatedBlog, err := testQueries.UpdateBlog(context.Background(), UpdateBlogParams{
		ID: oldBlog.ID,
		BannerImage: pgtype.Text{
			String: newBanner,
			Valid:  true,
		},
	})

	require.NoError(t, err)
	require.NotEqual(t, oldBlog.BannerImage, updatedBlog.BannerImage)
	require.Equal(t, oldBlog.Slug, updatedBlog.Slug)
	require.Equal(t, oldBlog.Title, updatedBlog.Title)
	require.Equal(t, oldBlog.Description, updatedBlog.Description)
	require.Equal(t, oldBlog.Body, updatedBlog.Body)

}
