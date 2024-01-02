package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/jwt"
	"github.com/maadiii/hertzwrapper/server"
)

// @action /api/v1/accounts [POST] 201 application/json
func Create(ctx *server.Context, req *RequestCreate) (out *ResponseCreate, err error) {
	entity := req.into(ctx)

	if err = service.Create(entity); err != nil {
		return
	}

	out = &ResponseCreate{ID: entity.ID}

	return
}

type RequestCreate struct {
	jwt.Token
	Fund float64 `json:"fund,omitempty"`
}

func (r RequestCreate) into(ctx *server.Context) (account *entity.Account) {
	username := ctx.Identity()["username"].(string)
	account = &entity.Account{
		Fund:     r.Fund,
		Username: username,
	}

	return
}

type ResponseCreate struct {
	ID int `json:"id,omitempty"`
}
