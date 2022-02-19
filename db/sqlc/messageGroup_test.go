package db

import (
	"context"
	"database/sql"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateMessageGroup(t *testing.T) {
	acc1, err := CreateRandomUser()
	require.NoError(t, err)
	acc2, err := CreateRandomUser()
	require.NoError(t, err)
	var identifier string
	if acc1.ID > acc2.ID {
		identifier = strconv.Itoa(int(acc1.ID)) + "-" + strconv.Itoa(int(acc2.ID))
	} else {
		identifier = strconv.Itoa(int(acc2.ID)) + "-" + strconv.Itoa(int(acc1.ID))

	}
	params := CreateMessageGroupParams{
		Identifier: identifier,
		FromUser:   acc1.ID,
		ToUser:     acc2.ID,
	}
	mg, err := testQueries.CreateMessageGroup(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, mg)
	require.Equal(t, identifier, mg.Identifier)
	require.Equal(t, acc1.ID, mg.FromUser)
	require.Equal(t, acc2.ID, mg.ToUser)
}

func CreateRandomMessageGroup() (MessageGroup, error) {
	acc1, err := CreateRandomUser()
	if err != nil {
		return MessageGroup{}, err
	}
	acc2, err := CreateRandomUser()
	if err != nil {
		return MessageGroup{}, err
	}
	var identifier string
	if acc1.ID > acc2.ID {
		identifier = strconv.Itoa(int(acc1.ID)) + "-" + strconv.Itoa(int(acc2.ID))
	} else {
		identifier = strconv.Itoa(int(acc2.ID)) + "-" + strconv.Itoa(int(acc1.ID))

	}
	params := CreateMessageGroupParams{
		Identifier: identifier,
		FromUser:   acc1.ID,
		ToUser:     acc2.ID,
	}
	return testQueries.CreateMessageGroup(context.Background(), params)
}

func ExtendedCreateRandomMessageGroup() (Users, Users, MessageGroup, error) {
	acc1, err := CreateRandomUser()
	if err != nil {
		return Users{}, Users{}, MessageGroup{}, err
	}
	acc2, err := CreateRandomUser()
	if err != nil {
		return Users{}, Users{}, MessageGroup{}, err

	}
	var identifier string
	if acc1.ID > acc2.ID {
		identifier = strconv.Itoa(int(acc1.ID)) + "-" + strconv.Itoa(int(acc2.ID))
	} else {
		identifier = strconv.Itoa(int(acc2.ID)) + "-" + strconv.Itoa(int(acc1.ID))

	}
	params := CreateMessageGroupParams{
		Identifier: identifier,
		FromUser:   acc1.ID,
		ToUser:     acc2.ID,
	}
	msg, err := testQueries.CreateMessageGroup(context.Background(), params)
	return acc1, acc2, msg, err
}

func TestGetMessageGroup(t *testing.T) {
	acc, err := CreateRandomMessageGroup()
	require.NoError(t, err)
	acc2, err2 := testQueries.GetMessageGroup(context.Background(), acc.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc, acc2)
}

func TestDeleteMessageGroup(t *testing.T) {
	acc, err := CreateRandomMessageGroup()
	require.NoError(t, err)
	err2 := testQueries.DeleteMessageGroup(context.Background(), acc.ID)
	require.NoError(t, err2)
	acc2, err3 := testQueries.GetMessageGroup(context.Background(), acc.ID)
	require.Error(t, err3)
	require.Equal(t, err3, sql.ErrNoRows)
	require.Empty(t, acc2)
}

func TestListMessageGroup(t *testing.T) {
	acc1, acc2, msg, err := ExtendedCreateRandomMessageGroup()
	require.NoError(t, err)
	require.NotEmpty(t, msg)
	accs, err := testQueries.ListGroups(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.Contains(t, accs, msg)
	accs2, err := testQueries.ListGroups(context.Background(), acc2.ID)
	require.NoError(t, err)
	require.Contains(t, accs2, msg)
}

func TestMessageGroupFromIdentifier(t *testing.T) {
	acc1, acc2, mg, err := ExtendedCreateRandomMessageGroup()
	require.NoError(t, err)
	var identifier string
	if acc1.ID > acc2.ID {
		identifier = strconv.Itoa(int(acc1.ID)) + "-" + strconv.Itoa(int(acc2.ID))
	} else {
		identifier = strconv.Itoa(int(acc2.ID)) + "-" + strconv.Itoa(int(acc1.ID))

	}
	require.NotEmpty(t, mg)
	msg, err := testQueries.GetMessageGroupFromIdentifier(context.Background(), identifier)
	require.NoError(t, err)
	require.Equal(t, mg, msg)
}
