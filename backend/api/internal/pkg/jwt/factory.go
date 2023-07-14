package jwtGin

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewManager(key string) *Manager {
	return &Manager{[]byte(key)}
}

func NewClaims(id string) *Claims {
	nowTime := time.Now()
	tomorrow := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, nowTime.UTC().Location()).AddDate(0, 0, 1)
	return &Claims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tomorrow),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			Issuer:    "sigma",
			Subject:   "prototype",
			ID:        id,
			Audience:  []string{""},
		},
	}

}
