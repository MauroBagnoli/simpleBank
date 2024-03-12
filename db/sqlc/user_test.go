package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"tutorial.sqlc.dev/app/util"
)

func createRandomUser(t *testing.T) (User, error) {
	hashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.Time.IsZero())

	require.NotZero(t, user.CreatedAt)

	return user, err
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1, _ := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user1.Email)
	require.Equal(t, user1.HashedPassword, user1.HashedPassword)
	require.Equal(t, user1.FullName, user1.FullName)

	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}
