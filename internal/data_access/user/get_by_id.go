package user

import "github.com/maadiii/helli/internal/entity"

func (d *dataAccess) GetByID(entity *entity.User) (err error) {
	for _, user := range d.users {
		if user.ID == entity.ID {
			*entity = *user
		}
	}

	return
}
