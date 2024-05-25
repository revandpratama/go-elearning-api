package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/revandpratama/go-elearning-api/model"
)

type JWTCustomClaims struct {
	ID int
	jwt.RegisteredClaims
}

var privateKey = []byte("this is between us")

func GenerateToken(user *model.User) (*string, error) {
	claims := &JWTCustomClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(privateKey)

	return &ss, err
}

func ValidateToken(tokenString string) {

}
