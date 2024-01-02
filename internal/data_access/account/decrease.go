package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/hertzwrapper/errors"
)

func (d *dataAccess) Decrease(entity *entity.Account, amount float64) (err error) {
	d.Lock()
	defer d.Unlock()

	for username := range d.accounts {
		if username == entity.Username {
			return d.decrease(username, entity.ID, amount)
		}
	}

	return errors.NotFound.Pattern("account not found")
}

func (d *dataAccess) decrease(username string, id int, amount float64) (err error) {
	account, err := d.canDecrease(username, id, amount)
	if err != nil {
		return
	}

	account.Fund -= amount

	return
}
