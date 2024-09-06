package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const minSecretKeySize = 32

var errTokenExpired = errors.New("token is expired")
var errTokenInvalid = errors.New("token is invalid")

// JWTMaker is a JSON Web Token Maker
type JWTMaker struct {
	secretKey string
	claims    jwt.Claims
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (*JWTMaker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

// NewJWTClaims creates a new token payload with a specific username and duration
func NewJWTClaims(username string, duration time.Duration) (*jwt.MapClaims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	claims := &jwt.MapClaims{
		"id":       tokenID,
		"username": username,
		"iat":      now.Unix(),
		"exp":      now.Add(duration).Unix(),
	}

	return claims, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	var err error
	maker.claims, err = NewJWTClaims(username, duration)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, maker.claims)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", nil, err
	}

	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", nil, err
	}
	return token, payload, nil
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (jwt.Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errTokenInvalid
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, maker.claims, keyFunc)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenInvalidClaims) {
			return nil, errTokenExpired
		}
		return nil, errTokenInvalid
	}

	return jwtToken.Claims.(*jwt.MapClaims), nil
}

// TokenValid checks if the token payload is valid or not
func TokenValid(expiredAt time.Time) error {
	if time.Now().After(expiredAt) {
		return errTokenExpired
	}
	return nil
}
