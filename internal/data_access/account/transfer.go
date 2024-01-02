package account

import (
	"github.com/maadiii/helli/internal/entity"
)

func (d *dataAccess) Transfer(src *entity.Account, dst *entity.Account, amount float64) (err error) {
	d.Lock()
	defer d.Unlock()

	src, err = d.canDecrease(src.Username, src.ID, amount)
	if err != nil {
		return
	}

	dst, err = d.get(dst.Username, dst.ID)
	if err != nil {
		return
	}

	src.Fund -= amount
	dst.Fund += amount

	return
}
