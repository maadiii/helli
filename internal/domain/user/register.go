package user

import (
	"github.com/maadiii/helli/internal/entity"
)

func (s *service) Register(entity *entity.User) error {
	return s.users.Insert(entity)
}
