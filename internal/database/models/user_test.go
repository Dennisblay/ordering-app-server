package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"ordering-server/util"
	"testing"
	"time"
)

func CreateRandomUser(t *testing.T) CreateUserRow {
	arg := CreateUserParams{
		FirstName:    util.RandomUser(),
		LastName:     util.RandomUser(),
		Email:        util.RandomEmail(),
		PasswordHash: util.RandomPassword(),
		Phone:        util.RandomPhone(),
		Address:      util.RandomAddress(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Phone, user.Phone)
	require.Equal(t, arg.Address, user.Address)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserById(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotZero(t, user2.ID)

	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.FirstName, user1.FirstName)
	require.Equal(t, user2.LastName, user1.LastName)
	require.Equal(t, user2.Email, user1.Email)
	require.Equal(t, user2.Address, user1.Address)

	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	arg := UpdateUserParams{
		ID:        user1.ID,
		FirstName: user1.FirstName,
		LastName:  user1.LastName,
		Email:     user1.Email,
		Phone:     user1.Phone,
		Address:   user1.Address,
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user2.ID, user1.ID)
	require.Equal(t, user2.FirstName, user1.FirstName)
	require.Equal(t, user2.LastName, user1.LastName)
	require.Equal(t, user2.Email, user1.Email)
	require.Equal(t, user2.Address, user1.Address)

	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)
}
