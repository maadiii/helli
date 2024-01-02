package user

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/hertzwrapper/errors"
)

func (d *dataAccess) GetByUsername(entity *entity.User) (err error) {
	d.Lock()
	defer d.Unlock()

	user, ok := d.users[entity.Username]
	if !ok {
		return errors.NotFound.Pattern("username not found")
	}

	*entity = *user

	return
}
