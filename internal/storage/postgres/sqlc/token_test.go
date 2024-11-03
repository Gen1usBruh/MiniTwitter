package postgresdb

import (
	"context"
	"testing"
	"time"

	"github.com/Gen1usBruh/MiniTwitter/util/random"
	"github.com/stretchr/testify/require"
)

func CreateRandomToken(t *testing.T) RefreshToken {
	user := createRandomAccount(t)
	arg := CreateTokenParams{}
	arg.UserID = user.ID
	arg.Token = random.GenerateRandomAlphanumeric(64)
	arg.ExpiresAt.Time = time.Now().Add(24 * time.Hour)
	arg.ExpiresAt.Valid = true
	token, err := testQueries.CreateToken(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, token.ID)
	require.Equal(t, token.UserID, arg.UserID)
	require.Equal(t, token.Token, arg.Token)
	// require.Equal(t, token.ExpiresAt, arg.ExpiresAt)

	return token
}

func TestCreateToken(t *testing.T) {
	CreateRandomToken(t)
}

func TestDeleteToken(t *testing.T) {
	token := CreateRandomToken(t)
	arg := DeleteTokenParams{}
	arg.Token = token.Token
	arg.UserID = token.UserID
	err := testQueries.DeleteToken(context.Background(), arg)
	require.NoError(t, err)
}

func TestSelectToken(t *testing.T) {
	token := CreateRandomToken(t)
	arg := SelectTokenParams{}
	arg.Token = token.Token
	arg.UserID = token.UserID
	tokenSelect, err := testQueries.SelectToken(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tokenSelect)
	require.Equal(t, token.ID, tokenSelect.ID)
	require.Equal(t, token.UserID, tokenSelect.UserID)
	require.Equal(t, token.Token, tokenSelect.Token)
}
