package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("segreto-bem-guardado")

type Jwt struct {
	Token  string
	Expire string
}

func GenerateJwt(username string) (*Jwt, error) {
	exp := time.Now().Add(15 * time.Minute)
	claims := jwt.MapClaims{
		"username": username,
		"exp":      exp.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(SecretKey)
	if err != nil {
		return nil, err
	}

	return &Jwt{Token: tokenStr, Expire: exp.Local().String()}, nil
}
