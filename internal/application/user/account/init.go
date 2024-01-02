package account

import (
	"github.com/maadiii/helli/internal/domain/account"
	"github.com/maadiii/helli/pkg/jwt"
	"github.com/maadiii/hertzwrapper/server"
)

var service account.Service

func Init(s account.Service) {
	service = s

	server.Handle(jwt.Auth, Create)
	server.Handle(jwt.Auth, Increase)
	server.Handle(jwt.Auth, Decrease)
	server.Handle(jwt.Auth, Transfer)
	server.Handle(jwt.Auth, Get)
}
