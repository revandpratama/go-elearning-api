package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/revandpratama/go-elearning-api/model"
)

type JWTCustomClaims struct {
	ID   int
	Role string
	jwt.RegisteredClaims
}

var privateKey = []byte("this is between us")

func GenerateToken(user *model.User) (*string, error) {
	claims := &JWTCustomClaims{
		user.ID,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 2)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(privateKey)

	return &ss, err
}

func ValidateToken(tokenString string) (*int, *string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, nil, errors.New("invalid token")
		}

		return nil, nil, errors.New("token expired")
	}

	claims, ok := token.Claims.(*JWTCustomClaims)

	if !ok || !token.Valid {
		return nil, nil, errors.New("token expired")
	}

	return &claims.ID, &claims.Role, nil
}
