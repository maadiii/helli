package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/event"
)

func (s *service) Create(entity *entity.Account) (err error) {
	if err = s.accounts.Insert(entity); err != nil {
		return
	}

	s.raiseCreate(entity)

	return
}

func (s *service) raiseCreate(entity *entity.Account) {
	go func() {
		event := &event.Event{
			Action:    event.Create,
			Username:  entity.Username,
			AccountID: entity.ID,
			Amount:    entity.Fund,
		}

		s.event.Raise(*event)
	}()
}
