package postgresdb

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Gen1usBruh/MiniTwitter/util/random"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MediaLinks struct {
	Video []string `json:"video"`
	Photo []string `json:"photo"`
}

func TestCreateTweet(t *testing.T) {
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
	require.NotEmpty(t, tweet.ID)
	require.Equal(t, tweet.UserID, user.ID)
	require.Equal(t, tweet.Content, arg.Content)
	// require.Equal(t, tweet.Media, arg.Media)
	require.NotEmpty(t, tweet.CreatedAt)
	require.NotEmpty(t, tweet.UpdatedAt)
}

func TestSelectTweet(t *testing.T) {
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

	tweetSelect, err := testQueries.SelectTweet(context.Background(), tweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tweetSelect)
	require.Equal(t, tweetSelect.ID, tweet.ID)
	require.Equal(t, tweetSelect.Content, tweet.Content)
	require.Equal(t, tweetSelect.Media, tweet.Media)
	require.Equal(t, tweetSelect.UserID, tweet.UserID)
	require.Equal(t, tweetSelect.CreatedAt, tweet.CreatedAt)
	require.Equal(t, tweetSelect.UpdatedAt, tweet.UpdatedAt)
}

func TestDeleteTweet(t *testing.T) {
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

	err = testQueries.DeleteTweet(context.Background(), tweet.ID)
	require.NoError(t, err)

	tweet_select, err := testQueries.SelectTweet(context.Background(), tweet.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, tweet_select)
}
