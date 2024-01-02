package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/maadiii/helli/config"
	"github.com/maadiii/hertzwrapper/errors"
	"github.com/maadiii/hertzwrapper/server"
)

func Auth[IN ijwt, OUT any](ctx *server.Context, in IN) (out OUT, err error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(in.AccessToken(), claims, func(*jwt.Token) (any, error) {
		return []byte(config.JWT().Secret), nil
	})
	if err != nil {
		return out, errors.Unauthorized
	}

	if token == nil || !token.Valid {
		exp, ex := token.Claims.GetExpirationTime()
		if ex != nil {
			return out, errors.Unauthorized
		}

		if exp.Before(time.Now()) {
			return out, errors.Unauthorized.Pattern("expired token")
		}
	}

	identity := server.Identity{
		"id":         claims.ID,
		"first_name": claims.FirstName,
		"last_name":  claims.LastName,
		"username":   claims.Username,
	}

	ctx.SetIdentity(identity)

	return
}

type ijwt interface {
	AccessToken() string
	RefreshToken() string
}

type Token struct {
	Access  string `header:"access_token"`
	Refresh string `header:"refresh_token"`
}

func (t Token) AccessToken() string {
	return t.Access
}

func (t Token) RefreshToken() string {
	return t.Refresh
}

type Out interface{}
