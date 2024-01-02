package user

import (
	"github.com/maadiii/helli/internal/domain/user"
	"github.com/maadiii/hertzwrapper/server"
)

var service user.Service

func Init(s user.Service) {
	service = s

	server.Handle(Register)
	server.Handle(Login)
}
