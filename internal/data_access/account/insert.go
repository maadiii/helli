package account

import "github.com/maadiii/helli/internal/entity"

func (d *dataAccess) Insert(entity *entity.Account) (err error) {
	d.Lock()
	defer d.Unlock()

	entity.ID = len(d.accounts[entity.Username]) + 1
	d.accounts[entity.Username] = append(d.accounts[entity.Username], entity)

	return
}
