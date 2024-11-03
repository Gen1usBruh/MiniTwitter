package postgresdb

import (
	"context"
	"testing"

	"github.com/Gen1usBruh/MiniTwitter/util/hash"
	"github.com/Gen1usBruh/MiniTwitter/util/random"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) User {
	hashPassw, err := hash.HashPassword(random.GenerateRandomAlphanumeric(16))
	assert.Nil(t, err, "No hashing errors for the sake of testing")

	arg := CreateUserParams{}
	arg.Username = random.GenerateRandomAlphanumeric(32)
	arg.Email = random.GenerateRandomAlphanumeric(32)
	arg.Password = hashPassw
	err = arg.Bio.Scan(random.GenerateRandomAlphanumeric(32))
	assert.Nil(t, err, "No Scan errors for the sake of testing")

	account, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.Password, account.Password)
	require.Equal(t, arg.Bio, account.Bio)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)

	return account
}

func TestCreateUser(t *testing.T) {
	createRandomAccount(t)
}

func TestGetUserCred(t *testing.T) {
	user_create := createRandomAccount(t)
	user_select, err := testQueries.SelectUserCred(context.Background(), user_create.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user_select)

	require.Equal(t, user_create.Username, user_select.Username)
	require.Equal(t, user_create.Email, user_select.Email)
	require.Equal(t, user_create.Password, user_select.Password)
}

func TestGetUserData(t *testing.T) {
	user_create := createRandomAccount(t)
	user_select, err := testQueries.SelectUserData(context.Background(), user_create.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user_select)

	require.Equal(t, user_create.ID, user_select.ID)
	require.Equal(t, user_create.Username, user_select.Username)
	require.Equal(t, user_create.Email, user_select.Email)
	require.Equal(t, user_create.Password, user_select.Password)
	require.Equal(t, user_create.Bio, user_select.Bio)
	require.Equal(t, user_create.CreatedAt, user_select.CreatedAt)
	require.Equal(t, user_create.UpdatedAt, user_select.CreatedAt)
}

func TestDeleteUser(t *testing.T) {
	user_create := createRandomAccount(t)
	err := testQueries.DeleteUser(context.Background(), user_create.ID)
	require.NoError(t, err)

	user_select, err := testQueries.SelectUserData(context.Background(), user_create.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, user_select)
}

func TestUpdateUserBio(t *testing.T) {
	user_create := createRandomAccount(t)

	arg := UpdateUserBioParams{}
	arg.ID = user_create.ID
	err := arg.Bio.Scan(random.GenerateRandomAlphanumeric(32))
	assert.Nil(t, err, "No Scan errors for the sake of testing")

	err = testQueries.UpdateUserBio(context.Background(), arg)
	require.NoError(t, err)

	user_select, err := testQueries.SelectUserData(context.Background(), user_create.ID)
	require.NoError(t, err)
	require.NotEqual(t, user_create.Bio, user_select.Bio)
	require.NotEqual(t, user_create.UpdatedAt, user_select.UpdatedAt)
	require.Equal(t, user_create.ID, user_select.ID)
	require.Equal(t, user_create.Username, user_select.Username)
	require.Equal(t, user_create.Email, user_select.Email)
	require.Equal(t, user_create.Password, user_select.Password)
	require.Equal(t, user_create.CreatedAt, user_select.CreatedAt)

	/*
		err = testQueries.DeleteUser(context.Background(), user_create.ID)
		require.NoError(t, err)
	*/
}

func TestUpdateUserName(t *testing.T) {
	user_create := createRandomAccount(t)

	arg := UpdateUserNameParams{}
	arg.ID = user_create.ID
	arg.Username = random.GenerateRandomAlphanumeric(32)

	err := testQueries.UpdateUserName(context.Background(), arg)
	require.NoError(t, err)

	user_select, err := testQueries.SelectUserData(context.Background(), user_create.ID)
	require.NoError(t, err)
	require.NotEqual(t, user_create.Username, user_select.Username)
	require.NotEqual(t, user_create.UpdatedAt, user_select.UpdatedAt)
	require.Equal(t, user_create.ID, user_select.ID)
	require.Equal(t, user_create.Bio, user_select.Bio)
	require.Equal(t, user_create.Email, user_select.Email)
	require.Equal(t, user_create.Password, user_select.Password)
	require.Equal(t, user_create.CreatedAt, user_select.CreatedAt)

	/*
		err = testQueries.DeleteUser(context.Background(), user_create.ID)
		require.NoError(t, err)
	*/
}

func TestSelectUserFollowers(t *testing.T) {
	user_er1 := createRandomAccount(t)
	user_er2 := createRandomAccount(t)
	user_ing := createRandomAccount(t)

	arg := CreateFollowerParams{}
	arg.FollowerID = user_er1.ID
	arg.FollowingID = user_ing.ID
	follow1, err := testQueries.CreateFollower(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow1)

	arg1 := CreateFollowerParams{}
	arg1.FollowerID = user_er2.ID
	arg1.FollowingID = user_ing.ID
	follow2, err := testQueries.CreateFollower(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, follow2)

	follow_arr, err := testQueries.SelectUserFollowers(context.Background(), user_ing.ID)
	require.NoError(t, err)
	require.NotEmpty(t, follow_arr)
	require.Equal(t, len(follow_arr), 2)

	require.Equal(t, follow_arr[0].ID, user_er2.ID)
	require.Equal(t, follow_arr[0].Username, user_er2.Username)
	require.Equal(t, follow_arr[0].Bio, user_er2.Bio)
	require.Equal(t, follow_arr[0].FollowDate, follow2.CreatedAt)

	require.Equal(t, follow_arr[1].ID, user_er1.ID)
	require.Equal(t, follow_arr[1].Username, user_er1.Username)
	require.Equal(t, follow_arr[1].Bio, user_er1.Bio)
	require.Equal(t, follow_arr[1].FollowDate, follow1.CreatedAt)
}

func TestSelectUserFollowing(t *testing.T) {
	user_er1 := createRandomAccount(t)
	user_er2 := createRandomAccount(t)
	user_ing := createRandomAccount(t)

	arg := CreateFollowerParams{}
	arg.FollowerID = user_ing.ID
	arg.FollowingID = user_er1.ID
	follow1, err := testQueries.CreateFollower(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow1)

	arg1 := CreateFollowerParams{}
	arg1.FollowerID = user_ing.ID
	arg1.FollowingID = user_er2.ID
	follow2, err := testQueries.CreateFollower(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, follow2)

	follow_arr, err := testQueries.SelectUserFollowing(context.Background(), user_ing.ID)
	require.NoError(t, err)
	require.NotEmpty(t, follow_arr)
	require.Equal(t, len(follow_arr), 2)

	require.Equal(t, follow_arr[0].ID, user_er2.ID)
	require.Equal(t, follow_arr[0].Username, user_er2.Username)
	require.Equal(t, follow_arr[0].Bio, user_er2.Bio)
	require.Equal(t, follow_arr[0].FollowDate, follow2.CreatedAt)

	require.Equal(t, follow_arr[1].ID, user_er1.ID)
	require.Equal(t, follow_arr[1].Username, user_er1.Username)
	require.Equal(t, follow_arr[1].Bio, user_er1.Bio)
	require.Equal(t, follow_arr[1].FollowDate, follow1.CreatedAt)
}
