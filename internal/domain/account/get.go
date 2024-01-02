package account

import "github.com/maadiii/helli/internal/entity"

func (s *service) Get(entity *entity.Account) error {
	return s.accounts.Get(entity)
}
