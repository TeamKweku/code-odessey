package db

import (
	"context"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomFavorite(t *testing.T, blog Blog) Favorite {
	favorite, err := testQueries.CreateFavorite(context.Background(), blog.ID)
	require.NoError(t, err)
	require.NotEmpty(t, favorite)

	require.NotEqual(t, uuid.Nil, favorite.ID)
	require.Equal(t, blog.ID, favorite.BlogID)
	require.WithinDuration(t, time.Now(), favorite.CreatedAt, 2*time.Second)
	require.True(t, favorite.UpdatedAt.IsZero())

	return favorite
}
func TestCreateFavorite(t *testing.T) {
	blog := createRandomBlog(t)
	createRandomFavorite(t, blog)
}

func TestGetFavorite(t *testing.T) {
	blog := createRandomBlog(t)
	favorite1 := createRandomFavorite(t, blog)

	favorite2, err := testQueries.GetFavorite(context.Background(), favorite1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, favorite2)

	require.Equal(t, favorite1.ID, favorite2.ID)
	require.Equal(t, favorite1.BlogID, favorite2.BlogID)
	require.WithinDuration(t, favorite1.CreatedAt, favorite2.CreatedAt, time.Second)
	require.Equal(t, favorite1.UpdatedAt, favorite2.UpdatedAt)
}

func TestListFavoritesByBlog(t *testing.T) {
	blog := createRandomBlog(t)

	for i := 0; i < 10; i++ {
		createRandomFavorite(t, blog)
	}

	arg := ListFavoritesByBlogParams{
		BlogID: blog.ID,
		Limit:  5,
		Offset: 0,
	}

	favorites, err := testQueries.ListFavoritesByBlog(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, favorites, 5)

	for _, favorite := range favorites {
		require.NotEmpty(t, favorite)
		require.Equal(t, blog.ID, favorite.BlogID)
	}
}

func TestDeleteFavorite(t *testing.T) {
	blog := createRandomBlog(t)
	favorite := createRandomFavorite(t, blog)

	err := testQueries.DeleteFavorite(context.Background(), favorite.ID)
	require.NoError(t, err)

	deletedFavorite, err := testQueries.GetFavorite(context.Background(), favorite.ID)
	require.Error(t, err)
	require.Empty(t, deletedFavorite)
}
