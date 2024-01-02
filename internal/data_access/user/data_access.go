package user

import (
	"sync"

	"github.com/maadiii/helli/internal/entity"
)

type dataAccess struct {
	users map[string]*entity.User
	*sync.Mutex
}

func New() DataAccess {
	return &dataAccess{
		users: make(map[string]*entity.User),
		Mutex: new(sync.Mutex),
	}
}

type DataAccess interface {
	Insert(entity *entity.User) error
	GetByUsername(entity *entity.User) error
	GetByID(entity *entity.User) error
}
