package postgresdb

import (
	"context"
	"encoding/json"
	// "fmt"
	"testing"

	"github.com/Gen1usBruh/MiniTwitter/util/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateRandomComment(t *testing.T) Comment {
	user := createRandomAccount(t)
	arg := CreateTweetParams{}
	arg.UserID = user.ID
	arg.Content = random.GenerateRandomAlphanumeric(64)
	media := MediaLinks{
		Video: []string{random.GenerateRandomAlphanumeric(16)},
		Photo: []string{random.GenerateRandomAlphanumeric(16),
			random.GenerateRandomAlphanumeric(16)},
	}
	var err error
	arg.Media, err = json.Marshal(media)
	assert.NoError(t, err, "No error marshaling JSON for the sake of test")
	tweet, err := testQueries.CreateTweet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)

	argComm := CreateCommentParams{}
	argComm.UserID = user.ID
	argComm.TweetID.Int32 = tweet.ID
	argComm.TweetID.Valid = true
	argComm.RetweetID.Valid = false
	argComm.ParentCommentID.Valid = false
	argComm.PostType = "tweet"
	argComm.Content = random.GenerateRandomAlphanumeric(64)
	argComm.Media = arg.Media

	comm, err := testQueries.CreateComment(context.Background(), argComm)
	require.NoError(t, err)
	require.NotEmpty(t, comm)
	require.NotEmpty(t, comm.ID)
	require.NotEmpty(t, comm.CreatedAt)
	require.NotEmpty(t, comm.UpdatedAt)
	require.Equal(t, argComm.Content, comm.Content)
	// require.Equal(t, argComm.Media, comm.Media)
	require.Equal(t, argComm.UserID, comm.UserID)
	require.Equal(t, argComm.TweetID, comm.TweetID)
	require.Equal(t, argComm.PostType, comm.PostType)

	return comm
}

func TestCreateComment(t *testing.T) {
	CreateRandomComment(t)
}

func TestSelectComment(t *testing.T) {
	user := createRandomAccount(t)
	arg := CreateTweetParams{}
	arg.UserID = user.ID
	arg.Content = random.GenerateRandomAlphanumeric(64)
	media := MediaLinks{
		Video: []string{random.GenerateRandomAlphanumeric(16)},
		Photo: []string{random.GenerateRandomAlphanumeric(16),
			random.GenerateRandomAlphanumeric(16)},
	}
	var err error
	arg.Media, err = json.Marshal(media)
	assert.NoError(t, err, "No error marshaling JSON for the sake of test")
	tweet, err := testQueries.CreateTweet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)

	argComm := CreateCommentParams{}
	argComm.UserID = user.ID
	argComm.TweetID.Int32 = tweet.ID
	argComm.TweetID.Valid = true
	argComm.RetweetID.Valid = false
	argComm.ParentCommentID.Valid = false
	argComm.PostType = "tweet"
	argComm.Content = random.GenerateRandomAlphanumeric(64)
	argComm.Media = arg.Media

	comm, err := testQueries.CreateComment(context.Background(), argComm)
	require.NoError(t, err)
	require.NotEmpty(t, comm)

	argSelect := SelectCommentParams{}
	argSelect.TweetID.Int32 = tweet.ID
	argSelect.TweetID.Valid = true
	argSelect.Column2 = string("tweet")

	// fmt.Println("Old: ", argSelect.TweetID)

	comm_arr, err := testQueries.SelectComment(context.Background(), argSelect)
	require.NoError(t, err)
	require.NotEmpty(t, comm_arr)
	require.Equal(t, len(comm_arr), 1)
	require.NotEmpty(t, comm_arr[0].ID)
	require.Equal(t, comm_arr[0].Content, argComm.Content)
	// require.Equal(t, comm_arr[0].Media, argComm.Media)
	require.NotEmpty(t, comm_arr[0].Username)
	require.NotEmpty(t, comm_arr[0].UserID, argComm.UserID)
	require.Equal(t, comm_arr[0].CreatedAt, comm.CreatedAt)
}

func TestDeleteComment(t *testing.T) {
	comm := CreateRandomComment(t)

	arg := DeleteCommentParams{}
	arg.ID = comm.ID
	arg.UserID = comm.UserID

	err := testQueries.DeleteComment(context.Background(), arg)
	require.NoError(t, err)
}
