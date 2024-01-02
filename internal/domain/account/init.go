package account

import (
	"github.com/maadiii/helli/internal/data_access/account"
	"github.com/maadiii/helli/internal/data_access/user"
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/event"
)

type service struct {
	accounts account.DataAccess
	users    user.DataAccess
	event    event.Raiser
}

func New(
	dataAccess account.DataAccess,
	usersDataAccess user.DataAccess,
	event event.Raiser,
) Service {
	return &service{dataAccess, usersDataAccess, event}
}

type Service interface {
	Create(entity *entity.Account) error
	Increase(entity *entity.Account, amount float64) error
	Decrease(entity *entity.Account, amount float64) error
	Transfer(src *entity.Account, dstUserID, dstAccountID int, amount float64) error
	Get(entity *entity.Account) error
}
