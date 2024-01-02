package account

import (
	"sync"

	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/hertzwrapper/errors"
)

type dataAccess struct {
	accounts map[string][]*entity.Account
	*sync.Mutex
}

func New() DataAccess {
	return &dataAccess{
		accounts: make(map[string][]*entity.Account, 0),
		Mutex:    new(sync.Mutex),
	}
}

type DataAccess interface {
	Insert(entity *entity.Account) error
	Increase(entity *entity.Account, amount float64) error
	Decrease(entity *entity.Account, amount float64) error
	Transfer(src *entity.Account, dst *entity.Account, amount float64) error
	Get(entity *entity.Account) error
}

func (d *dataAccess) canDecrease(username string, id int, amount float64) (account *entity.Account, err error) {
	account, err = d.get(username, id)
	if err != nil {
		return nil, err
	}

	if account.Fund-amount < 0 {
		return nil, errors.Conflict.Pattern("there is not enough balance in this account")
	}

	return
}

func (d *dataAccess) get(username string, id int) (*entity.Account, error) {
	for uname, accounts := range d.accounts {
		if uname == username {
			for i := range accounts {
				if accounts[i].ID == id {
					return accounts[i], nil
				}
			}
		}
	}

	return nil, errors.NotFound.Pattern("account not found")
}
