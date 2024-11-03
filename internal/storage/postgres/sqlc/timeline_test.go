package postgresdb

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Gen1usBruh/MiniTwitter/util/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSelectTimeline(t *testing.T) {
	user := createRandomAccount(t)
	user1 := createRandomAccount(t)
	user2 := createRandomAccount(t)
	var err error

	arg1 := CreateFollowerParams{}
	arg1.FollowerID = user.ID
	arg1.FollowingID = user1.ID
	follow1, err := testQueries.CreateFollower(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, follow1)

	arg2 := CreateFollowerParams{}
	arg2.FollowerID = user.ID
	arg2.FollowingID = user2.ID
	follow2, err := testQueries.CreateFollower(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, follow2)

	media := MediaLinks{
		Video: []string{random.GenerateRandomAlphanumeric(16)},
		Photo: []string{random.GenerateRandomAlphanumeric(16),
			random.GenerateRandomAlphanumeric(16)},
	}

	argTweet := CreateTweetParams{}
	argTweet.UserID = user.ID
	argTweet.Content = random.GenerateRandomAlphanumeric(64)
	argTweet.Media, err = json.Marshal(media)
	assert.NoError(t, err, "No error marshaling JSON for the sake of test")
	tweet, err := testQueries.CreateTweet(context.Background(), argTweet)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)

	argTweet1 := CreateTweetParams{}
	argTweet1.UserID = user1.ID
	argTweet1.Content = random.GenerateRandomAlphanumeric(64)
	argTweet1.Media, err = json.Marshal(media)
	assert.NoError(t, err, "No error marshaling JSON for the sake of test")
	tweet1, err := testQueries.CreateTweet(context.Background(), argTweet1)
	require.NoError(t, err)
	require.NotEmpty(t, tweet1)

	argTweet2 := CreateRetweetParams{}
	argTweet2.UserID = user2.ID
	argTweet2.ParentTweetID.Int32 = tweet.ID
	argTweet2.ParentTweetID.Valid = true
	argTweet2.ParentRetweetID.Valid = false
	argTweet2.Quote.Scan(random.GenerateRandomAlphanumeric(64))
	argTweet2.IsQuote = true
	retweet, err := testQueries.CreateRetweet(context.Background(), argTweet2)
	require.NoError(t, err)
	require.NotEmpty(t, retweet)

	arg := SelectTimelineParams{}
	arg.FollowerID = user.ID
	arg.Limit = 10
	arg.Offset = 0
	timeline, err := testQueries.SelectTimeline(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, timeline)
	require.Equal(t, len(timeline), 2)

	require.Equal(t, timeline[0].PostID, retweet.ID)
	require.Equal(t, timeline[0].PostType, "retweet")
	require.Equal(t, timeline[0].UserID, retweet.UserID)

	require.Equal(t, timeline[1].PostID, tweet1.ID)
	require.Equal(t, timeline[1].PostType, "tweet")
	require.Equal(t, timeline[1].UserID, tweet1.UserID)
}
