package db

import (
	"context"
	"testing"
	"time"

	"github.com/samudra-ajri/simple-user/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	name := util.RandomName()

	user, err := testQueries.CreateUser(context.Background(), name)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, name, user.Name)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestDisplayUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.DisplayUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Name, user2.Name)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestDisplayAllUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}

	arg := DisplayAllUsersParams{
		Limit:  5,
		Offset: 0,
	}

	users, err := testQueries.DisplayAllUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
