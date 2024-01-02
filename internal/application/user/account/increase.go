package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/jwt"
	"github.com/maadiii/hertzwrapper/server"
)

// @action /api/v1/accounts/increase [POST] 204 application/json
func Increase(ctx *server.Context, req *RequestIncrease) (out any, err error) {
	entity := req.into(ctx)

	return out, service.Increase(entity, req.Amount)
}

type RequestIncrease struct {
	jwt.Token
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
}

func (r RequestIncrease) into(ctx *server.Context) (account *entity.Account) {
	username := ctx.Identity()["username"].(string)
	account = &entity.Account{
		Username: username,
		ID:       r.ID,
	}

	return
}
