package jwtGin

import (
	jwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
}
