package user

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/hertzwrapper/errors"
)

func (d *dataAccess) Insert(entity *entity.User) (err error) {
	d.Lock()
	defer d.Unlock()

	if _, ok := d.users[entity.Username]; ok {
		return errors.Conflict.Pattern("username already exist")
	}

	entity.ID = len(d.users) + 1
	d.users[entity.Username] = entity

	return
}
