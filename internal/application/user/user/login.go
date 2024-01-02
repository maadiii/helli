package user

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/hertzwrapper/server"
)

// @action /api/v1/auth/login [POST] 200 application/json
func Login(_ *server.Context, req *RequestLogin) (out *ResponseLogin, err error) {
	access, refresh, err := service.Login(req.into())
	if err != nil {
		return
	}

	out = new(ResponseLogin).from(access, refresh)

	return
}

type RequestLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r RequestLogin) into() *entity.User {
	return &entity.User{
		Username: r.Username,
		Password: r.Password,
	}
}

type ResponseLogin struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func (r *ResponseLogin) from(access, refresh string) *ResponseLogin {
	r.AccessToken = access
	r.RefreshToken = refresh

	return r
}
