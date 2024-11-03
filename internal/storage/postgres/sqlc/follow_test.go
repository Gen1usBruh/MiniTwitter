package postgresdb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateFollower(t *testing.T) {
	user_er := createRandomAccount(t)
	user_ing := createRandomAccount(t)

	arg := CreateFollowerParams{}
	arg.FollowerID = user_er.ID
	arg.FollowingID = user_ing.ID

	follow, err := testQueries.CreateFollower(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow)
	require.Equal(t, follow.FollowerID, user_er.ID)
	require.Equal(t, follow.FollowingID, user_ing.ID)
	require.NotEmpty(t, follow.CreatedAt)
	require.NotEmpty(t, follow.UpdatedAt)
}

func TestDeleteFollower(t *testing.T) {
	user_er := createRandomAccount(t)
	user_ing := createRandomAccount(t)

	arg := CreateFollowerParams{}
	arg.FollowerID = user_er.ID
	arg.FollowingID = user_ing.ID

	follow, err := testQueries.CreateFollower(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow)

	argDel := DeleteFollowerParams{}
	argDel.FollowerID = follow.FollowerID
	argDel.FollowingID = follow.FollowingID
	err = testQueries.DeleteFollower(context.Background(), argDel)
	require.NoError(t, err)
}
