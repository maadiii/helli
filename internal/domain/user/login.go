package user

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/jwt"
	"github.com/maadiii/hertzwrapper/errors"
)

func (s *service) Login(entity *entity.User) (access, refresh string, err error) {
	password := entity.Password //nolint

	if err = s.users.GetByUsername(entity); err != nil {
		return
	}

	if password != entity.Password {
		return "", "", errors.BadRequest.Pattern("username or password is wrong")
	}

	return s.login(entity)
}

func (s *service) login(entity *entity.User) (access, refresh string, err error) {
	claims := jwt.Claims{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Username:  entity.Username,
	}

	return jwt.Create(claims)
}
