package user

import (
	"github.com/maadiii/helli/internal/data_access/user"
	"github.com/maadiii/helli/internal/entity"
)

type service struct {
	users user.DataAccess
}

func New(dataAccess user.DataAccess) Service {
	return &service{dataAccess}
}

type Service interface {
	Register(entity *entity.User) error
	Login(entity *entity.User) (access, refresh string, err error)
}
