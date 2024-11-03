package postgresdb

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Gen1usBruh/MiniTwitter/util/random"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func CreateRandomUsersTweetAndRetweet(t *testing.T) Retweet {
	user_orig := createRandomAccount(t)
	arg := CreateTweetParams{}
	arg.UserID = user_orig.ID
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

	user_new := createRandomAccount(t)
	arg_tweet := CreateRetweetParams{}
	arg_tweet.UserID = user_new.ID
	arg_tweet.ParentTweetID.Int32 = tweet.ID
	arg_tweet.ParentTweetID.Valid = true
	arg_tweet.ParentRetweetID.Valid = false
	arg_tweet.Quote.Scan(random.GenerateRandomAlphanumeric(64))
	arg_tweet.IsQuote = true

	retweet, err := testQueries.CreateRetweet(context.Background(), arg_tweet)
	require.NoError(t, err)
	require.NotEmpty(t, retweet)
	require.NotEmpty(t, retweet.ID)
	require.Equal(t, retweet.UserID, arg_tweet.UserID)
	require.Equal(t, retweet.ParentTweetID, arg_tweet.ParentTweetID)
	require.Equal(t, retweet.ParentRetweetID, arg_tweet.ParentRetweetID)
	require.Equal(t, retweet.Quote, arg_tweet.Quote)
	require.Equal(t, retweet.IsQuote, arg_tweet.IsQuote)
	require.NotEmpty(t, retweet.CreatedAt)
	require.NotEmpty(t, retweet.UpdatedAt)

	return retweet
}

func TestCreateRetweet(t *testing.T) {
	CreateRandomUsersTweetAndRetweet(t)
}

func TestSelectRetweet(t *testing.T) {
	retweet := CreateRandomUsersTweetAndRetweet(t)

	retweet_select, err := testQueries.SelectRetweet(context.Background(), retweet.ID)
	require.NoError(t, err)
	require.NotEmpty(t, retweet_select)
	require.Equal(t, retweet.ID, retweet_select.ID)
	require.Equal(t, retweet.UserID, retweet_select.UserID)
	require.Equal(t, retweet.ParentTweetID, retweet_select.ParentTweetID)
	require.Equal(t, retweet.ParentRetweetID, retweet_select.ParentRetweetID)
	require.Equal(t, retweet.Quote, retweet_select.Quote)
	require.Equal(t, retweet.CreatedAt, retweet_select.CreatedAt)
	require.Equal(t, retweet.UpdatedAt, retweet_select.UpdatedAt)
}

func TestDeleteRetweet(t *testing.T) {
	retweet := CreateRandomUsersTweetAndRetweet(t)

	err := testQueries.DeleteRetweet(context.Background(), retweet.ID)
	require.NoError(t, err)

	retweet_select, err := testQueries.SelectRetweet(context.Background(), retweet.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, retweet_select)
}

func TestSelectRetweetsOfTweet(t *testing.T) {
	user_orig := createRandomAccount(t)
	arg := CreateTweetParams{}
	arg.UserID = user_orig.ID
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

	user_new := createRandomAccount(t)

	arg_tweet := CreateRetweetParams{}
	arg_tweet.UserID = user_new.ID
	arg_tweet.ParentTweetID.Int32 = tweet.ID
	arg_tweet.ParentTweetID.Valid = true
	arg_tweet.ParentRetweetID.Valid = false
	arg_tweet.Quote.Scan(random.GenerateRandomAlphanumeric(64))
	arg_tweet.IsQuote = true
	retweet, err := testQueries.CreateRetweet(context.Background(), arg_tweet)
	require.NoError(t, err)
	require.NotEmpty(t, retweet)
	require.NotEmpty(t, retweet.ID)
	require.Equal(t, retweet.UserID, arg_tweet.UserID)
	require.Equal(t, retweet.ParentTweetID, arg_tweet.ParentTweetID)
	require.Equal(t, retweet.ParentRetweetID, arg_tweet.ParentRetweetID)
	require.Equal(t, retweet.Quote, arg_tweet.Quote)
	require.Equal(t, retweet.IsQuote, arg_tweet.IsQuote)
	require.NotEmpty(t, retweet.CreatedAt)
	require.NotEmpty(t, retweet.UpdatedAt)

	arg_tweet1 := CreateRetweetParams{}
	arg_tweet1.UserID = user_new.ID
	arg_tweet1.ParentTweetID.Int32 = tweet.ID
	arg_tweet1.ParentTweetID.Valid = true
	arg_tweet1.ParentRetweetID.Valid = false
	arg_tweet1.Quote.Scan(random.GenerateRandomAlphanumeric(64))
	arg_tweet1.IsQuote = true
	retweet1, err := testQueries.CreateRetweet(context.Background(), arg_tweet1)
	require.NoError(t, err)
	require.NotEmpty(t, retweet1)
	require.NotEmpty(t, retweet1.ID)
	require.Equal(t, retweet1.UserID, arg_tweet1.UserID)
	require.Equal(t, retweet1.ParentTweetID, arg_tweet1.ParentTweetID)
	require.Equal(t, retweet1.ParentRetweetID, arg_tweet1.ParentRetweetID)
	require.Equal(t, retweet1.Quote, arg_tweet1.Quote)
	require.Equal(t, retweet1.IsQuote, arg_tweet1.IsQuote)
	require.NotEmpty(t, retweet1.CreatedAt)
	require.NotEmpty(t, retweet1.UpdatedAt)

	retweet_arr, err := testQueries.SelectRetweetsOfTweet(context.Background(), pgtype.Int4{Int32: tweet.ID, Valid: true})
	require.NoError(t, err)
	require.Equal(t, len(retweet_arr), 2)

	require.NotEmpty(t, retweet_arr[0])
	require.Equal(t, user_new.ID, retweet_arr[0].ID)
	require.Equal(t, retweet1.ID, retweet_arr[0].RetweetID)
	require.Equal(t, retweet1.Quote, retweet_arr[0].Quote)
	require.Equal(t, retweet1.IsQuote, retweet_arr[0].IsQuote)
	require.Equal(t, retweet1.UserID, retweet_arr[0].UserID)
	require.Equal(t, retweet1.CreatedAt, retweet_arr[0].CreatedAt)

	require.NotEmpty(t, retweet_arr[1])
	require.Equal(t, user_new.ID, retweet_arr[1].ID)
	require.Equal(t, retweet.ID, retweet_arr[1].RetweetID)
	require.Equal(t, retweet.Quote, retweet_arr[1].Quote)
	require.Equal(t, retweet.IsQuote, retweet_arr[1].IsQuote)
	require.Equal(t, retweet.UserID, retweet_arr[1].UserID)
	require.Equal(t, retweet.CreatedAt, retweet_arr[1].CreatedAt)
}

func TestSelectRetweetsOfRetweet(t *testing.T) {
	user_orig := createRandomAccount(t)
	arg := CreateTweetParams{}
	arg.UserID = user_orig.ID
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

	user_new := createRandomAccount(t)

	arg_tweet := CreateRetweetParams{}
	arg_tweet.UserID = user_new.ID
	arg_tweet.ParentTweetID.Int32 = tweet.ID
	arg_tweet.ParentTweetID.Valid = true
	arg_tweet.ParentRetweetID.Valid = false
	arg_tweet.Quote.Scan(random.GenerateRandomAlphanumeric(64))
	arg_tweet.IsQuote = true
	retweet, err := testQueries.CreateRetweet(context.Background(), arg_tweet)
	require.NoError(t, err)
	require.NotEmpty(t, retweet)
	require.NotEmpty(t, retweet.ID)
	require.Equal(t, retweet.UserID, arg_tweet.UserID)
	require.Equal(t, retweet.ParentTweetID, arg_tweet.ParentTweetID)
	require.Equal(t, retweet.ParentRetweetID, arg_tweet.ParentRetweetID)
	require.Equal(t, retweet.Quote, arg_tweet.Quote)
	require.Equal(t, retweet.IsQuote, arg_tweet.IsQuote)
	require.NotEmpty(t, retweet.CreatedAt)
	require.NotEmpty(t, retweet.UpdatedAt)

	arg_tweet1 := CreateRetweetParams{}
	arg_tweet1.UserID = user_orig.ID
	arg_tweet1.ParentRetweetID.Int32 = retweet.ID
	arg_tweet1.ParentRetweetID.Valid = true
	arg_tweet1.ParentTweetID.Valid = false
	arg_tweet1.Quote.Scan(random.GenerateRandomAlphanumeric(64))
	arg_tweet1.IsQuote = true
	retweet1, err := testQueries.CreateRetweet(context.Background(), arg_tweet1)
	require.NoError(t, err)
	require.NotEmpty(t, retweet1)
	require.NotEmpty(t, retweet1.ID)
	require.Equal(t, retweet1.UserID, arg_tweet1.UserID)
	require.Equal(t, retweet1.ParentTweetID, arg_tweet1.ParentTweetID)
	require.Equal(t, retweet1.ParentRetweetID, arg_tweet1.ParentRetweetID)
	require.Equal(t, retweet1.Quote, arg_tweet1.Quote)
	require.Equal(t, retweet1.IsQuote, arg_tweet1.IsQuote)
	require.NotEmpty(t, retweet1.CreatedAt)
	require.NotEmpty(t, retweet1.UpdatedAt)

	arg_tweet2 := CreateRetweetParams{}
	arg_tweet2.UserID = user_orig.ID
	arg_tweet2.ParentRetweetID.Int32 = retweet.ID
	arg_tweet2.ParentRetweetID.Valid = true
	arg_tweet2.ParentTweetID.Valid = false
	arg_tweet2.Quote.Scan(random.GenerateRandomAlphanumeric(64))
	arg_tweet2.IsQuote = true
	retweet2, err := testQueries.CreateRetweet(context.Background(), arg_tweet2)
	require.NoError(t, err)
	require.NotEmpty(t, retweet2)
	require.NotEmpty(t, retweet2.ID)
	require.Equal(t, retweet2.UserID, arg_tweet2.UserID)
	require.Equal(t, retweet2.ParentTweetID, arg_tweet2.ParentTweetID)
	require.Equal(t, retweet2.ParentRetweetID, arg_tweet2.ParentRetweetID)
	require.Equal(t, retweet2.Quote, arg_tweet2.Quote)
	require.Equal(t, retweet2.IsQuote, arg_tweet2.IsQuote)
	require.NotEmpty(t, retweet2.CreatedAt)
	require.NotEmpty(t, retweet2.UpdatedAt)

	retweet_arr, err := testQueries.SelectRetweetsOfRetweet(context.Background(), pgtype.Int4{Int32: retweet.ID, Valid: true})
	require.NoError(t, err)
	require.Equal(t, len(retweet_arr), 2)

	require.NotEmpty(t, retweet_arr[1])
	require.Equal(t, user_orig.ID, retweet_arr[1].ID)
	require.Equal(t, retweet1.ID, retweet_arr[1].RetweetID)
	require.Equal(t, retweet1.Quote, retweet_arr[1].Quote)
	require.Equal(t, retweet1.IsQuote, retweet_arr[1].IsQuote)
	require.Equal(t, retweet1.UserID, retweet_arr[1].UserID)
	require.Equal(t, retweet1.CreatedAt, retweet_arr[1].CreatedAt)

	require.NotEmpty(t, retweet_arr[0])
	require.Equal(t, user_orig.ID, retweet_arr[0].ID)
	require.Equal(t, retweet2.ID, retweet_arr[0].RetweetID)
	require.Equal(t, retweet2.Quote, retweet_arr[0].Quote)
	require.Equal(t, retweet2.IsQuote, retweet_arr[0].IsQuote)
	require.Equal(t, retweet2.UserID, retweet_arr[0].UserID)
	require.Equal(t, retweet2.CreatedAt, retweet_arr[0].CreatedAt)
}
