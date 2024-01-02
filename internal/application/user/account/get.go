package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/jwt"
	"github.com/maadiii/hertzwrapper/server"
)

// @action /api/v1/accounts/:id [GET] 200 application/json
func Get(ctx *server.Context, req *RequestGet) (out *ResponseGet, err error) {
	entity := req.into(ctx)

	if err = service.Get(entity); err != nil {
		return
	}

	return new(ResponseGet).from(entity), nil
}

type RequestGet struct {
	jwt.Token
	ID int `path:"id"`
}

func (r RequestGet) into(ctx *server.Context) (account *entity.Account) {
	username := ctx.Identity()["username"].(string)
	account = &entity.Account{
		ID:       r.ID,
		Username: username,
	}

	return
}

type ResponseGet struct {
	ID   int     `json:"id,omitempty"`
	Fund float64 `json:"fund,omitempty"`
}

func (r *ResponseGet) from(entity *entity.Account) *ResponseGet {
	r.ID = entity.ID
	r.Fund = entity.Fund

	return r
}
