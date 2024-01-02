package user

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/hertzwrapper/server"
)

// @action /api/v1/auth/register [POST] 201 application/json
func Register(_ *server.Context, req *RequestRegister) (out *ResponseRegister, err error) {
	entity := req.into()
	if err = service.Register(entity); err != nil {
		return
	}

	out = new(ResponseRegister).from(entity)

	return
}

type RequestRegister struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (r RequestRegister) into() *entity.User {
	return &entity.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Username:  r.Username,
		Password:  r.Password,
	}
}

type ResponseRegister struct {
	ID int `json:"id,omitempty"`
}

func (r *ResponseRegister) from(entity *entity.User) *ResponseRegister {
	r.ID = entity.ID

	return r
}
