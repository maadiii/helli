package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maadiii/helli/config"
)

type Claims struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
	jwt.RegisteredClaims
}

type refreshClaims struct {
	ID int `json:"id,omitempty"`
	jwt.RegisteredClaims
}

func Create(claims Claims) (access, refresh string, err error) {
	claims.ExpiresAt = &jwt.NumericDate{
		Time: time.Now().Add(config.JWT().Expiration),
	}

	access, err = jwt.
		NewWithClaims(config.JWT().Algorithm, claims).
		SignedString([]byte(config.JWT().Secret))
	if err != nil {
		return
	}

	refreshClaims := refreshClaims{ID: claims.ID}
	refresh, err = jwt.
		NewWithClaims(config.JWT().RefreshAlgorithm, refreshClaims).
		SignedString([]byte(config.JWT().RefreshSecret))

	return
}
