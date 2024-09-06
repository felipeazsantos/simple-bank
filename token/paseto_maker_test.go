package token

import (
	"testing"
	"time"

	"github.com/felipeazsantos/simple_bank/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payloadResponse, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotNil(t, payload)

	payloadObj := *payloadResponse.(*Payload)

	require.NotZero(t, payloadObj.ID)
	require.Equal(t, username, payloadObj.Username)
	require.WithinDuration(t, issuedAt, payloadObj.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payloadObj.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, payload, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payloadResponse, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, jwt.ErrTokenExpired.Error())
	require.Nil(t, payloadResponse)
}
