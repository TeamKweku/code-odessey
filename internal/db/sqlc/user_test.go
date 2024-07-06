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

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashedPassword(utils.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       utils.RandomUsername(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.True(t, user.UpdatedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {

	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}

func TestGetNonExistentUser(t *testing.T) {
	nonExistentID := uuid.New()

	user, err := testQueries.GetUser(context.Background(), nonExistentID)
	require.Error(t, err)
	require.Empty(t, user)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
}

// Test Update User
func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID: user1.ID,
		Username: pgtype.Text{
			String: utils.RandomUsername(),
			Valid:  true,
		},
		FullName: pgtype.Text{
			String: utils.RandomOwner(),
			Valid:  true,
		},
		Email: pgtype.Text{
			String: utils.RandomEmail(),
			Valid:  true,
		},
		HashedPassword: pgtype.Text{
			String: "new_hashed_password",
			Valid:  true,
		},
		IsEmailVerified: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
		PasswordChangedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, arg.ID, updatedUser.ID)
	require.Equal(t, arg.Username.String, updatedUser.Username)
	require.Equal(t, arg.FullName.String, updatedUser.FullName)
	require.Equal(t, arg.Email.String, updatedUser.Email)
	require.Equal(t, arg.HashedPassword.String, updatedUser.HashedPassword)
	require.Equal(t, arg.IsEmailVerified.Bool, updatedUser.IsEmailVerified)
	require.WithinDuration(t, arg.PasswordChangedAt.Time, updatedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, updatedUser.CreatedAt, time.Second)
	require.NotEqual(t, user1.UpdatedAt, updatedUser.UpdatedAt)
}
