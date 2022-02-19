package db

import (
	"Messaging/util"
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateMessage(t *testing.T) {
	acc1, acc2, mg, err := ExtendedCreateRandomMessageGroup()
	require.NoError(t, err)
	var identifier string
	if acc1.ID > acc2.ID {
		identifier = strconv.Itoa(int(acc1.ID)) + "-" + strconv.Itoa(int(acc2.ID))
	} else {
		identifier = strconv.Itoa(int(acc2.ID)) + "-" + strconv.Itoa(int(acc1.ID))

	}
	require.NotEmpty(t, mg)
	message := util.RandomString(15)
	params := createMessageParams{
		Group:    identifier,
		Message:  message,
		SentFrom: acc1.ID,
		SentTo:   acc2.ID,
	}
	msg, err := testQueries.createMessage(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, msg)
	require.Equal(t, msg.Group, identifier)
	require.Equal(t, msg.SentFrom, acc1.ID)
	require.Equal(t, msg.SentTo, acc2.ID)
	require.Equal(t, msg.Message, message)
	require.Equal(t, msg.Read, false)
}

func createRandomMessage() (Message, error) {
	acc1, acc2, _, err := ExtendedCreateRandomMessageGroup()
	if err != nil {
		return Message{}, err
	}
	var identifier string
	if acc1.ID > acc2.ID {
		identifier = strconv.Itoa(int(acc1.ID)) + "-" + strconv.Itoa(int(acc2.ID))
	} else {
		identifier = strconv.Itoa(int(acc2.ID)) + "-" + strconv.Itoa(int(acc1.ID))

	}
	message := util.RandomString(15)
	params := createMessageParams{
		Group:    identifier,
		Message:  message,
		SentFrom: acc1.ID,
		SentTo:   acc2.ID,
	}
	msg, err := testQueries.createMessage(context.Background(), params)
	return msg, err
}

func TestRetrieveAllMessages(t *testing.T) {
	acc1, acc2, _, err := ExtendedCreateRandomMessageGroup()
	require.NoError(t, err)
	var identifier string
	if acc1.ID > acc2.ID {
		identifier = strconv.Itoa(int(acc1.ID)) + "-" + strconv.Itoa(int(acc2.ID))
	} else {
		identifier = strconv.Itoa(int(acc2.ID)) + "-" + strconv.Itoa(int(acc1.ID))

	}
	for i := 0; i < 10; i++ {
		params := createMessageParams{
			Group:    identifier,
			Message:  util.RandomString(15),
			SentFrom: acc1.ID,
			SentTo:   acc2.ID,
		}
		_, err := testQueries.createMessage(context.Background(), params)
		require.NoError(t, err)
	}
	arg := RetrieveAllParams{
		Group:  identifier,
		Limit:  5,
		Offset: 5,
	}
	msgs, err := testQueries.RetrieveAll(context.Background(), arg)
	require.NoError(t, err)
	for i := 0; i < 5; i++ {
		require.Equal(t, msgs[i].Group, identifier)
		require.Equal(t, msgs[i].SentFrom, acc1.ID)
		require.Equal(t, msgs[i].SentTo, acc2.ID)
	}
}

func TestGetLatestUnreadMessage(t *testing.T) {
	msg, err := createRandomMessage()
	require.NoError(t, err)
	require.NotEmpty(t, msg)
	msg2, err := testQueries.GetLatestUnreadMessage(context.Background(), msg.SentTo)
	require.NoError(t, err)
	require.NotEmpty(t, msg2)
	require.Equal(t, msg, msg2)
}

func TestReadMessageGroup(t *testing.T) {
	msg, err := createRandomMessage()
	require.NoError(t, err)
	require.NotEmpty(t, msg)
	err = testQueries.ReadMessageGroup(context.Background(), msg.Group)
	require.NoError(t, err)
	args := RetrieveAllParams{
		Group:  msg.Group,
		Limit:  1,
		Offset: 0,
	}
	msgs, err := testQueries.RetrieveAll(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, msgs[0].Read, true)
}
