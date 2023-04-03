package jwt

import (
	"time"

	"github.com/Babatunde50/green-light/internal/data"
	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) *JWTMaker {

	return &JWTMaker{secretKey}
}

func (maker *JWTMaker) CreateToken(user data.User, duration time.Duration) (string, error) {
	claim := NewClaim(user, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, err
}

func (maker *JWTMaker) VerifyToken(token string) (*Claims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// _, ok := token.Method.(*jwt.SigningMethodHMAC)
		// if !ok {
		// 	return nil, ErrInvalidToken
		// }
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, keyFunc)
	if err != nil {
		return nil, ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
