package jwtGin

import (
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	signinKey []byte
}

func (j *Manager) Generate(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signinKey)
}

func (j *Manager) Parse(signedStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(signedStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signinKey, nil
	})

	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("could not handle this token")

	} else {
		return nil, errors.New("could not handle this token")
	}
}
