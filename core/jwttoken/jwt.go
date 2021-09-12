package jwttoken

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strings"
	"time"
)

type jwtToken struct{}

type JWTClaims struct {
	jwt.StandardClaims
	UserId int64
	Role   string
}

type JWTPayload struct {
	UserId int64
	Role   string
}

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrTokenExpired = errors.New("token expired")
)

func NewJWTToken() *jwtToken {
	return &jwtToken{}
}

func (*jwtToken) GenerateToken(payload JWTPayload) (signedToken string, err error) {
	claims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    os.Getenv("APP_NAME"),
		},
		UserId: payload.UserId,
		Role: payload.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func (*jwtToken) ValidateToken(tokenStr string) (result *JWTClaims, err error) {
	payload := strings.Split(tokenStr, " ")
	if len(payload) != 2 || payload[0] != "Bearer" {
		err = ErrInvalidToken
		return
	}

	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(payload[1], claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, ErrInvalidToken
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		if v.Errors == jwt.ValidationErrorExpired {
			err = ErrTokenExpired
			return
		}
		return
	}

	if c, ok := token.Claims.(*JWTClaims); ok {
		return c, nil
	}

	return nil, ErrInvalidToken
}
