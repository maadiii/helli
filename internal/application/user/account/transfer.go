package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/jwt"
	"github.com/maadiii/hertzwrapper/server"
)

// @action /api/v1/accounts/transfer [POST] 204 application/json
func Transfer(ctx *server.Context, req *RequestTransfer) (out any, err error) {
	username := ctx.Identity()["username"].(string)
	src := &entity.Account{ID: req.SrcAccountID, Username: username}

	err = service.Transfer(src, req.DstUserID, req.DstAccountID, req.Amount)

	return
}

type RequestTransfer struct {
	jwt.Token
	SrcAccountID int     `json:"src_account_id"`
	DstUserID    int     `json:"dst_user_id"`
	DstAccountID int     `json:"dst_account_id"`
	Amount       float64 `json:"amount"`
}
