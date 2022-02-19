package db

import (
	"Messaging/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	username := util.RandomString(6)
	acc, err := testQueries.createUser(context.Background(), username)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, username, acc.Username)
	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
}

func CreateRandomUser() (Users, error) {
	username := util.RandomString(6)
	return testQueries.createUser(context.Background(), username)
}

func TestGetUser(t *testing.T) {
	acc, err := CreateRandomUser()
	require.NoError(t, err)
	acc2, err2 := testQueries.GetUser(context.Background(), acc.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc, acc2)
}

func TestDeleteUser(t *testing.T) {
	acc, err := CreateRandomUser()
	require.NoError(t, err)
	err2 := testQueries.DeleteUser(context.Background(), acc.ID)
	require.NoError(t, err2)
	acc2, err3 := testQueries.GetUser(context.Background(), acc.ID)
	require.Error(t, err3)
	require.Equal(t, err3, sql.ErrNoRows)
	require.Empty(t, acc2)
}

func TestUpdateUser(t *testing.T) {
	acc, err := CreateRandomUser()
	require.NoError(t, err)
	newUsername := util.RandomString(5)
	updateParams := UpdateUserParams{
		ID:       acc.ID,
		Username: newUsername,
	}
	acc2, err2 := testQueries.UpdateUser(context.Background(), updateParams)
	require.NoError(t, err2)
	require.NotEmpty(t, acc2)
	require.NotEqual(t, acc, acc2)
	require.Equal(t, newUsername, acc2.Username)
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUser()
	}
	params := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	accs, err := testQueries.ListUsers(context.Background(), params)
	require.NoError(t, err)
	require.Len(t, accs, 5)
	for i := 0; i < 5; i++ {
		require.NotEmpty(t, accs[i])
	}
}
