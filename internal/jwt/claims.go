package jwt

import (
	"errors"
	"time"

	"github.com/Babatunde50/green-light/internal/data"
	"github.com/golang-jwt/jwt/v5"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload contains the payload data of the token
type Claims struct {
	data.User
	jwt.RegisteredClaims
}

func NewClaim(user data.User, duration time.Duration) *Claims {

	claim := &Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return claim
}
