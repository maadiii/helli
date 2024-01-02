package account

import "github.com/maadiii/helli/internal/entity"

func (d *dataAccess) Increase(entity *entity.Account, amount float64) (err error) {
	d.Lock()
	defer d.Unlock()

	entity, err = d.get(entity.Username, entity.ID)
	if err != nil {
		return
	}

	entity.Fund += amount

	return
}
