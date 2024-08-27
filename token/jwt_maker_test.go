package token

import (
	"testing"
	"time"

	"github.com/felipeazsantos/simple_bank/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, claims)

	mapClaims := *claims.(*jwt.MapClaims)

	issuedAtToken := time.Unix(int64(mapClaims["iat"].(float64)), 0)
	expiredAtToken := time.Unix(int64(mapClaims["exp"].(float64)), 0)

	require.NotZero(t, mapClaims["id"])
	require.Equal(t, username, mapClaims["username"])
	require.WithinDuration(t, issuedAt, issuedAtToken, time.Second)
	require.WithinDuration(t, expiredAt, expiredAtToken, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, jwt.ErrTokenExpired.Error())
	require.Nil(t, claims)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	claims, err := NewJWTClaims(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	mapClaims, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errTokenInvalid.Error())
	require.Nil(t, mapClaims)
}
