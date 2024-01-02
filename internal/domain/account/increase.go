package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/event"
)

func (s *service) Increase(entity *entity.Account, amount float64) (err error) {
	if err = s.accounts.Increase(entity, amount); err != nil {
		return
	}

	s.raiseIncrease(entity, amount)

	return
}

func (s *service) raiseIncrease(entity *entity.Account, amount float64) {
	go func() {
		event := event.Event{
			Action:    event.Increase,
			Username:  entity.Username,
			AccountID: entity.ID,
			Amount:    amount,
		}

		s.event.Raise(event)
	}()
}
