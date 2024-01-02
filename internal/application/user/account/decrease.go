package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/jwt"
	"github.com/maadiii/hertzwrapper/server"
)

// @action /api/v1/accounts/decrease [POST] 204 application/json
func Decrease(ctx *server.Context, req *RequestDecrease) (out any, err error) {
	entity := req.into(ctx)

	return out, service.Decrease(entity, req.Amount)
}

type RequestDecrease struct {
	jwt.Token
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
}

func (r RequestDecrease) into(ctx *server.Context) (account *entity.Account) {
	username := ctx.Identity()["username"].(string)
	account = &entity.Account{
		Username: username,
		ID:       r.ID,
	}

	return
}
