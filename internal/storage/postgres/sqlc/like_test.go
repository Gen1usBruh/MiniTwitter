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

func CreateRandomLikeTweet(t *testing.T) LikeTweet {
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

	argLike := CreateLikeTweetParams{}
	argLike.UserID = user.ID
	argLike.TweetID.Int32 = tweet.ID
	argLike.TweetID.Valid = true
	argLike.RetweetID.Valid = false 
	argLike.PostType = "tweet"
	
	like, err := testQueries.CreateLikeTweet(context.Background(), argLike)
	require.NoError(t, err)
	require.NotEmpty(t, like)
	require.NotEmpty(t, like.ID)
	require.Equal(t, argLike.TweetID, like.TweetID)
	require.Equal(t, argLike.RetweetID, like.RetweetID)
	require.Equal(t, argLike.UserID, like.UserID)
	require.Equal(t, argLike.PostType, like.PostType)
	require.NotEmpty(t, like.CreatedAt)
	require.NotEmpty(t, like.UpdatedAt)

	return like
}

func TestCreateLikeTweet(t *testing.T) {
	CreateRandomLikeTweet(t)
}

func TestDeleteLikeTweet(t *testing.T) {
	tweet := CreateRandomLikeTweet(t)

	arg := DeleteLikeTweetParams{}
	arg.UserID = tweet.UserID
	arg.Column2 = "tweet"
	arg.TweetID.Int32 = tweet.ID
	arg.TweetID.Valid = true

	err := testQueries.DeleteLikeTweet(context.Background(), arg)
	require.NoError(t, err)
} 

func TestSelectLikeTweet(t *testing.T) {
	user := createRandomAccount(t)

	argTweet := CreateTweetParams{}
	argTweet.UserID = user.ID
	argTweet.Content = random.GenerateRandomAlphanumeric(64)
	media := MediaLinks{
		Video: []string{random.GenerateRandomAlphanumeric(16)},
		Photo: []string{random.GenerateRandomAlphanumeric(16),
			random.GenerateRandomAlphanumeric(16)},
	}
	var err error
	argTweet.Media, err = json.Marshal(media)
	assert.NoError(t, err, "No error marshaling JSON for the sake of test")

	tweet, err := testQueries.CreateTweet(context.Background(), argTweet)
	require.NoError(t, err)
	require.NotEmpty(t, tweet)

	// tweet := CreateRandomLikeTweet(t)
	// fmt.Println(tweet)
	user1 := createRandomAccount(t)
	user2 := createRandomAccount(t)
	// fmt.Println(user1)
	// fmt.Println(user2)
	argLike1 := CreateLikeTweetParams{}
	argLike1.UserID = user1.ID
	argLike1.TweetID.Int32 = tweet.ID
	argLike1.TweetID.Valid = true
	argLike1.RetweetID.Valid = false 
	argLike1.PostType = "tweet"
	// fmt.Println(argLike1)
	like1, err := testQueries.CreateLikeTweet(context.Background(), argLike1)
	require.NoError(t, err)
	require.NotEmpty(t, like1)

	argLike2 := CreateLikeTweetParams{}
	argLike2.UserID = user2.ID
	argLike2.TweetID.Int32 = tweet.ID
	argLike2.TweetID.Valid = true
	argLike2.RetweetID.Valid = false 
	argLike2.PostType = "tweet"
	like2, err := testQueries.CreateLikeTweet(context.Background(), argLike2)
	require.NoError(t, err)
	require.NotEmpty(t, like2)

	arg := SelectLikeTweetParams{}
	arg.Column1 = "tweet"
	arg.TweetID.Int32 = tweet.ID
	arg.TweetID.Valid = true 

	like_arr, err := testQueries.SelectLikeTweet(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, len(like_arr), 2)
	
	require.NotEmpty(t, like_arr[1])
	require.NotEmpty(t, like_arr[1].ID)
	require.Equal(t, argLike1.UserID, like_arr[1].UserID)
	require.NotEmpty(t, like_arr[1].Username)
	require.NotEmpty(t, like_arr[1].CreatedAt)

	require.NotEmpty(t, like_arr[0])
	require.NotEmpty(t, like_arr[0].ID)
	require.Equal(t, argLike2.UserID, like_arr[0].UserID)
	require.NotEmpty(t, like_arr[0].Username)
	require.NotEmpty(t, like_arr[0].CreatedAt)
}

func TestCreateLikeComment(t *testing.T) {
	comm := CreateRandomComment(t)
	arg := CreateLikeCommentParams{}
	arg.UserID = comm.UserID
	arg.CommentID = comm.ID
	likeComm, err := testQueries.CreateLikeComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, likeComm)
}

func TestDeleteLikeComment(t *testing.T) {
	comm := CreateRandomComment(t)
	arg := CreateLikeCommentParams{}
	arg.UserID = comm.UserID
	arg.CommentID = comm.ID
	likeComm, err := testQueries.CreateLikeComment(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, likeComm)
	argDel := DeleteLikeCommentParams{}
	argDel.CommentID = comm.ID
	argDel.UserID = comm.UserID
	err = testQueries.DeleteLikeComment(context.Background(), argDel)
	require.NoError(t, err)
}

func TestSelectLikeComment(t *testing.T) {
	comm := CreateRandomComment(t)

	user1 := createRandomAccount(t)
	user2 := createRandomAccount(t)

	arg1 := CreateLikeCommentParams{}
	arg1.UserID = user1.ID
	arg1.CommentID = comm.ID
	likeComm1, err := testQueries.CreateLikeComment(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, likeComm1)	

	arg2 := CreateLikeCommentParams{}
	arg2.UserID = user2.ID
	arg2.CommentID = comm.ID
	likeComm2, err := testQueries.CreateLikeComment(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, likeComm2)

	like_arr, err := testQueries.SelectLikeComment(context.Background(), comm.ID)
	require.NoError(t, err)
	require.Equal(t, len(like_arr), 2)

	require.Equal(t, like_arr[0].UserID, likeComm2.UserID)
	require.Equal(t, like_arr[0].ID, likeComm2.ID)
	require.NotEmpty(t, like_arr[0].Username)
	require.Equal(t, like_arr[0].CreatedAt, likeComm2.CreatedAt)

	require.Equal(t, like_arr[1].UserID, likeComm1.UserID)
	require.Equal(t, like_arr[1].ID, likeComm1.ID)
	require.NotEmpty(t, like_arr[1].Username)
	require.Equal(t, like_arr[1].CreatedAt, likeComm1.CreatedAt)
}
