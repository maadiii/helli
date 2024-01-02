package account

import (
	"github.com/maadiii/helli/internal/entity"
)

func (d *dataAccess) Get(entity *entity.Account) (err error) {
	d.Lock()
	defer d.Unlock()

	account, err := d.get(entity.Username, entity.ID)
	if err != nil {
		return
	}

	*entity = *account

	return
}
